package mysunxapi

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"

	"github.com/robfig/cron/v3"
)

type RestProxy struct {
	ProxyUrl string //代理的协议IP端口URL

	PublicWeight       ProxyWeight
	PrivateTradeWeight ProxyWeight
	PrivateReadWeight  ProxyWeight
}

type ProxyWeight struct {
	RemainWeight  int  //剩余可用权重
	Is1032Limited bool //是否已被限制
}

func (w *ProxyWeight) restore(apiType APIType) {
	switch apiType {
	case PUBLIC:
		w.RemainWeight = 240
	case PRIVATE_TRADE, PRIVATE_READ:
		w.RemainWeight = 72
	}
	w.Is1032Limited = false
}

var proxyList = []*RestProxy{}

var UseProxy = false
var WsUseProxy = false

func GetCurrentProxyList() []*RestProxy {
	return proxyList
}

func SetUseProxy(useProxy bool, proxyUrls ...string) {
	UseProxy = useProxy
	var newProxyList []*RestProxy
	for _, proxyUrl := range proxyUrls {
		newProxyList = append(newProxyList, &RestProxy{
			ProxyUrl: proxyUrl,
			PublicWeight: ProxyWeight{
				RemainWeight:  240,
				Is1032Limited: false,
			},
			PrivateTradeWeight: ProxyWeight{
				RemainWeight:  72,
				Is1032Limited: false,
			},
			PrivateReadWeight: ProxyWeight{
				RemainWeight:  72,
				Is1032Limited: false,
			},
		})
	}
	proxyList = newProxyList
}

func SetWsUseProxy(useProxy bool) error {
	if !UseProxy {
		return errors.New("please set UseProxy first")
	}
	WsUseProxy = useProxy
	return nil
}

func isUseProxy() bool {
	return UseProxy
}

func init() {
	c := cron.New(cron.WithSeconds())
	//每3秒权重清零，状态恢复
	_, err := c.AddFunc("*/3 * * * * *", func() {
		for _, proxy := range proxyList {
			proxy.PublicWeight.restore(PUBLIC)
			proxy.PrivateTradeWeight.restore(PRIVATE_TRADE)
			proxy.PrivateReadWeight.restore(PRIVATE_READ)
		}
	})
	if err != nil {
		log.Error(err)
		return
	}
	c.Start()
}

// 获取最佳代理
func getBestProxyAndWeight(apiType APIType) (*RestProxy, *ProxyWeight, error) {
	var maxWeightProxy *RestProxy
	var maxWeight *ProxyWeight
	for _, proxy := range proxyList {
		var proxyWeight *ProxyWeight
		switch apiType {
		case PUBLIC:
			proxyWeight = &proxy.PublicWeight
		case PRIVATE_TRADE:
			proxyWeight = &proxy.PrivateTradeWeight
		case PRIVATE_READ:
			proxyWeight = &proxy.PrivateReadWeight
		default:
			return nil, nil, errors.New("apiType is invalid")
		}

		if proxyWeight.Is1032Limited {
			continue
		}
		if maxWeightProxy == nil {
			maxWeightProxy = proxy
			maxWeight = proxyWeight
			continue
		}
		if proxyWeight.RemainWeight > maxWeight.RemainWeight {
			maxWeightProxy = proxy
			maxWeight = proxyWeight
		}
	}
	log.Debug("maxHeightProxy: %v", maxWeightProxy)
	return maxWeightProxy, maxWeight, nil
}

// 获取随机代理
func getRandomProxy() (*RestProxy, error) {
	length := len(proxyList)
	if length == 0 {
		return nil, errors.New("proxyList is empty")
	}

	return proxyList[rand.Intn(length)], nil
}

// Request 发送请求
func Request(url string, reqBody []byte, method string, isGzip bool, apiType APIType) ([]byte, error) {
	return RequestWithHeader(url, reqBody, method, map[string]string{}, isGzip, apiType)
}

func RequestWithHeader(urlStr string, reqBody []byte, method string, headerMap map[string]string, isGzip bool, apiType APIType) ([]byte, error) {
	req, err := http.NewRequest(method, urlStr, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range headerMap {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	if isGzip { // 请求 header 添加 gzip
		req.Header.Add("Content-Encoding", "gzip")
		req.Header.Add("Accept-Encoding", "gzip")
	}

	log.Debug("reqURL: ", req.URL.String())
	if len(reqBody) > 0 {
		log.Debug("reqBody: ", string(reqBody))
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	var currentProxy *RestProxy
	var currentProxyWeight *ProxyWeight
	if UseProxy {
		currentProxy, currentProxyWeight, err = getBestProxyAndWeight(apiType)
		if err != nil {
			return nil, err
		}
		if currentProxy == nil || currentProxyWeight == nil || currentProxyWeight.RemainWeight <= 0 {
			return nil, errors.New("all proxy ip weight limit reached")
		}

		url_i := url.URL{}
		bestProxy, _ := url_i.Parse(currentProxy.ProxyUrl)

		reqProxy := &http.Transport{}
		reqProxy.Proxy = http.ProxyURL(bestProxy)                        // set proxy
		reqProxy.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //set ssl

		client.Transport = reqProxy
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		body, err = gzip.NewReader(resp.Body)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	data, err := io.ReadAll(body)
	// log.Debug(string(data))
	log.Debug(string(data))
	log.Debug(resp.Header)
	if isUseProxy() {
		//回填权重
		rateLimitRemaining := resp.Header.Get("ratelimit-remaining")
		if rateLimitRemaining == "" {
			// 尝试其他常见的大小写或变体
			rateLimitRemaining = resp.Header.Get("x-ratelimit-remaining")
		}

		if rateLimitRemaining != "" {
			log.Debug("rateLimitRemaining: ", rateLimitRemaining)
			remainWeight, err := strconv.Atoi(rateLimitRemaining)
			if err != nil {
				log.Errorf("parse ratelimit-remaining error: %v, val: %s", err, rateLimitRemaining)
			} else {
				if remainWeight < currentProxyWeight.RemainWeight {
					currentProxyWeight.RemainWeight = remainWeight
				}

			}
		} else {
			// 如果ratelimit-remaining为空，则本地维护剩余权重
			currentProxyWeight.RemainWeight -= 1
		}
		//回填是否接口权重已用完
		if currentProxyWeight.RemainWeight <= 0 {
			currentProxyWeight.Is1032Limited = true
		}

		// 回填是否1032限制
		var errRes SunxErrorRes
		if err := json.Unmarshal(data, &errRes); err == nil {
			if errRes.Code == 1032 {
				currentProxyWeight.Is1032Limited = true
				// log.Warn("Proxy 1032 Limited by Body Code")
			}
		}
	}
	return data, err
}
