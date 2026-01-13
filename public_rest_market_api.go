package mysunxapi

// GET PUBLIC_READ 获取行情深度数据
func (client *PublicRestClient) NewPublicRestMarketDepth() *PublicRestMarketDepthAPI {
	return &PublicRestMarketDepthAPI{
		client: client,
		req:    &PublicRestMarketDepthReq{},
	}
}

func (api *PublicRestMarketDepthAPI) Do() (*SunxRestRes[PublicRestMarketDepthRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketDepth])
	middleRes, err := sunxCallApi[PublicRestMarketDepthResMiddle](url, NIL_REQBODY, GET, PUBLIC)
	if err != nil {
		return nil, err
	}
	res := &SunxRestRes[PublicRestMarketDepthRes]{
		SunxErrorRes: middleRes.SunxErrorRes,
		SunxTimeRes:  middleRes.SunxTimeRes,
		Ch:           middleRes.Ch,
		Data:         *middleRes.Data.ConvertToRes(),
	}
	return res, nil
}

// GET PUBLIC_READ 获取K线数据
func (client *PublicRestClient) NewPublicRestMarketHistoryKline() *PublicRestMarketHistoryKlineAPI {
	return &PublicRestMarketHistoryKlineAPI{
		client: client,
		req:    &PublicRestMarketHistoryKlineReq{},
	}
}

func (api *PublicRestMarketHistoryKlineAPI) Do() (*SunxRestRes[PublicRestMarketHistoryKlineRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketHistoryKline])
	return sunxCallApi[PublicRestMarketHistoryKlineRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PUBLIC_READ 获取聚合行情
func (client *PublicRestClient) NewPublicRestMarketDetailMerged() *PublicRestMarketDetailMergedAPI {
	return &PublicRestMarketDetailMergedAPI{
		client: client,
		req:    &PublicRestMarketDetailMergedReq{},
	}
}

func (api *PublicRestMarketDetailMergedAPI) Do() (*SunxRestRes[PublicRestMarketDetailMergedRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketDetailMerged])
	middleRes, err := sunxCallApi[PublicRestMarketDetailMergedResMiddle](url, NIL_REQBODY, GET, PUBLIC)
	if err != nil {
		return nil, err
	}
	res := &SunxRestRes[PublicRestMarketDetailMergedRes]{
		SunxErrorRes: middleRes.SunxErrorRes,
		SunxTimeRes:  middleRes.SunxTimeRes,
		Ch:           middleRes.Ch,
		Data:         *middleRes.Data.ConvertToRes(),
	}
	return res, nil
}

// GET PUBLIC_READ 获取最新成交
func (client *PublicRestClient) NewPublicRestMarketTrade() *PublicRestMarketTradeAPI {
	return &PublicRestMarketTradeAPI{
		client: client,
		req:    &PublicRestMarketTradeReq{},
	}
}

func (api *PublicRestMarketTradeAPI) Do() (*SunxRestRes[PublicRestMarketTradeRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketTrade])
	return sunxCallApi[PublicRestMarketTradeRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PUBLIC_READ 获取市场最优挂单
func (client *PublicRestClient) NewPublicRestMarketBBO() *PublicRestMarketBBOAPI {
	return &PublicRestMarketBBOAPI{
		client: client,
		req:    &PublicRestMarketBBOReq{},
	}
}

func (api *PublicRestMarketBBOAPI) Do() (*SunxRestRes[PublicRestMarketBBORes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketBBO])
	middleRes, err := sunxCallApi[PublicRestMarketBBOResMiddle](url, NIL_REQBODY, GET, PUBLIC)
	if err != nil {
		return nil, err
	}
	res := &SunxRestRes[PublicRestMarketBBORes]{
		SunxErrorRes: middleRes.SunxErrorRes,
		SunxTimeRes:  middleRes.SunxTimeRes,
		Ch:           middleRes.Ch,
		Data:         *middleRes.Data.ConvertToRes(),
	}
	return res, nil
}

// GET PUBLIC_READ 批量获取最近的交易记录
func (client *PublicRestClient) NewPublicRestMarketHistoryTrade() *PublicRestMarketHistoryTradeAPI {
	return &PublicRestMarketHistoryTradeAPI{
		client: client,
		req:    &PublicRestMarketHistoryTradeReq{},
	}
}

func (api *PublicRestMarketHistoryTradeAPI) Do() (*SunxRestRes[PublicRestMarketHistoryTradeRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestMarketHistoryTrade])
	return sunxCallApi[PublicRestMarketHistoryTradeRes](url, NIL_REQBODY, GET, PUBLIC)
}
