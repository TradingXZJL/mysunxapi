package mysunxapi

type PrivateRestTradeOrderPostReq struct {
	ContractCode       *string  `json:"contract_code"`         // 合约代码 支持大小写，示例："BTC-USDT" ...
	MarginMode         *string  `json:"margin_mode"`           // 保证金模式 cross：全仓 ；
	PositionSide       *string  `json:"position_side"`         // 仓位方向 "long":多 "short":空 “both”:单向持仓，开平模式必填，买卖模式默认为both。
	Side               *string  `json:"side"`                  // 订单方向 "buy":买， "sell":卖
	Type               *string  `json:"type"`                  // 订单类型；枚举 "market": 市价，"limit":限价, "post_only":只做maker
	PriceMatch         *string  `json:"price_match"`           // 最优档位，和price互斥 opponent-对手价、opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
	ClientOrderId      *int64   `json:"client_order_id"`       // 用户自定义订单ID [1-9223372036854775807]
	Price              *float64 `json:"price"`                 // 价格，仅当限价单的时候有效，市价无需输入。
	Volume             *float64 `json:"volume"`                // 委托张数。
	ReduceOnly         *int     `json:"reduce_only"`           // 只减仓 0-否，1-是
	TimeInForce        *string  `json:"time_in_force"`         // 枚举fok, ioc, gtc，非必填，默认是gtc
	TpTriggerPrice     *string  `json:"tp_trigger_price"`      // 止盈触发价格
	TpOrderPrice       *string  `json:"tp_order_price"`        // 止盈委托价格（最优N档委托类型时无需填写价格）
	TpType             *string  `json:"tp_type"`               // 止盈委托类型,不填默认为market；市价：market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	TpTriggerPriceType *string  `json:"tp_trigger_price_type"` // 止盈价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	SlTriggerPrice     *string  `json:"sl_trigger_price"`      // 止损触发价格
	SlTriggerPriceType *string  `json:"sl_trigger_price_type"` // 止损价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	SlOrderPrice       *string  `json:"sl_order_price"`        // 止损委托价格（最优N档委托类型时无需填写价格）
	SlType             *string  `json:"sl_type"`               // 止损委托类型,不填默认为market; 市价:market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	PriceProtect       *string  `json:"price_protect"`         // 价差保护，默认为false。仅当设置止盈/止损需要该参数。 false 或者 true
	SelfMatchPrevent   *string  `json:"self_match_prevent"`    // 防自成交 cancel_maker：撤销maker单 cancel_both：撤销全部订单 默认值：cancel_taker
}
type PrivateRestTradeOrderPostAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderPostReq
}

// contract_code true 合约代码 支持大小写，示例："BTC-USDT" ...
func (api *PrivateRestTradeOrderPostAPI) ContractCode(contractCode string) *PrivateRestTradeOrderPostAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式 cross：全仓 ；
func (api *PrivateRestTradeOrderPostAPI) MarginMode(marginMode string) *PrivateRestTradeOrderPostAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// position_side false 仓位方向 "long":多 "short":空 “both”:单向持仓，开平模式必填，买卖模式默认为both。
func (api *PrivateRestTradeOrderPostAPI) PositionSide(positionSide string) *PrivateRestTradeOrderPostAPI {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// side true 订单方向 "buy":买， "sell":卖
func (api *PrivateRestTradeOrderPostAPI) Side(side string) *PrivateRestTradeOrderPostAPI {
	api.req.Side = GetPointer(side)
	return api
}

// type true 订单类型；枚举 "market": 市价，"limit":限价, "post_only":只做maker
func (api *PrivateRestTradeOrderPostAPI) Type(t string) *PrivateRestTradeOrderPostAPI {
	api.req.Type = GetPointer(t)
	return api
}

// price_match false 最优档位，和price互斥 opponent-对手价、opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
func (api *PrivateRestTradeOrderPostAPI) PriceMatch(priceMatch string) *PrivateRestTradeOrderPostAPI {
	api.req.PriceMatch = GetPointer(priceMatch)
	return api
}

// client_order_id false 用户自定义订单ID [1-9223372036854775807]
func (api *PrivateRestTradeOrderPostAPI) ClientOrderId(clientOrderId int64) *PrivateRestTradeOrderPostAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

// price false 价格，仅当限价单的时候有效，市价无需输入。
func (api *PrivateRestTradeOrderPostAPI) Price(price float64) *PrivateRestTradeOrderPostAPI {
	api.req.Price = GetPointer(price)
	return api
}

// volume true 委托张数。
func (api *PrivateRestTradeOrderPostAPI) Volume(volume float64) *PrivateRestTradeOrderPostAPI {
	api.req.Volume = GetPointer(volume)
	return api
}

// reduce_only false 只减仓 0-否，1-是
func (api *PrivateRestTradeOrderPostAPI) ReduceOnly(reduceOnly int) *PrivateRestTradeOrderPostAPI {
	api.req.ReduceOnly = GetPointer(reduceOnly)
	return api
}

// time_in_force false 枚举fok, ioc, gtc，非必填，默认是gtc
func (api *PrivateRestTradeOrderPostAPI) TimeInForce(timeInForce string) *PrivateRestTradeOrderPostAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// tp_trigger_price false 止盈触发价格
func (api *PrivateRestTradeOrderPostAPI) TpTriggerPrice(tpTriggerPrice string) *PrivateRestTradeOrderPostAPI {
	api.req.TpTriggerPrice = GetPointer(tpTriggerPrice)
	return api
}

// tp_order_price false 止盈委托价格（最优N档委托类型时无需填写价格）
func (api *PrivateRestTradeOrderPostAPI) TpOrderPrice(tpOrderPrice string) *PrivateRestTradeOrderPostAPI {
	api.req.TpOrderPrice = GetPointer(tpOrderPrice)
	return api
}

// tp_type false 止盈委托类型,不填默认为market；市价：market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
func (api *PrivateRestTradeOrderPostAPI) TpType(tpType string) *PrivateRestTradeOrderPostAPI {
	api.req.TpType = GetPointer(tpType)
	return api
}

// tp_trigger_price_type false 止盈价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
func (api *PrivateRestTradeOrderPostAPI) TpTriggerPriceType(tpTriggerPriceType string) *PrivateRestTradeOrderPostAPI {
	api.req.TpTriggerPriceType = GetPointer(tpTriggerPriceType)
	return api
}

// sl_trigger_price false 止损触发价格
func (api *PrivateRestTradeOrderPostAPI) SlTriggerPrice(slTriggerPrice string) *PrivateRestTradeOrderPostAPI {
	api.req.SlTriggerPrice = GetPointer(slTriggerPrice)
	return api
}

// sl_trigger_price_type false 止损价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
func (api *PrivateRestTradeOrderPostAPI) SlTriggerPriceType(slTriggerPriceType string) *PrivateRestTradeOrderPostAPI {
	api.req.SlTriggerPriceType = GetPointer(slTriggerPriceType)
	return api
}

// sl_order_price false 止损委托价格（最优N档委托类型时无需填写价格）
func (api *PrivateRestTradeOrderPostAPI) SlOrderPrice(slOrderPrice string) *PrivateRestTradeOrderPostAPI {
	api.req.SlOrderPrice = GetPointer(slOrderPrice)
	return api
}

// sl_type false 止损委托类型,不填默认为market; 市价:market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
func (api *PrivateRestTradeOrderPostAPI) SlType(slType string) *PrivateRestTradeOrderPostAPI {
	api.req.SlType = GetPointer(slType)
	return api
}

// price_protect false 价差保护，默认为false。仅当设置止盈/止损需要该参数。 false 或者 true
func (api *PrivateRestTradeOrderPostAPI) PriceProtect(priceProtect string) *PrivateRestTradeOrderPostAPI {
	api.req.PriceProtect = GetPointer(priceProtect)
	return api
}

// self_match_prevent false 防自成交 cancel_maker：撤销maker单 cancel_both：撤销全部订单 默认值：cancel_taker
func (api *PrivateRestTradeOrderPostAPI) SelfMatchPrevent(selfMatchPrevent string) *PrivateRestTradeOrderPostAPI {
	api.req.SelfMatchPrevent = GetPointer(selfMatchPrevent)
	return api
}

type PrivateRestTradeCancelOrderReq struct {
	ContractCode  *string `json:"contract_code"`   // 合约代码
	OrderId       *string `json:"order_id"`        // 订单ID order_id和client_order_id必须传一个，若两个都传，以order_id为准
	ClientOrderId *string `json:"client_order_id"` // 用户自定义订单ID order_id和client_order_id必须传一个，若两个都传，以order_id为准
}

type PrivateRestTradeCancelOrderAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelOrderReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeCancelOrderAPI) ContractCode(contractCode string) *PrivateRestTradeCancelOrderAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// order_id false 订单ID order_id和client_order_id必须传一个，若两个都传，以order_id为准
func (api *PrivateRestTradeCancelOrderAPI) OrderId(orderId string) *PrivateRestTradeCancelOrderAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// client_order_id false 用户自定义订单ID order_id和client_order_id必须传一个，若两个都传，以order_id为准
func (api *PrivateRestTradeCancelOrderAPI) ClientOrderIds(clientOrderId string) *PrivateRestTradeCancelOrderAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

type PrivateRestTradeBatchOrdersReq []PrivateRestTradeOrderPostReq
type PrivateRestTradeBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeBatchOrdersReq
}

func (api *PrivateRestTradeBatchOrdersAPI) AddOrder(order PrivateRestTradeOrderPostReq) *PrivateRestTradeBatchOrdersAPI {
	if api.req == nil {
		api.req = &PrivateRestTradeBatchOrdersReq{}
	}
	*api.req = append(*api.req, order)
	return api
}

func (api *PrivateRestTradeBatchOrdersAPI) SetOrderList(orderList []PrivateRestTradeOrderPostReq) *PrivateRestTradeBatchOrdersAPI {
	api.req = &PrivateRestTradeBatchOrdersReq{}
	for _, v := range orderList {
		*api.req = append(*api.req, v)
	}
	return api
}

type PrivateRestTradeCancelBatchOrdersReq struct {
	ContractCode  *string   `json:"contract_code"`   // 合约代码
	OrderId       *[]string `json:"order_id"`        // 订单ID
	ClientOrderId *[]string `json:"client_order_id"` // 用户自定义订单ID
}
type PrivateRestTradeCancelBatchOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelBatchOrdersReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeCancelBatchOrdersAPI) ContractCode(contractCode string) *PrivateRestTradeCancelBatchOrdersAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// order_id false 订单ID列表
func (api *PrivateRestTradeCancelBatchOrdersAPI) OrderId(orderIdList []string) *PrivateRestTradeCancelBatchOrdersAPI {
	api.req.OrderId = GetPointer(orderIdList)
	return api
}

// client_order_id false 用户自定义订单ID列表
func (api *PrivateRestTradeCancelBatchOrdersAPI) ClientOrderId(clientOrderIdList []string) *PrivateRestTradeCancelBatchOrdersAPI {
	api.req.ClientOrderId = GetPointer(clientOrderIdList)
	return api
}

type PrivateRestTradeCancelAllOrdersReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码
	Side         *string `json:"side"`          // 订单方向 buy:买 "sell":卖
	PositionSide *string `json:"position_side"` // 仓位方向 "long":多 "short":空 “both”:单向持仓
}
type PrivateRestTradeCancelAllOrdersAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeCancelAllOrdersReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeCancelAllOrdersAPI) ContractCode(contractCode string) *PrivateRestTradeCancelAllOrdersAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// side true 订单方向 buy:买 "sell":卖
func (api *PrivateRestTradeCancelAllOrdersAPI) Side(side string) *PrivateRestTradeCancelAllOrdersAPI {
	api.req.Side = GetPointer(side)
	return api
}

// position_side true 仓位方向 "long":多 "short":空 “both”:单向持仓
func (api *PrivateRestTradeCancelAllOrdersAPI) PositionSide(positionSide string) *PrivateRestTradeCancelAllOrdersAPI {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

type PrivateRestTradePositionReq struct {
	ContractCode  *string `json:"contract_code"`   // 合约代码 true
	MarginMode    *string `json:"margin_mode"`     // 保证金模式 true cross：全仓 ；
	PositionSide  *string `json:"position_side"`   // 仓位方向 true 持仓方向 买卖模式下默认值both， 开平仓模式下： long：多 ，short：空
	ClientOrderId *string `json:"client_order_id"` // 用户自定义订单ID false [1-9223372036854775807]
}
type PrivateRestTradePositionAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionReq
}

// contract_code true 合约代码
func (api *PrivateRestTradePositionAPI) ContractCode(contractCode string) *PrivateRestTradePositionAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestTradePositionAPI) MarginMode(marginMode string) *PrivateRestTradePositionAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// position_side true 仓位方向
func (api *PrivateRestTradePositionAPI) PositionSide(positionSide string) *PrivateRestTradePositionAPI {
	api.req.PositionSide = GetPointer(positionSide)
	return api
}

// client_order_id false 用户自定义订单ID [1-9223372036854775807]
func (api *PrivateRestTradePositionAPI) ClientOrderId(clientOrderId string) *PrivateRestTradePositionAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

type PrivateRestTradePositionAllReq struct{}
type PrivateRestTradePositionAllAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradePositionAllReq
}

type PrivateRestTradeOrderOpensReq struct {
	ContractCode  *string `json:"contract_code"`   // 合约代码
	MarginMode    *string `json:"margin_mode"`     // 保证金模式 cross：全仓
	OrderId       *string `json:"order_id"`        // 订单ID
	ClientOrderId *string `json:"client_order_id"` // 用户自定义订单ID
	From          *int64  `json:"from"`            // 查询的起始id，默认从0开始
	Limit         *int    `json:"limit"`           // 分页页面大小，默认为10，最大为100。
	Direct        *string `json:"direct"`          // 翻页方向，prev, next，默认next
}
type PrivateRestTradeOrderOpensAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderOpensReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeOrderOpensAPI) ContractCode(contractCode string) *PrivateRestTradeOrderOpensAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestTradeOrderOpensAPI) MarginMode(marginMode string) *PrivateRestTradeOrderOpensAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// order_id false 订单ID
func (api *PrivateRestTradeOrderOpensAPI) OrderId(orderId string) *PrivateRestTradeOrderOpensAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// client_order_id false 用户自定义订单ID
func (api *PrivateRestTradeOrderOpensAPI) ClientOrderId(clientOrderId string) *PrivateRestTradeOrderOpensAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

// from false 查询的起始id，默认从0开始
func (api *PrivateRestTradeOrderOpensAPI) From(from int64) *PrivateRestTradeOrderOpensAPI {
	api.req.From = GetPointer(from)
	return api
}

// limit false 分页页面大小，默认为10，最大为100。
func (api *PrivateRestTradeOrderOpensAPI) Limit(limit int) *PrivateRestTradeOrderOpensAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// direct false 翻页方向，prev, next，默认next
func (api *PrivateRestTradeOrderOpensAPI) Direct(direct string) *PrivateRestTradeOrderOpensAPI {
	api.req.Direct = GetPointer(direct)
	return api
}

type PrivateRestTradeOrderDetailsReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码
	OrderId      *string `json:"order_id"`      // 订单ID
	StartTime    *string `json:"start_time"`    // 查询开始时间，UNIX时间戳，以毫秒为单位。取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值 (now) – 48h
	EndTime      *string `json:"end_time"`      // 查询结束时间，取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值：now
	From         *int64  `json:"from"`          // 查询的起始id，默认从0开始
	Limit        *int    `json:"limit"`         // 分页页面大小，默认为10，最大为100。
	Direct       *string `json:"direct"`        // 翻页方向，prev, next，默认next
}
type PrivateRestTradeOrderDetailsAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderDetailsReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeOrderDetailsAPI) ContractCode(contractCode string) *PrivateRestTradeOrderDetailsAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// order_id false 订单ID
func (api *PrivateRestTradeOrderDetailsAPI) OrderId(orderId string) *PrivateRestTradeOrderDetailsAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// start_time false 查询开始时间，UNIX时间戳，以毫秒为单位。取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值 (now) – 48h
func (api *PrivateRestTradeOrderDetailsAPI) StartTime(startTime string) *PrivateRestTradeOrderDetailsAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// end_time false 查询结束时间，取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值：now
func (api *PrivateRestTradeOrderDetailsAPI) EndTime(endTime string) *PrivateRestTradeOrderDetailsAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// from false 查询的起始id，默认从0开始
func (api *PrivateRestTradeOrderDetailsAPI) From(from int64) *PrivateRestTradeOrderDetailsAPI {
	api.req.From = GetPointer(from)
	return api
}

// limit false 分页页面大小，默认为10，最大为100。
func (api *PrivateRestTradeOrderDetailsAPI) Limit(limit int) *PrivateRestTradeOrderDetailsAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// direct false 翻页方向，prev, next，默认next
func (api *PrivateRestTradeOrderDetailsAPI) Direct(direct string) *PrivateRestTradeOrderDetailsAPI {
	api.req.Direct = GetPointer(direct)
	return api
}

type PrivateRestTradeOrderHistoryReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 true
	MarginMode   *string `json:"margin_mode"`   // 保证金模式 true cross：全仓
	States       *string `json:"states"`        // 订单状态 false 可以查多个状态，使用逗号分隔。“filled“:”已成交”, ”partially_canceled”:””部分成交已撤销, “canceled“:”已撤销”
	Type         *string `json:"type"`          // 订单类型 false "market": 市价，"limit":限价，"post_only":只做maker单
	PriceMatch   *string `json:"price_match"`   // 最优档位 false ,opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
	TimeInForce  *string `json:"time_in_force"` // 生效时间 false 枚举fok, ioc, gtc，非必填，默认是gtc
	StartTime    *string `json:"start_time"`    // 查询开始时间，UNIX时间戳，以毫秒为单位。取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值 (now) – 48h
	EndTime      *string `json:"end_time"`      // 查询结束时间，取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值：now
	From         *int64  `json:"from"`          // 查询的起始id，默认从0开始
	Limit        *int    `json:"limit"`         // 分页页面大小，默认为10，最大为100。
	Direct       *string `json:"direct"`        // 翻页方向，prev, next，默认next
}
type PrivateRestTradeOrderHistoryAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderHistoryReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeOrderHistoryAPI) ContractCode(contractCode string) *PrivateRestTradeOrderHistoryAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestTradeOrderHistoryAPI) MarginMode(marginMode string) *PrivateRestTradeOrderHistoryAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// states false 订单状态 可以查多个状态，使用逗号分隔。“filled“:”已成交”, ”partially_canceled”:””部分成交已撤销, “canceled“:”已撤销”
func (api *PrivateRestTradeOrderHistoryAPI) States(states string) *PrivateRestTradeOrderHistoryAPI {
	api.req.States = GetPointer(states)
	return api
}

// type false 订单类型 "market": 市价，"limit":限价，"post_only":只做maker单
func (api *PrivateRestTradeOrderHistoryAPI) Type(t string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Type = GetPointer(t)
	return api
}

// price_match false 最优档位 opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
func (api *PrivateRestTradeOrderHistoryAPI) PriceMatch(priceMatch string) *PrivateRestTradeOrderHistoryAPI {
	api.req.PriceMatch = GetPointer(priceMatch)
	return api
}

// time_in_force false 生效时间 枚举fok, ioc, gtc，非必填，默认是gtc
func (api *PrivateRestTradeOrderHistoryAPI) TimeInForce(timeInForce string) *PrivateRestTradeOrderHistoryAPI {
	api.req.TimeInForce = GetPointer(timeInForce)
	return api
}

// start_time false 查询开始时间，UNIX时间戳，以毫秒为单位。取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值 (now) – 48h
func (api *PrivateRestTradeOrderHistoryAPI) StartTime(startTime string) *PrivateRestTradeOrderHistoryAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// end_time false 查询结束时间，取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。 默认值：now
func (api *PrivateRestTradeOrderHistoryAPI) EndTime(endTime string) *PrivateRestTradeOrderHistoryAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// from false 查询的起始id，默认从0开始
func (api *PrivateRestTradeOrderHistoryAPI) From(from int64) *PrivateRestTradeOrderHistoryAPI {
	api.req.From = GetPointer(from)
	return api
}

// limit false 分页页面大小，默认为10，最大为100。
func (api *PrivateRestTradeOrderHistoryAPI) Limit(limit int) *PrivateRestTradeOrderHistoryAPI {
	api.req.Limit = GetPointer(limit)
	return api
}

// direct false 翻页方向，prev, next，默认next
func (api *PrivateRestTradeOrderHistoryAPI) Direct(direct string) *PrivateRestTradeOrderHistoryAPI {
	api.req.Direct = GetPointer(direct)
	return api
}

type PrivateRestTradeOrderGetReq struct {
	ContractCode  *string `json:"contract_code"`   // 合约代码 true
	MarginMode    *string `json:"margin_mode"`     // 保证金模式 true cross：全仓
	OrderId       *string `json:"order_id"`        // 订单id，order_id和client_order_id必填写一个，若都填写了以order_id为准
	ClientOrderId *string `json:"client_order_id"` // 用户自定义订单ID
}
type PrivateRestTradeOrderGetAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderGetReq
}

// contract_code true 合约代码
func (api *PrivateRestTradeOrderGetAPI) ContractCode(contractCode string) *PrivateRestTradeOrderGetAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// margin_mode true 保证金模式
func (api *PrivateRestTradeOrderGetAPI) MarginMode(marginMode string) *PrivateRestTradeOrderGetAPI {
	api.req.MarginMode = GetPointer(marginMode)
	return api
}

// order_id false 订单id，order_id和client_order_id必填写一个，若都填写了以order_id为准
func (api *PrivateRestTradeOrderGetAPI) OrderId(orderId string) *PrivateRestTradeOrderGetAPI {
	api.req.OrderId = GetPointer(orderId)
	return api
}

// client_order_id false 用户自定义订单ID
func (api *PrivateRestTradeOrderGetAPI) ClientOrderId(clientOrderId string) *PrivateRestTradeOrderGetAPI {
	api.req.ClientOrderId = GetPointer(clientOrderId)
	return api
}

type PrivateRestTradeOrderLimitReq struct {
	ContractCode   *string `json:"contract_code"`    // 合约代码 false
	OrderPriceType *string `json:"order_price_type"` // 订单报价类型 true limit:限价，"opponent":对手价，"lightning":闪电平仓，"optimal_5":最优5档，"optimal_10":最优10档，"optimal_20":最优20档，"fok":FOK订单，"ioc":IOC订单,opponent_ioc"： 对手价-IOC下单，"lightning_ioc"：闪电平仓-IOC下单，"optimal_5_ioc"：最优5档-IOC下单，"optimal_10_ioc"：最优10档-IOC下单，"optimal_20_ioc"：最优20档-IOC下单,"opponent_fok"： 对手价-FOK下单，"lightning_fok"：闪电平仓-FOK下单，"optimal_5_fok"：最优5档-FOK下单，"optimal_10_fok"：最优10档-FOK下单，"optimal_20_fok"：最优20档-FOK下单
	Pair           *string `json:"pair"`             // 合约对 true 例如：BTC-USDT
	ContractType   *string `json:"contract_type"`    // 合约类型 false swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	BusinessType   *string `json:"business_type"`    // 业务类型 false futures：交割、swap：永续、all：全部
}
type PrivateRestTradeOrderLimitAPI struct {
	client *PrivateRestClient
	req    *PrivateRestTradeOrderLimitReq
}

// contract_code false 合约代码
func (api *PrivateRestTradeOrderLimitAPI) ContractCode(contractCode string) *PrivateRestTradeOrderLimitAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// order_price_type true 订单报价类型 limit:限价，"opponent":对手价，"lightning":闪电平仓，"optimal_5":最优5档，"optimal_10":最优10档，"optimal_20":最优20档，"fok":FOK订单，"ioc":IOC订单,opponent_ioc"： 对手价-IOC下单，"lightning_ioc"：闪电平仓-IOC下单，"optimal_5_ioc"：最优5档-IOC下单，"optimal_10_ioc"：最优10档-IOC下单，"optimal_20_ioc"：最优20档-IOC下单,"opponent_fok"： 对手价-FOK下单，"lightning_fok"：闪电平仓-FOK下单，"optimal_5_fok"：最优5档-FOK下单，"optimal_10_fok"：最优10档-FOK下单，"optimal_20_fok"：最优20档-FOK下单
func (api *PrivateRestTradeOrderLimitAPI) OrderPriceType(orderPriceType string) *PrivateRestTradeOrderLimitAPI {
	api.req.OrderPriceType = GetPointer(orderPriceType)
	return api
}

// pair true 合约对 例如：BTC-USDT
func (api *PrivateRestTradeOrderLimitAPI) Pair(pair string) *PrivateRestTradeOrderLimitAPI {
	api.req.Pair = GetPointer(pair)
	return api
}

// contract_type false 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
func (api *PrivateRestTradeOrderLimitAPI) ContractType(contractType string) *PrivateRestTradeOrderLimitAPI {
	api.req.ContractType = GetPointer(contractType)
	return api
}

// business_type false 业务类型 futures：交割、swap：永续、all：全部
func (api *PrivateRestTradeOrderLimitAPI) BusinessType(businessType string) *PrivateRestTradeOrderLimitAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}
