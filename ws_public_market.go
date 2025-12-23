package mysunxapi

import (
	"fmt"
)

type WsOrderBookSize int

const (
	WS_ORDER_BOOK_SIZE_20  WsOrderBookSize = 20  // 20: stands for 20 unmerged data.
	Ws_ORDER_BOOK_SIZE_150 WsOrderBookSize = 150 // 150:stands for 150 unmerged data.
)

type WsMarketCommonReq struct {
	Sub   string `json:"sub,omitempty"`
	Unsub string `json:"unsub,omitempty"`
	Id    string `json:"id"`
}

// 订阅买一卖一逐笔行情推送 market.$contract_code.bbo
func (ws *PublicWsStreamClient) SubscribeMarketBBO(contractCodes []string, isSubscribe bool) (*Subscription[WsMarketCommonReq, WsBBO], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	reqId := node.Generate().String()
	subReqs := []*WsMarketCommonReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			subKey := fmt.Sprintf("market.%s.bbo", contractCode)
			subReqs = append(subReqs, &WsMarketCommonReq{
				Sub: subKey,
				Id:  reqId,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subKey := fmt.Sprintf("market.%s.bbo", contractCode)
			subReqs = append(subReqs, &WsMarketCommonReq{
				Unsub: subKey,
				Id:    reqId,
			})
		}
	}

	sub, err := subscribe[WsMarketCommonReq, WsBBO](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMarketBBO error: ", err)
		return nil, err
	}

	for _, subReq := range subReqs {
		// log.Info("SubKey: ", subReq.Sub)
		ws.bboSubMap.Store(subReq.Sub, sub)
	}
	return sub, nil
}

// 订阅 Market Depth 数据 market.$contract_code.depth.$type
func (ws *PublicWsStreamClient) SubscribeMarketDepth(contractCodes []string, depthType []string, isSubscribe bool) (*Subscription[WsMarketCommonReq, WsDepth], error) {
	if len(contractCodes) == 0 || len(depthType) == 0 {
		return nil, fmt.Errorf("contractCode and depthType are required")
	}

	reqId := node.Generate().String()
	subReqs := []*WsMarketCommonReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			for _, depthType := range depthType {
				subKey := fmt.Sprintf("market.%s.depth.%s", contractCode, depthType)
				subReqs = append(subReqs, &WsMarketCommonReq{
					Sub: subKey,
					Id:  reqId,
				})
			}
		}
	} else {
		for _, contractCode := range contractCodes {
			for _, depthType := range depthType {
				subKey := fmt.Sprintf("market.%s.depth.%s", contractCode, depthType)
				subReqs = append(subReqs, &WsMarketCommonReq{
					Unsub: subKey,
					Id:    reqId,
				})
			}
		}
	}

	sub, err := subscribe[WsMarketCommonReq, WsDepth](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMarketDepth error: ", err)
		return nil, err
	}

	for _, subReq := range subReqs {
		ws.depthSubMap.Store(subReq.Sub, sub)
	}
	return sub, nil
}

type WsMarketDepthHighFreqReq struct {
	WsMarketCommonReq
	DataType string `json:"data_type"` // incremental
}

// 订阅Market Depth增量数据 market.$contract_code.depth.size_${size}.high_freq
func (ws *PublicWsStreamClient) SubscribeMarketDepthHighFreq(contractCodes []string, size []int, isSubscribe bool) (*Subscription[WsMarketDepthHighFreqReq, WsDepthHighFreq], error) {
	if len(contractCodes) == 0 || len(size) == 0 {
		return nil, fmt.Errorf("contractCode and size are required")
	}

	reqId := node.Generate().String()
	subReqs := []*WsMarketDepthHighFreqReq{}

	if isSubscribe {
		for _, contractCode := range contractCodes {
			for _, size := range size {
				subKey := fmt.Sprintf("market.%s.depth.size_%d.high_freq", contractCode, size)
				subReqs = append(subReqs, &WsMarketDepthHighFreqReq{
					WsMarketCommonReq: WsMarketCommonReq{
						Sub: subKey,
						Id:  reqId,
					},
					DataType: "incremental",
				})
			}
		}
	} else {
		for _, contractCode := range contractCodes {
			for _, size := range size {
				subKey := fmt.Sprintf("market.%s.depth.size_%d.high_freq", contractCode, size)
				subReqs = append(subReqs, &WsMarketDepthHighFreqReq{
					WsMarketCommonReq: WsMarketCommonReq{
						Unsub: subKey,
						Id:    reqId,
					},
					DataType: "incremental",
				})
			}
		}
	}

	sub, err := subscribe[WsMarketDepthHighFreqReq, WsDepthHighFreq](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMarketDepthHighFreq error: ", err)
		return nil, err
	}

	for _, subReq := range subReqs {
		ws.depthHighFreqSubMap.Store(subReq.WsMarketCommonReq.Sub, sub)
	}

	return sub, nil
}

// 订阅 KLine 数据 market.$contract_code.kline.$period
func (ws *PublicWsStreamClient) SubscribeMarketKline(contractCodes []string, period []string, isSubscribe bool) (*Subscription[WsMarketCommonReq, WsKline], error) {
	if len(contractCodes) == 0 || len(period) == 0 {
		return nil, fmt.Errorf("contractCode and period are required")
	}

	reqId := node.Generate().String()
	subReqs := []*WsMarketCommonReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			for _, period := range period {
				subKey := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
				subReqs = append(subReqs, &WsMarketCommonReq{
					Sub: subKey,
					Id:  reqId,
				})
			}
		}
	} else {
		for _, contractCode := range contractCodes {
			for _, period := range period {
				subKey := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
				subReqs = append(subReqs, &WsMarketCommonReq{
					Unsub: subKey,
					Id:    reqId,
				})
			}
		}
	}

	sub, err := subscribe[WsMarketCommonReq, WsKline](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMarketKline error: ", err)
		return nil, err
	}

	for _, subReq := range subReqs {
		ws.klineSubMap.Store(subReq.Sub, sub)
	}

	return sub, nil
}

// 订阅 Trade Detail 数据 market.$contract_code.trade.detail
func (ws *PublicWsStreamClient) SubscribeMarketTradeDetail(contractCodes []string, isSubscribe bool) (*Subscription[WsMarketCommonReq, WsTradeDetail], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	reqId := node.Generate().String()
	subReqs := []*WsMarketCommonReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			subKey := fmt.Sprintf("market.%s.trade.detail", contractCode)
			subReqs = append(subReqs, &WsMarketCommonReq{
				Sub: subKey,
				Id:  reqId,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subKey := fmt.Sprintf("market.%s.trade.detail", contractCode)
			subReqs = append(subReqs, &WsMarketCommonReq{
				Unsub: subKey,
				Id:    reqId,
			})
		}
	}

	sub, err := subscribe[WsMarketCommonReq, WsTradeDetail](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMarketTradeDetail error: ", err)
		return nil, err
	}

	for _, subReq := range subReqs {
		ws.tradeDetailSubMap.Store(subReq.Sub, sub)
	}

	return sub, nil
}

// // 订阅(sub)指数K线数据 market.$contract_code.index.$period
// func (ws *PublicWsStreamClient) SubscribeMarketIndexKline(contractCodes []string, period []string, isSubscribe bool) (*Subscription[WsMarketCommonReq, WsKline], error) {
// 	if len(contractCodes) == 0 || len(period) == 0 {
// 		return nil, fmt.Errorf("contractCode and period are required")
// 	}

// 	reqId := node.Generate().String()
// 	subReqs := []*WsMarketCommonReq{}
// 	if isSubscribe {
// 		for _, contractCode := range contractCodes {
// 			for _, period := range period {
// 				subKey := fmt.Sprintf("market.%s.index.%s", contractCode, period)
// 				subReqs = append(subReqs, &WsMarketCommonReq{
// 					Sub: subKey,
// 					Id:  reqId,
// 				})
// 			}
// 		}
// 	} else {
// 		for _, contractCode := range contractCodes {
// 			for _, period := range period {
// 				subKey := fmt.Sprintf("market.%s.index.%s", contractCode, period)
// 				subReqs = append(subReqs, &WsMarketCommonReq{
// 					Unsub: subKey,
// 					Id:    reqId,
// 				})
// 			}
// 		}
// 	}

// 	sub, err := subscribe[WsMarketCommonReq, WsKline](&ws.WsStreamClient, subReqs, reqId)
// 	if err != nil {
// 		// log.Error("SubscribeMarketIndexKline error: ", err)
// 		return nil, err
// 	}

// 	for _, subReq := range subReqs {
// 		ws.indexKlineSubMap.Store(subReq.Sub, sub)
// 	}

// 	return sub, nil
// }
