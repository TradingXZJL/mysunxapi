package mysunxapi

// GET 查询当前持仓
func (client *PrivateRestClient) NewPrivateRestTradePositionOpens() *PrivateRestTradePositionOpensAPI {
	return &PrivateRestTradePositionOpensAPI{
		client: client,
		req:    &PrivateRestTradePositionOpensReq{},
	}
}

func (api *PrivateRestTradePositionOpensAPI) Do() (*SunxRestRes[PrivateRestTradePositionOpensRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradePositionOpens])
	return sunxCallApi[PrivateRestTradePositionOpensRes](url, NIL_REQBODY, GET)
}

// GET 查询杠杆等级列表
func (client *PrivateRestClient) NewPrivateRestPositionLeverGet() *PrivateRestPositionLeverGetAPI {
	return &PrivateRestPositionLeverGetAPI{
		client: client,
		req:    &PrivateRestTradePositionLeverGetReq{},
	}
}

func (api *PrivateRestPositionLeverGetAPI) Do() (*SunxRestRes[PrivateRestPositionLeverGetRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestPositionLeverGet])
	return sunxCallApi[PrivateRestPositionLeverGetRes](url, NIL_REQBODY, GET)
}

// POST 设置杠杆等级
func (client *PrivateRestClient) NewPrivateRestPositionLeverPost() *PrivateRestPositionLeverPostAPI {
	return &PrivateRestPositionLeverPostAPI{
		client: client,
		req:    &PrivateRestPositionLeverPostReq{},
	}
}

func (api *PrivateRestPositionLeverPostAPI) Do() (*SunxRestRes[PrivateRestPositionLeverPostRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestPositionLeverPostReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestPositionLeverPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestPositionLeverPostRes](url, reqBody, POST)
}

// GET 查询持仓模式
func (client *PrivateRestClient) NewPrivateRestPositionModeGet() *PrivateRestPositionModeGetAPI {
	return &PrivateRestPositionModeGetAPI{
		client: client,
		req:    &PrivateRestPositionModeGetReq{},
	}
}

func (api *PrivateRestPositionModeGetAPI) Do() (*SunxRestRes[PrivateRestPositionModeGetRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestPositionModeGet])
	return sunxCallApi[PrivateRestPositionModeGetRes](url, NIL_REQBODY, GET)
}

// POST 设置持仓模式
func (client *PrivateRestClient) NewPrivateRestPositionModePost() *PrivateRestPositionModePostAPI {
	return &PrivateRestPositionModePostAPI{
		client: client,
		req:    &PrivateRestPositionModePostReq{},
	}
}

func (api *PrivateRestPositionModePostAPI) Do() (*SunxRestRes[PrivateRestPositionModePostRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestPositionModePostReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestPositionModePost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestPositionModePostRes](url, reqBody, POST)
}

// GET 查询持仓风险限额
func (client *PrivateRestClient) NewPrivateRestPositionRiskLimit() *PrivateRestPositionRiskLimitAPI {
	return &PrivateRestPositionRiskLimitAPI{
		client: client,
		req:    &PrivateRestTradePositionRiskLimitReq{},
	}
}

func (api *PrivateRestPositionRiskLimitAPI) Do() (*SunxRestRes[PrivateRestPositionRiskLimitRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestPositionRiskLimit])
	return sunxCallApi[PrivateRestPositionRiskLimitRes](url, NIL_REQBODY, GET)
}

// POST 用户持仓量限制的查询
func (client *PrivateRestClient) NewPrivateRestPositionPositionLimit() *PrivateRestPositionPositionLimitAPI {
	return &PrivateRestPositionPositionLimitAPI{
		client: client,
		req:    &PrivateRestTradePositionPositionLimitReq{},
	}
}

func (api *PrivateRestPositionPositionLimitAPI) Do() (*SunxRestRes[PrivateRestPositionPositionLimitRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionPositionLimitReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestPositionPositionLimit])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestPositionPositionLimitRes](url, reqBody, POST)
}
