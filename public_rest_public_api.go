package mysunxapi

// GET PUBLIC_READ 获取合约信息
func (client *PublicRestClient) NewPublicRestPublicContractInfo() *PublicRestPublicContractInfoAPI {
	return &PublicRestPublicContractInfoAPI{
		client: client,
		req:    &PublicRestPublicContractInfoReq{},
	}
}

func (api *PublicRestPublicContractInfoAPI) Do() (*SunxRestRes[PublicRestPublicContractInfoRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestPublicContractInfo])
	return sunxCallApi[PublicRestPublicContractInfoRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PUBLIC_READ 获取合约指数信息
func (client *PublicRestClient) NewPublicRestPublicIndex() *PublicRestPublicIndexAPI {
	return &PublicRestPublicIndexAPI{
		client: client,
		req:    &PublicRestPublicIndexReq{},
	}
}

func (api *PublicRestPublicIndexAPI) Do() (*SunxRestRes[PublicRestPublicIndexRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestPublicIndex])
	return sunxCallApi[PublicRestPublicIndexRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PRIVATE_READ 获取合约最高限价和最低限价
func (client *PublicRestClient) NewPublicRestPublicRiskLimit() *PublicRestPublicRiskLimitAPI {
	return &PublicRestPublicRiskLimitAPI{
		client: client,
		req:    &PublicRestPublicRiskLimitReq{},
	}
}

func (api *PublicRestPublicRiskLimitAPI) Do() (*SunxRestRes[PublicRestPublicRiskLimitRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PRIVATE_READ, api.req, PublicRestAPIMap[PublicRestPublicRiskLimit])
	return sunxCallApi[PublicRestPublicRiskLimitRes](url, NIL_REQBODY, GET, PRIVATE_READ)
}

// GET PUBLIC_READ 获取合约的资金费率
func (client *PublicRestClient) NewPublicRestPublicFundingRate() *PublicRestPublicFundingRateAPI {
	return &PublicRestPublicFundingRateAPI{
		client: client,
		req:    &PublicRestPublicFundingRateReq{},
	}
}

func (api *PublicRestPublicFundingRateAPI) Do() (*SunxRestRes[PublicRestPublicFundingRateRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestPublicFundingRate])
	return sunxCallApi[PublicRestPublicFundingRateRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PUBLIC_READ 获取合约的历史资金费率
func (client *PublicRestClient) NewPublicRestPublicFundingRateHistory() *PublicRestPublicFundingRateHistoryAPI {
	return &PublicRestPublicFundingRateHistoryAPI{
		client: client,
		req:    &PublicRestPublicFundingRateHistoryReq{},
	}
}

// GET PUBLIC_READ 获取合约最高限价和最低限价
func (client *PublicRestClient) NewPublicRestPublicPriceLimit() *PublicRestPublicPriceLimitAPI {
	return &PublicRestPublicPriceLimitAPI{
		client: client,
		req:    &PublicRestPublicPriceLimitReq{},
	}
}

func (api *PublicRestPublicPriceLimitAPI) Do() (*SunxRestRes[PublicRestPublicPriceLimitRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestPublicPriceLimit])
	return sunxCallApi[PublicRestPublicPriceLimitRes](url, NIL_REQBODY, GET, PUBLIC)
}

func (api *PublicRestPublicFundingRateHistoryAPI) Do() (*SunxRestRes[PublicRestPublicFundingRateHistoryRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PUBLIC, api.req, PublicRestAPIMap[PublicRestPublicFundingRateHistory])
	return sunxCallApi[PublicRestPublicFundingRateHistoryRes](url, NIL_REQBODY, GET, PUBLIC)
}

// GET PRIVATE_READ 查询联合保证金支持币种
func (client *PublicRestClient) NewPublicRestPublicMultiAssetsMargin() *PublicRestPublicMultiAssetsMarginAPI {
	return &PublicRestPublicMultiAssetsMarginAPI{
		client: client,
		req:    &PublicRestPublicMultiAssetsMarginReq{},
	}
}

func (api *PublicRestPublicMultiAssetsMarginAPI) Do() (*SunxRestRes[PublicRestPublicMultiAssetsMarginRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(PRIVATE_READ, api.req, PublicRestAPIMap[PublicRestPublicMultiAssetsMargin])
	return sunxCallApi[PublicRestPublicMultiAssetsMarginRes](url, NIL_REQBODY, GET, PRIVATE_READ)
}
