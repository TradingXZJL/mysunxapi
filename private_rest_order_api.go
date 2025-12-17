package mysunxapi

// POST 下单
func (client *PrivateRestClient) NewPrivateRestTradeOrderPost() *PrivateRestTradeOrderPostAPI {
	return &PrivateRestTradeOrderPostAPI{
		client: client,
		req:    &PrivateRestTradeOrderPostReq{},
	}
}

func (api *PrivateRestTradeOrderPostAPI) Do() (*SunxRestRes[PrivateRestTradeOrderResCommon], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradeOrderPostReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradeOrderPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradeOrderResCommon](url, reqBody, POST)
}

// POST 批量下单
func (client *PrivateRestClient) NewPrivateRestTradeBatchOrders() *PrivateRestTradeBatchOrdersAPI {
	return &PrivateRestTradeBatchOrdersAPI{
		client: client,
		req:    &PrivateRestTradeBatchOrdersReq{},
	}
}

func (api *PrivateRestTradeBatchOrdersAPI) Do() (*SunxRestRes[PrivateRestTradeBatchOrdersRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradeBatchOrdersReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradeBatchOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradeBatchOrdersRes](url, reqBody, POST)
}

// POST 撤单
func (client *PrivateRestClient) NewPrivateRestTradeCancelOrder() *PrivateRestTradeCancelOrderAPI {
	return &PrivateRestTradeCancelOrderAPI{
		client: client,
		req:    &PrivateRestTradeCancelOrderReq{},
	}
}

func (api *PrivateRestTradeCancelOrderAPI) Do() (*SunxRestRes[PrivateRestTradeCancelOrderRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradeCancelOrderReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradeCancelOrder])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradeCancelOrderRes](url, reqBody, POST)
}

// POST 批量撤单
func (client *PrivateRestClient) NewPrivateRestTradeCancelBatchOrders() *PrivateRestTradeCancelBatchOrdersAPI {
	return &PrivateRestTradeCancelBatchOrdersAPI{
		client: client,
		req:    &PrivateRestTradeCancelBatchOrdersReq{},
	}
}

func (api *PrivateRestTradeCancelBatchOrdersAPI) Do() (*SunxRestRes[PrivateRestTradeCancelBatchOrdersRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradeCancelBatchOrdersReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradeCancelBatchOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradeCancelBatchOrdersRes](url, reqBody, POST)
}

// POST 全部撤单
func (client *PrivateRestClient) NewPrivateRestTradeCancelAllOrders() *PrivateRestTradeCancelAllOrdersAPI {
	return &PrivateRestTradeCancelAllOrdersAPI{
		client: client,
		req:    &PrivateRestTradeCancelAllOrdersReq{},
	}
}

func (api *PrivateRestTradeCancelAllOrdersAPI) Do() (*SunxRestRes[PrivateRestTradeCancelAllOrdersRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradeCancelAllOrdersReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradeCancelAllOrders])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradeCancelAllOrdersRes](url, reqBody, POST)
}

// POST 市价全平
func (client *PrivateRestClient) NewPrivateRestTradePosition() *PrivateRestTradePositionAPI {
	return &PrivateRestTradePositionAPI{
		client: client,
		req:    &PrivateRestTradePositionReq{},
	}
}

func (api *PrivateRestTradePositionAPI) Do() (*SunxRestRes[PrivateRestTradePositionRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradePosition])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradePositionRes](url, reqBody, POST)
}

// POST 一键全平
func (client *PrivateRestClient) NewPrivateRestTradePositionAll() *PrivateRestTradePositionAllAPI {
	return &PrivateRestTradePositionAllAPI{
		client: client,
		req:    &PrivateRestTradePositionAllReq{},
	}
}

func (api *PrivateRestTradePositionAllAPI) Do() (*SunxRestRes[PrivateRestTradePositionAllRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionAllReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradePositionAll])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradePositionAllRes](url, reqBody, POST)
}

// GET 查询当前委托
func (client *PrivateRestClient) NewPrivateRestTradeOrderOpens() *PrivateRestTradeOrderOpensAPI {
	return &PrivateRestTradeOrderOpensAPI{
		client: client,
		req:    &PrivateRestTradeOrderOpensReq{},
	}
}

func (api *PrivateRestTradeOrderOpensAPI) Do() (*SunxRestRes[PrivateRestTradeOrderOpensRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradeOrderOpens])
	return sunxCallApi[PrivateRestTradeOrderOpensRes](url, NIL_REQBODY, GET)
}

// GET 查询成交明细
func (client *PrivateRestClient) NewPrivateRestTradeOrderDetails() *PrivateRestTradeOrderDetailsAPI {
	return &PrivateRestTradeOrderDetailsAPI{
		client: client,
		req:    &PrivateRestTradeOrderDetailsReq{},
	}
}

func (api *PrivateRestTradeOrderDetailsAPI) Do() (*SunxRestRes[PrivateRestTradeOrderDetailsRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradeOrderDetails])
	return sunxCallApi[PrivateRestTradeOrderDetailsRes](url, NIL_REQBODY, GET)
}

// GET 查询历史委托
func (client *PrivateRestClient) NewPrivateRestTradeOrderHistory() *PrivateRestTradeOrderHistoryAPI {
	return &PrivateRestTradeOrderHistoryAPI{
		client: client,
		req:    &PrivateRestTradeOrderHistoryReq{},
	}
}

func (api *PrivateRestTradeOrderHistoryAPI) Do() (*SunxRestRes[PrivateRestTradeOrderHistoryRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradeOrderHistory])
	return sunxCallApi[PrivateRestTradeOrderHistoryRes](url, NIL_REQBODY, GET)
}

// GET 查询订单信息
func (client *PrivateRestClient) NewPrivateRestTradeOrderGet() *PrivateRestTradeOrderGetAPI {
	return &PrivateRestTradeOrderGetAPI{
		client: client,
		req:    &PrivateRestTradeOrderGetReq{},
	}
}

func (api *PrivateRestTradeOrderGetAPI) Do() (*SunxRestRes[PrivateRestTradeOrderGetRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradeOrderGet])
	return sunxCallApi[PrivateRestTradeOrderGetRes](url, NIL_REQBODY, GET)
}

// GET 查询用户当前的下单量限制
func (client *PrivateRestClient) NewPrivateRestTradeOrderLimit() *PrivateRestTradeOrderLimitAPI {
	return &PrivateRestTradeOrderLimitAPI{
		client: client,
		req:    &PrivateRestTradeOrderLimitReq{},
	}
}

func (api *PrivateRestTradeOrderLimitAPI) Do() (*SunxRestRes[PrivateRestTradeOrderLimitRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradeOrderLimit])
	return sunxCallApi[PrivateRestTradeOrderLimitRes](url, NIL_REQBODY, GET)
}
