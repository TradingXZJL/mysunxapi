package mysunxapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"golang.org/x/sync/errgroup"
)

const SUNX_API_WS = "api.sunx.io"

const (
	WS_NOTIFICATION = "notification"
	WS_MARKET       = "market"
)

const (
	SUNX_NOTIFICATION_WS_STREAM = "/ws/v1/notification"
	SUNX_MARKET_WS_STREAM       = "/ws/v1/market"
)

// op类型
const (
	SUBSCRIBE   = "sub"   //订阅
	UNSUBSCRIBE = "unsub" //取消订阅
)

var (
	// WebsocketTimeout is an interval for sending ping/pong messages if WebsocketKeepalive is enabled
	WebsocketTimeout = time.Second * 10
	// WebsocketKeepalive enables sending ping/pong messages to check the connection stability
	WebsocketKeepalive = true
	//SUBSCRIBE_INTERVAL_TIME = 500 * time.Millisecond //订阅间隔

)

type WsAPIType int

const (
	// Private
	WsAPITypeNotification WsAPIType = iota

	// Public
	WsAPITypeMarket
)

var (
	// WebSocket Server 每隔 5s(这个频率可能会变化) 会向 WebSocket Client 发起⼀一次⼼心跳，WebSocket Client 忽略心跳 5 次后，WebSocket Server 将会主动断开连接。
	WebSocketTimeout   = 5 * 5 * time.Second
	WebSocketKeepAlive = true
)

var node *snowflake.Node

func init() {
	node, _ = snowflake.NewNode(33)
}

type WsStreamClient struct {
	client *Client
	wsType WsAPIType
	conn   *websocket.Conn
	connId string

	waitSubscribeResMap MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]
	waitSubResult       *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]
	currentSubMap       MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]

	// public
	bboSubMap           MySyncMap[string, *Subscription[WsMarketCommonReq, WsBBO]]
	depthSubMap         MySyncMap[string, *Subscription[WsMarketCommonReq, WsDepth]]
	depthHighFreqSubMap MySyncMap[string, *Subscription[WsMarketDepthHighFreqReq, WsDepthHighFreq]]
	klineSubMap         MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]
	tradeDetailSubMap   MySyncMap[string, *Subscription[WsMarketCommonReq, WsTradeDetail]]
	// indexKlineSubMap MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]

	// private
	accountSubMap     MySyncMap[string, *Subscription[WsNotificationCommonReq, WsAccountRes]]
	positionsSubMap   MySyncMap[string, *Subscription[WsPositionsReq, WsPositions]]
	matchOrdersSubMap MySyncMap[string, *Subscription[WsMatchOrdersReq, WsMatchOrders]]
	tradeSubMap       MySyncMap[string, *Subscription[WsTradeReq, WsTrade]]
	ordersSubMap      MySyncMap[string, *Subscription[WsOrdersReq, WsOrders]]

	resultChan chan []byte
	errChan    chan error
	isClose    bool

	AutoReConnectTimes int //自动重连次数
	writeMu            sync.Mutex
}

func (ws *WsStreamClient) sendSubscribeResultToChan(result WsSubscribeResCommon) {
	if ws.connId == "" && result.Id != "" {
		ws.connId = result.Id
	}
	if result.Id != "" {
		if sub, ok := ws.waitSubscribeResMap.Load(result.Id); ok {
			if result.ErrCode == 0 || result.ErrMsg == "" || result.Status == "ok" {
				sub.resultChan <- result
			} else {
				sub.errChan <- fmt.Errorf("errHandler: %+v", result)
			}
			ws.waitSubscribeResMap.Delete(result.Id)
			return
		}
	} else if result.Cid != "" {
		if sub, ok := ws.waitSubscribeResMap.Load(result.Cid); ok {
			if result.ErrCode != 0.0 {
				sub.errChan <- fmt.Errorf("errHandler: %+v", result)
			} else {
				sub.resultChan <- result
			}
		}
		ws.waitSubscribeResMap.Delete(result.Cid)
		return
	}
	if ws.waitSubResult != nil {
		if result.ErrCode != 0 || result.ErrMsg != "" {
			ws.waitSubResult.errChan <- fmt.Errorf("errHandler: %+v", result)
		} else {
			ws.waitSubResult.resultChan <- result
		}
	}
}

// 解除所有订阅，清除所有订阅者
func (ws *WsStreamClient) sendWsCloseToAllSub() {
	ws.currentSubMap.Range(func(reqId string, sub *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]) bool {
		subKeys := make([]string, 0, len(sub.SubReqs))
		for _, req := range sub.SubReqs {
			key := req.Sub
			if key == "" {
				key = req.Topic
			}
			if key != "" {
				subKeys = append(subKeys, key)
			}
		}
		ws.sendUnSubscribeSuccessToCloseChan(reqId, subKeys)
		return true
	})
}

func (ws *WsStreamClient) sendUnSubscribeSuccessToCloseChan(reqId string, subKeys []string) {
	if _, ok := ws.currentSubMap.Load(reqId); ok {
		ws.currentSubMap.Delete(reqId)
	}

	for _, key := range subKeys {
		if key == "" {
			continue
		}

		// 删除 bbo 订阅者
		if sub, ok := ws.bboSubMap.Load(key); ok {
			ws.bboSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除深度订阅者
		if sub, ok := ws.depthSubMap.Load(key); ok {
			ws.depthSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除深度增量订阅者
		if sub, ok := ws.depthHighFreqSubMap.Load(key); ok {
			ws.depthHighFreqSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除K线订阅者
		if sub, ok := ws.klineSubMap.Load(key); ok {
			ws.klineSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除 Trade Detail 订阅者
		if sub, ok := ws.tradeDetailSubMap.Load(key); ok {
			ws.tradeDetailSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除账户资金订阅者
		if sub, ok := ws.accountSubMap.Load(key); ok {
			ws.accountSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除持仓变动订阅者
		if sub, ok := ws.positionsSubMap.Load(key); ok {
			ws.positionsSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除撮合订单订阅者
		if sub, ok := ws.matchOrdersSubMap.Load(key); ok {
			ws.matchOrdersSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除成交变动订阅者
		if sub, ok := ws.tradeSubMap.Load(key); ok {
			ws.tradeSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}

		// 删除订单订阅者
		if sub, ok := ws.ordersSubMap.Load(key); ok {
			ws.ordersSubMap.Delete(key)
			if sub.closeChan != nil {
				select {
				case sub.closeChan <- struct{}{}:
				default:
				}
			}
		}
	}
}

func (ws *WsStreamClient) reSubscribeForReconnect() error {
	var errG errgroup.Group
	ws.currentSubMap.Range(func(_ string, sub *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]) bool {
		// 如果没有保存请求参数，跳过
		if len(sub.SubReqs) == 0 {
			return true
		}

		errG.Go(func() error {
			reqId := sub.SubReqs[0].Id
			if reqId == "" {
				reqId = sub.SubReqs[0].Cid
			}
			if reqId == "" {
				reqId = node.Generate().String()
			}
			newSub, err := subscribe[WsSubscribeReqCommon, WsSubscribeResCommon](ws, sub.SubReqs, reqId)
			if err != nil {
				log.Error(err)
				return err
			}

			sub.SubId = newSub.SubId
			return nil
		})
		return true
	})
	return errG.Wait()
}

type PublicWsStreamClient struct {
	WsStreamClient
}
type PrivateWsStreamClient struct {
	WsStreamClient
}

func (*MySunx) NewPublicWsStreamClient(wsType WsAPIType) *PublicWsStreamClient {
	return &PublicWsStreamClient{
		WsStreamClient: WsStreamClient{
			wsType:              wsType,
			waitSubscribeResMap: MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]{},
			currentSubMap:       MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]{},

			bboSubMap:           MySyncMap[string, *Subscription[WsMarketCommonReq, WsBBO]]{},
			depthSubMap:         MySyncMap[string, *Subscription[WsMarketCommonReq, WsDepth]]{},
			depthHighFreqSubMap: MySyncMap[string, *Subscription[WsMarketDepthHighFreqReq, WsDepthHighFreq]]{},
			klineSubMap:         MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]{},
			tradeDetailSubMap:   MySyncMap[string, *Subscription[WsMarketCommonReq, WsTradeDetail]]{},
			// indexKlineSubMap: MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]{},

			accountSubMap:     MySyncMap[string, *Subscription[WsNotificationCommonReq, WsAccountRes]]{},
			positionsSubMap:   MySyncMap[string, *Subscription[WsPositionsReq, WsPositions]]{},
			matchOrdersSubMap: MySyncMap[string, *Subscription[WsMatchOrdersReq, WsMatchOrders]]{},
			tradeSubMap:       MySyncMap[string, *Subscription[WsTradeReq, WsTrade]]{},

			resultChan: make(chan []byte),
			errChan:    make(chan error),
			isClose:    false,

			AutoReConnectTimes: 0,
			writeMu:            sync.Mutex{},
		},
	}
}

func (*MySunx) NewPrivateWsStreamClient(accessKey, secretKey string, wsType WsAPIType) *PrivateWsStreamClient {
	return &PrivateWsStreamClient{
		WsStreamClient: WsStreamClient{
			client: &Client{
				AccessKey: accessKey,
				SecretKey: secretKey,
			},
			wsType:              wsType,
			waitSubscribeResMap: MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]{},
			currentSubMap:       MySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]{},

			bboSubMap:           MySyncMap[string, *Subscription[WsMarketCommonReq, WsBBO]]{},
			depthSubMap:         MySyncMap[string, *Subscription[WsMarketCommonReq, WsDepth]]{},
			depthHighFreqSubMap: MySyncMap[string, *Subscription[WsMarketDepthHighFreqReq, WsDepthHighFreq]]{},
			klineSubMap:         MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]{},
			tradeDetailSubMap:   MySyncMap[string, *Subscription[WsMarketCommonReq, WsTradeDetail]]{},
			// indexKlineSubMap: MySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]{},

			accountSubMap:     MySyncMap[string, *Subscription[WsNotificationCommonReq, WsAccountRes]]{},
			positionsSubMap:   MySyncMap[string, *Subscription[WsPositionsReq, WsPositions]]{},
			matchOrdersSubMap: MySyncMap[string, *Subscription[WsMatchOrdersReq, WsMatchOrders]]{},
			tradeSubMap:       MySyncMap[string, *Subscription[WsTradeReq, WsTrade]]{},
			ordersSubMap:      MySyncMap[string, *Subscription[WsOrdersReq, WsOrders]]{},

			resultChan: make(chan []byte),
			errChan:    make(chan error),
			isClose:    false,

			AutoReConnectTimes: 0,
			writeMu:            sync.Mutex{},
		},
	}
}

type Subscription[T any, R any] struct {
	SubId      string                  //订阅ID
	Ws         *WsStreamClient         //订阅的连接
	Op         string                  //订阅方法
	Topic      string                  //订阅主题
	Req        *T                      //订阅参数
	SubReqs    []*WsSubscribeReqCommon //订阅参数列表(用于重连)
	Res        *R                      //订阅返回结果
	resultChan chan R                  //接收订阅结果的通道
	errChan    chan error              //接收订阅错误的通道
	closeChan  chan struct{}           //接收订阅关闭的通道
	// subResultMap map[string]bool //订阅结果
}

// 获取订阅结果
func (s *Subscription[T, R]) ResultChan() chan R {
	return s.resultChan
}

// 获取错误订阅
func (s *Subscription[T, R]) ErrChan() chan error {
	return s.errChan
}

// 获取关闭订阅信号
func (s *Subscription[T, R]) CloseChan() chan struct{} {
	return s.closeChan
}

// 取消订阅
func (s *Subscription[T, R]) Unsubscribe() error {
	if s.Ws == nil || s.Ws.conn == nil || s.Ws.isClose {
		return fmt.Errorf("websocket is closed")
	}

	if len(s.SubReqs) == 0 {
		return fmt.Errorf("no subscription requests to unsubscribe")
	}

	unsubReqs := make([]*WsSubscribeReqCommon, 0, len(s.SubReqs))
	subKeys := make([]string, 0, len(s.SubReqs))

	for _, req := range s.SubReqs {
		unsubReq := &WsSubscribeReqCommon{
			Id:       req.Id,
			Cid:      req.Cid,
			DataType: req.DataType,
		}

		if req.Sub != "" {
			unsubReq.Unsub = req.Sub
			subKeys = append(subKeys, req.Sub)
		}
		if req.Op != "" {
			unsubReq.Op = "unsub"
			unsubReq.Topic = req.Topic
			unsubReq.ContractCode = req.ContractCode
			if req.Topic != "" {
				subKeys = append(subKeys, req.Topic)
			}
		}

		unsubReqs = append(unsubReqs, unsubReq)
	}

	// 发送取消订阅请求
	for _, req := range unsubReqs {
		data, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("marshal unsubscribe request failed: %w", err)
		}
		log.Debugf("ws unsubscribe req: %s", string(data))
		s.Ws.writeMu.Lock()
		err = s.Ws.conn.WriteMessage(websocket.TextMessage, data)
		s.Ws.writeMu.Unlock()
		if err != nil {
			return fmt.Errorf("send unsubscribe request failed: %w", err)
		}
	}

	log.Debugf("Unsubscribe Success args:%v", s.SubReqs)

	// 取消订阅成功，统一清理资源
	s.Ws.sendUnSubscribeSuccessToCloseChan(s.SubId, subKeys)

	return nil
}

// 通用订阅请求结构
type WsSubscribeReqCommon struct {
	// public
	Sub      string `json:"sub,omitempty"`
	Unsub    string `json:"unsub,omitempty"`
	Id       string `json:"id"`
	DataType string `json:"data_type,omitempty"`

	// Private
	Op           string `json:"op,omitempty"`
	Cid          string `json:"cid,omitempty"`
	Topic        string `json:"topic,omitempty"`
	ContractCode string `json:"contractCode,omitempty"`
}

// 通用订阅返回结果结构
type WsSubscribeResCommon struct {
	// Common
	Ts int64 `json:"ts"`

	// Public
	Id       string `json:"id,omitempty"`
	Status   string `json:"status,omitempty"`
	DataType string `json:"data_type,omitempty"`
	Subbed   string `json:"subbed,omitempty"`
	UnSub    string `json:"unsub,omitempty"`

	// Private
	// Data         T      `json:"data,omitempty"`
	Op           string `json:"op,omitempty"`
	Topic        string `json:"topic,omitempty"`
	ContractCode string `json:"contract_code,omitempty"`
	Cid          string `json:"cid,omitempty"`
	Uid          string `json:"uid,omitempty"`
	Event        string `json:"event,omitempty"`    // snapshot
	ErrCode      any    `json:"err-code,omitempty"` // int or string
	ErrMsg       string `json:"err-msg,omitempty"`
	Type         string `json:"type,omitempty"`
}

func subscribe[T any, R any](ws *WsStreamClient, subscribeReq []*T, reqId string) (*Subscription[T, R], error) {
	if ws == nil || ws.conn == nil || ws.isClose {
		return nil, fmt.Errorf("websocket is closed")
	}

	// 发送请求前，构造一个接收订阅结果的订阅者
	waitSubResult := &Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]{
		SubId:      node.Generate().String(),
		Ws:         ws,
		resultChan: make(chan WsSubscribeResCommon, 50), // 给个缓冲区
		errChan:    make(chan error, 1),
		closeChan:  make(chan struct{}, 1),
		SubReqs:    make([]*WsSubscribeReqCommon, 0, len(subscribeReq)),
	}

	for _, req := range subscribeReq {
		data, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		// 将请求参数转换为通用结构保存，用于重连
		commonReq := &WsSubscribeReqCommon{}
		if err := json.Unmarshal(data, commonReq); err == nil {
			waitSubResult.SubReqs = append(waitSubResult.SubReqs, commonReq)
		}

		log.Debugf("ws subscribe req: %s", string(data))
		ws.writeMu.Lock()
		err = ws.conn.WriteMessage(websocket.TextMessage, data)
		ws.writeMu.Unlock()
		if err != nil {
			return nil, err
		}
	}

	ws.waitSubscribeResMap.Store(reqId, waitSubResult)

	// 同步捕获订阅结果
	err := ws.catchSubscribeResult(waitSubResult)
	if err != nil {
		return nil, err
	}

	dataSubResult := &Subscription[T, R]{
		SubId:      reqId,
		Ws:         ws,
		resultChan: make(chan R, 50),
		errChan:    make(chan error, 1),
		closeChan:  make(chan struct{}, 1),
		SubReqs:    waitSubResult.SubReqs,
	}

	return dataSubResult, nil
}

func (ws *WsStreamClient) Close() error {
	ws.isClose = true

	err := ws.conn.Close()
	if err != nil {
		return err
	}

	//手动关闭成功，给所有订阅发送关闭信号
	ws.sendWsCloseToAllSub()

	//初始化连接状态
	ws.conn = nil
	ws.connId = ""
	close(ws.resultChan)
	close(ws.errChan)
	ws.waitSubscribeResMap = NewMySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]()

	ws.currentSubMap = NewMySyncMap[string, *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]]()
	ws.bboSubMap = NewMySyncMap[string, *Subscription[WsMarketCommonReq, WsBBO]]()
	ws.depthSubMap = NewMySyncMap[string, *Subscription[WsMarketCommonReq, WsDepth]]()
	ws.depthHighFreqSubMap = NewMySyncMap[string, *Subscription[WsMarketDepthHighFreqReq, WsDepthHighFreq]]()
	ws.klineSubMap = NewMySyncMap[string, *Subscription[WsMarketCommonReq, WsKline]]()

	return nil
}

func (ws *WsStreamClient) catchSubscribeResult(sub *Subscription[WsSubscribeReqCommon, WsSubscribeResCommon]) error {
	isBreak := false
	for {
		select {
		case err := <-sub.ErrChan():
			log.Error(err)
			return fmt.Errorf("subscribe failed: %s", err)
		case subResult := <-sub.ResultChan():
			if subResult.Status != "ok" && subResult.ErrCode != 0.0 {
				log.Errorf("subscribe failed: %s", subResult.Status)
				return fmt.Errorf("subscribe failed: %s", subResult.Status)
			}
			if ws.connId == "" {
				if subResult.Id != "" {
					ws.connId = subResult.Id
				} else if subResult.Cid != "" {
					ws.connId = subResult.Cid
				}
			}
			sub.Res = &subResult
			ws.waitSubscribeResMap.Delete(sub.SubId)
			ws.currentSubMap.Store(sub.SubId, sub)

			isBreak = true
		case <-sub.CloseChan():
			return errors.New("subscribe closed")
		}
		if isBreak {
			break
		}
	}

	// d, _ := json.Marshal(sub.Res)
	// log.Debugf("subscribe result: %s", string(d))

	return nil
}

func (ws *WsStreamClient) handleResult(resultChan chan []byte, errChan chan error) {
	go func() {
		for {
			select {
			case err, ok := <-errChan:
				if !ok {
					log.Error("errChan is closed")
					return
				}
				log.Error(err)
				//错误处理 重连等
				//ws标记为非关闭 且返回错误包含EOF、close、reset时自动重连
				if !ws.isClose && (strings.Contains(err.Error(), "EOF") ||
					strings.Contains(err.Error(), "close") ||
					strings.Contains(err.Error(), "reset")) {
					//重连
					err := ws.OpenConn()
					for err != nil {
						time.Sleep(1500 * time.Millisecond)
						err = ws.OpenConn()
					}
					ws.AutoReConnectTimes += 1
					go func() {
						//重新订阅
						err = ws.reSubscribeForReconnect()
						if err != nil {
							log.Error(err)
						}
					}()
				} else {
					continue
				}
			case data, ok := <-resultChan:
				if !ok {
					log.Error("resultChan is closed")
					return
				}

				// Auth Success
				if strings.Contains(string(data), "op\":\"auth") {
					log.Debug("auth success: ", string(data))
					continue
				}

				// Auth Error
				if strings.Contains(string(data), "op\":\"error") {
					log.Error("sub error: ", string(data))
					ws.isClose = true
					ws.errChan <- fmt.Errorf("sub error: %s", string(data))
					continue
				}

				// log.Warn(string(data))
				//处理订阅或查询订阅列表请求返回结果
				if strings.Contains(string(data), "subbed") || strings.Contains(string(data), "unsub") ||
					strings.Contains(string(data), "status") || strings.Contains(string(data), "op\":\"sub") {
					result := WsSubscribeResCommon{}
					err := json.Unmarshal(data, &result)
					if err != nil {
						log.Error(err)
						continue
					}
					d, _ := json.MarshalIndent(result, "", "  ")
					log.Debugf("subscribe result: %s", string(d))
					ws.sendSubscribeResultToChan(result)
					continue
				}

				// 处理买一卖一逐笔行情推送订阅结果
				if strings.Contains(string(data), ".bbo") {
					bboRes, err := handleWsData[WsBBORes](data)
					// log.Info("data: ", string(data))
					// log.Info("bboRes: ", bboRes)
					if sub, ok := ws.bboSubMap.Load(bboRes.Ch); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *bboRes.convertToWsBBO()
						continue
					}
				}

				// 处理 Market Depth 数据
				if strings.Contains(string(data), ".depth.") {
					ch, err := handleWsData[WsMarketCh](data)
					if err != nil {
						log.Error(err)
						continue
					}
					if len(strings.Split(ch.Ch, ".")) == 4 {
						depthRes, err := handleWsData[WsDepthRes](data)
						if err != nil {
							log.Error(err)
							continue
						}
						if sub, ok := ws.depthSubMap.Load(depthRes.Ch); ok {
							sub.resultChan <- *depthRes.convertToWsDepthRes()
							continue
						}
					}
				}

				// 处理 Market Depth 增量数据
				if strings.Contains(string(data), ".depth.size_") {
					depthHFRes, err := handleWsData[WsDepthHighFreqRes](data)
					if sub, ok := ws.depthHighFreqSubMap.Load(depthHFRes.Ch); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *depthHFRes.convertToWsDepthHighFreq()
						continue
					}
				}

				// 处理 K线数据
				if strings.Contains(string(data), ".kline.") {
					klineRes, err := handleWsData[WsKline](data)
					if sub, ok := ws.klineSubMap.Load(klineRes.Ch); ok {
						if err != nil {
							sub.errChan <- err
						}
						sub.resultChan <- *klineRes
						continue
					}
				}

				// 处理 Trade Detail 数据
				if strings.Contains(string(data), ".trade.detail") {
					tradeDetailRes, err := handleWsData[WsTradeDetail](data)
					if sub, ok := ws.tradeDetailSubMap.Load(tradeDetailRes.Ch); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *tradeDetailRes
						continue
					}
				}

				// 处理指数K线数据
				// if strings.Contains(string(data), ".index.") {
				// 	indexKlineRes, err := handleWsData[WsKline](data)
				// 	if sub, ok := ws.indexKlineSubMap.Load(indexKlineRes.Ch); ok {
				// 		if err != nil {
				// 			sub.errChan <- err
				// 		}
				// 		sub.resultChan <- *indexKlineRes
				// 	}
				// }

				// 处理账户资金数据
				if strings.Contains(string(data), "topic\":\"account") {
					accountRes, err := handleWsData[WsAccountRes](data)
					if sub, ok := ws.accountSubMap.Load(accountRes.Topic); ok {
						// log.Warn("accountRes: ", accountRes)
						if err != nil {
							sub.errChan <- err
						}
						sub.resultChan <- *accountRes
						continue
					}
				}

				// 处理持仓变动数据
				if strings.Contains(string(data), "topic\":\"positions") {
					positionsRes, err := handleWsData[WsPositions](data)
					if sub, ok := ws.positionsSubMap.Load(positionsRes.Topic); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *positionsRes
						continue
					}
				}

				// 处理撮合后的订单数据
				if strings.Contains(string(data), "topic\":\"match_orders") {
					matchOrdersRes, err := handleWsData[WsMatchOrders](data)
					if sub, ok := ws.matchOrdersSubMap.Load(matchOrdersRes.Topic); ok {
						if err != nil {
							sub.errChan <- err
							continue
						}
						sub.resultChan <- *matchOrdersRes
						continue
					}
				}

				// 处理成交变动数据
				if strings.Contains(string(data), "topic\":\"trade") {
					topicRes, err := handleWsData[WsNotificationTopic](data)
					if err != nil {
						// log.Error(err)
						continue
					}

					if len(strings.Split(topicRes.Topic, ".")) == 1 {
						tradeRes, err := handleWsData[WsTrade](data)

						if sub, ok := ws.tradeSubMap.Load(tradeRes.Topic); ok {
							if err != nil {
								sub.errChan <- err
								continue
							}
							sub.resultChan <- *tradeRes
							continue
						}
					}
				}

				// 处理订单数据
				if strings.Contains(string(data), "orders") {
					topicRes, err := handleWsData[WsNotificationTopic](data)
					if err != nil {
						continue
					}
					if len(strings.Split(topicRes.Topic, ".")) == 1 {
						ordersRes, err := handleWsData[WsOrders](data)
						if sub, ok := ws.ordersSubMap.Load(ordersRes.Topic); ok {
							if err != nil {
								sub.errChan <- err
								continue
							}
							sub.resultChan <- *ordersRes
							continue
						}
					}
				}
			}
		}
	}()
}
func (ws *WsStreamClient) OpenConn() error {
	if ws.resultChan != nil {
		ws.resultChan = make(chan []byte)
	}
	if ws.errChan != nil {
		ws.errChan = make(chan error)
	}

	apiUrl, isPrivate := handlerWsStreamRequestApi(ws)
	if ws.conn == nil {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan, isPrivate)
		ws.conn = conn
		ws.isClose = false
		ws.connId = ""
		log.Infof("OpenConn success to: %s", apiUrl)

		ws.handleResult(ws.resultChan, ws.errChan)

		// 如果是私有连接，连接建立后发送鉴权消息
		if isPrivate {
			err = ws.SendAuthMessage()
			if err != nil {
				log.Errorf("SendAuthMessage error: %v", err)
				return err
			}
		}

		return err
	} else {
		conn, err := wsStreamServe(apiUrl, ws.resultChan, ws.errChan, isPrivate)
		ws.conn = conn
		ws.connId = ""
		log.Info("Auto ReOpenConn success to:", apiUrl)

		ws.handleResult(ws.resultChan, ws.errChan)

		// 如果是私有连接，重连后也需要发送鉴权消息
		if isPrivate {
			err = ws.SendAuthMessage()
			if err != nil {
				log.Errorf("SendAuthMessage error: %v", err)
				return err
			}
		}

		return err
	}
}

// WsAuthReq 鉴权请求结构
type WsAuthReq struct {
	Op               string `json:"op"`
	Type             string `json:"type"`
	AccessKeyId      string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
}

func handlerWsStreamRequestApi(ws *WsStreamClient) (string, bool) {
	path, isPrivate := getWsApiPath(ws.wsType)

	if isPrivate {
		// Private 接口需要先建立连接，然后发送鉴权消息
		// 鉴权消息在连接建立后发送，不放在 URL 参数里
		u := url.URL{
			Scheme: "wss",
			Host:   SUNX_API_WS,
			Path:   path,
		}
		return u.String(), isPrivate
	}

	u := url.URL{
		Scheme: "wss",
		Host:   SUNX_API_WS,
		Path:   path,
	}
	return u.String(), isPrivate
}

// SendAuthMessage 发送鉴权消息
func (ws *WsStreamClient) SendAuthMessage() error {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	reqMap := make(map[string]string)
	reqMap["AccessKeyId"] = ws.client.AccessKey
	reqMap["SignatureMethod"] = "HmacSHA256"
	reqMap["SignatureVersion"] = "2"
	reqMap["Timestamp"] = timestamp

	// 签名计算逻辑
	keys := []string{}
	for k := range reqMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	query := ""
	for _, k := range keys {
		// 签名时：参数名=参数值&... (参数值需要 URL Encode)
		query += k + "=" + url.QueryEscape(reqMap[k]) + "&"
	}
	query = strings.TrimRight(query, "&")

	path, _ := getWsApiPath(ws.wsType)
	hmacSha256Data := "GET" + "\n" + SUNX_API_WS + "\n" + path + "\n" + query
	// log.Infof("hmacSha256Data: %v", hmacSha256Data)
	sign := HmacSha256(ws.client.SecretKey, hmacSha256Data)
	signature := base64.StdEncoding.EncodeToString(sign)

	authReq := WsAuthReq{
		Op:               "auth",
		Type:             "api",
		AccessKeyId:      ws.client.AccessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature,
	}

	data, err := json.Marshal(authReq)
	if err != nil {
		return err
	}
	log.Debugf("Sending auth message: %s", string(data))

	ws.writeMu.Lock()
	defer ws.writeMu.Unlock()
	return ws.conn.WriteMessage(websocket.TextMessage, data)
}

func getWsApiPath(wsType WsAPIType) (string, bool) {
	switch wsType {
	case WsAPITypeNotification:
		return SUNX_NOTIFICATION_WS_STREAM, true
	case WsAPITypeMarket:
		return SUNX_MARKET_WS_STREAM, false
	}
	return "", false
}

// Gzip解压函数
func GzipDecode(in []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return io.ReadAll(reader)
}

func wsStreamServe(api string, resultChan chan []byte, errChan chan error, isPrivate bool) (*websocket.Conn, error) {
	dialer := websocket.DefaultDialer
	if WsUseProxy {
		proxy, err := getRandomProxy()
		if err != nil {
			return nil, err
		}
		url_i := url.URL{}
		targetProxy, _ := url_i.Parse(proxy.ProxyUrl)
		dialer.Proxy = http.ProxyURL(targetProxy)
	}
	c, _, err := dialer.Dial(api, nil)
	if err != nil {
		return nil, err
	}
	c.SetReadLimit(6553500)
	go func() {
		if WebsocketKeepalive {
			keepAlive(c, WebsocketTimeout)
		}
		isPrivate := isPrivate
		defer c.Close()
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				errChan <- err
				return
			}
			// log.Info("message: ", string(message))

			// 尝试 Gzip 解压
			// Gzip Magic Number: 1f 8b
			if len(message) >= 2 && message[0] == 0x1f && message[1] == 0x8b {
				decoded, err := GzipDecode(message)
				if err != nil {
					log.Error("Gzip decode error:", err)
					// 如果解压失败，可能不是Gzip，或者数据损坏，视情况继续处理原始数据或丢弃
				} else {
					message = decoded
				}
			}

			isPing, err := handlePingPong(c, message, isPrivate)
			if err != nil {
				errChan <- err
				return
			}
			if isPing {
				continue
			}
			resultChan <- message
		}
	}()

	return c, err
}

func handlePingPong(c *websocket.Conn, message []byte, isPrivate bool) (bool, error) {
	if isPrivate {
		var ppr PrivatePingPongRes
		err := json.Unmarshal(message, &ppr)
		// 不是pingpong信息
		if err != nil {
			return false, nil
		}
		if ppr.Op == "error" {
			return true, fmt.Errorf("private ping pong error: %s", ppr.ErrMsg)
		} else if ppr.Op == "ping" {
			err = ppr.SendPong(c)
			if err != nil {
				return true, err
			}
			return true, nil
		}
	} else {
		var ppr PublicPingResp
		err := json.Unmarshal(message, &ppr)
		if err != nil {
			return false, nil
		}

		if ppr.Ping > 0 {
			err = ppr.SendPong(c)
			if err != nil {
				return true, err
			}
			return true, nil
		}
	}
	return false, nil
}

// public ping
type PublicPingResp struct {
	Ping int64 `json:"ping"`
}

// public pong
type PublicPongReq struct {
	Pong int64 `json:"pong"`
}

func (ppr *PublicPingResp) SendPong(c *websocket.Conn) error {
	var pongReq PublicPongReq
	pongReq.Pong = ppr.Ping
	data, err := json.Marshal(pongReq)
	if err != nil {
		return err
	}
	err = c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	log.Debug("public ping pong send pong: ", string(data))
	return nil
}

// private ping
type PrivatePingpongErr struct {
	ErrCode int    `json:"err-code,omitempty"`
	ErrMsg  string `json:"err-msg,omitempty"`
}

type PrivatePongReq struct {
	Op string `json:"op"` // pong
	Ts string `json:"ts"`
}

type PrivatePingPongRes struct {
	Op string `json:"op"` // ping or pong
	Ts string `json:"ts"`
	PrivatePingpongErr
}

func (ppr *PrivatePingPongRes) SendPong(c *websocket.Conn) error {
	pongReq := PrivatePongReq{
		Op: "pong",
		Ts: ppr.Ts,
	}
	data, err := json.Marshal(pongReq)
	if err != nil {
		return err
	}
	log.Debug("private ping pong send pong: ", string(data))
	return c.WriteMessage(websocket.TextMessage, data)
}

// 发送ping/pong消息以检查连接稳定性
func keepAlive(c *websocket.Conn, timeout time.Duration) {
	ticker := time.NewTicker(timeout)

	lastResponse := time.Now()
	c.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > 3*timeout {
				err := c.Close()
				if err != nil {
					return
				}
				return
			}
		}
	}()
}
