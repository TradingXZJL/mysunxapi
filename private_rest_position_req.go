package mysunxapi

type PrivateRestTradePositionOpensReq struct {
	ContractCode *string `json:"contract_code"`
}
type PrivateRestTradePositionOpensAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionOpensReq
}

// contract_code string false 合约代码 永续："BTC-USDT"... ， 交割：“BTC-USDT-210625” ...
func (api *PrivateRestTradePositionOpensAPI) ContractCode(contractCode string) *PrivateRestTradePositionOpensAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

type PrivateRestTradePositionLeverGetReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 false
	MarginMode   *string `json:"margin_mode"`   // 保证金模式 true cross：全仓
}
type PrivateRestPositionLeverGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionLeverGetReq
}

// contract_code false 合约代码
func (api *PrivateRestPositionLeverGetAPI) ContractCode(contractCode string) *PrivateRestPositionLeverGetAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestPositionLeverGetAPI) MarginMode(marginMode string) *PrivateRestPositionLeverGetAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

type PrivateRestPositionLeverPostReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 true
	MarginMode   *string `json:"margin_mode"`   // 保证金模式 true cross：全仓
	LeverRate    *string `json:"lever_rate"`    // 杠杆等级 true
}
type PrivateRestPositionLeverPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestPositionLeverPostReq
}

// contract_code true 合约代码
func (api *PrivateRestPositionLeverPostAPI) ContractCode(contractCode string) *PrivateRestPositionLeverPostAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestPositionLeverPostAPI) MarginMode(marginMode string) *PrivateRestPositionLeverPostAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// lever_rate true 杠杆等级
func (api *PrivateRestPositionLeverPostAPI) LeverRate(leverRate string) *PrivateRestPositionLeverPostAPI {
	api.req.LeverRate = GetPointer(leverRate)
	return api
}

type PrivateRestPositionModeGetReq struct{}
type PrivateRestPositionModeGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestPositionModeGetReq
}

type PrivateRestPositionModePostReq struct {
	PositionMode *string `json:"position_mode"` // 持仓模式 true single_side：单向持仓；dual_side：双向持仓
}
type PrivateRestPositionModePostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestPositionModePostReq
}

// position_mode true 持仓模式
func (api *PrivateRestPositionModePostAPI) PositionMode(positionMode string) *PrivateRestPositionModePostAPI {
	api.req.PositionMode = GetPointer(positionMode)
	return api
}

type PrivateRestTradePositionRiskLimitReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 false
	MarginMode   *string `json:"margin_mode"`   // 保证金模式 false cross：全仓
	PositionSide *string `json:"position_side"` // 仓位方向 false long:多，short:空，both：单向持仓
}
type PrivateRestPositionRiskLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionRiskLimitReq
}

// contract_code false 合约代码
func (api *PrivateRestPositionRiskLimitAPI) ContractCode(contractCode string) *PrivateRestPositionRiskLimitAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode false 保证金模式
func (api *PrivateRestPositionRiskLimitAPI) MarginMode(marginMode string) *PrivateRestPositionRiskLimitAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// position_side false 仓位方向
func (api *PrivateRestPositionRiskLimitAPI) PositionSide(positionSide string) *PrivateRestPositionRiskLimitAPI {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

type PrivateRestTradePositionPositionLimitReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 false 永续："BTC-USDT"... ，交割：”BTC-USDT-210625“
	Pair         *string `json:"pair"`          // 合约对 false BTC-USDT
	ContractType *string `json:"contract_type"` // 合约类型 false swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
}
type PrivateRestPositionPositionLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionPositionLimitReq
}

// contract_code false 合约代码
func (api *PrivateRestPositionPositionLimitAPI) ContractCode(contractCode string) *PrivateRestPositionPositionLimitAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// pair false 合约对
func (api *PrivateRestPositionPositionLimitAPI) Pair(pair string) *PrivateRestPositionPositionLimitAPI {
	api.req.Pair = GetPointer(pair)
	return api
}

// contract_type false 合约类型
func (api *PrivateRestPositionPositionLimitAPI) ContractType(contractType string) *PrivateRestPositionPositionLimitAPI {
	api.req.ContractType = GetPointer(contractType)
	return api
}
