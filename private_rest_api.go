package mysunxapi

type PrivateRestAPI int

const (
	// Account
	PrivateRestAccountBalance    PrivateRestAPI = iota // PRIVATE_READ 查询账户余额
	PrivateRestAccountBillRecord                       // PRIVATE_READ 组合查询用户财务记录
	PrivateRestAccountFeeRate                          // PRIVATE_READ 查询用户当前的手续费费率

	// Order
	PrivateRestTradeOrderPost         // PRIVATE_TRADE 下单
	PrivateRestTradeBatchOrders       // PRIVATE_TRADE 批量下单
	PrivateRestTradeCancelOrder       // PRIVATE_TRADE 撤单
	PrivateRestTradeCancelBatchOrders // PRIVATE_TRADE 批量撤单
	PrivateRestTradeCancelAllOrders   // PRIVATE_TRADE 全部撤单
	PrivateRestTradePosition          // PRIVATE_TRADE 市价全平
	PrivateRestTradePositionAll       // PRIVATE_TRADE 一键全平
	PrivateRestTradeOrderOpens        // PRIVATE_READ 查询当前委托
	PrivateRestTradeOrderDetails      // PRIVATE_READ 查询成交明细
	PrivateRestTradeOrderHistory      // PRIVATE_READ 查询历史委托
	PrivateRestTradeOrderGet          // PRIVATE_READ 查询订单信息
	PrivateRestTradeOrderLimit        // PRIVATE_READ 查询用户当前的下单量限制

	// Position
	PrivateRestTradePositionOpens    // PRIVATE_READ 查询当前持仓
	PrivateRestPositionLeverGet      // PRIVATE_READ 查询杠杆等级列表
	PrivateRestPositionLeverPost     // PRIVATE_TRADE 设置杠杆等级
	PrivateRestPositionModeGet       // PRIVATE_READ 查询持仓模式
	PrivateRestPositionModePost      // PRIVATE_TRADE 设置持仓模式
	PrivateRestPositionRiskLimit     // PRIVATE_READ 查询持仓风险限额
	PrivateRestPositionPositionLimit // PRIVATE_READ 用户持仓量限制的查询
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	// Account
	PrivateRestAccountBalance:    "/sapi/v1/account/balance",     // GET PRIVATE_READ 查询账户余额
	PrivateRestAccountBillRecord: "/sapi/v1/account/bill_record", // POST PRIVATE_READ 组合查询用户财务记录
	PrivateRestAccountFeeRate:    "/sapi/v1/account/fee_rate",    // POST PRIVATE_READ 查询用户当前的手续费费率

	// Order
	PrivateRestTradeOrderPost:         "/sapi/v1/trade/order",               // POST PRIVATE_TRADE 下单
	PrivateRestTradeBatchOrders:       "/sapi/v1/trade/batch_orders",        // POST PRIVATE_TRADE 批量下单
	PrivateRestTradeCancelOrder:       "/sapi/v1/trade/cancel_order",        // POST PRIVATE_TRADE 撤单
	PrivateRestTradeCancelBatchOrders: "/sapi/v1/trade/cancel_batch_orders", // POST PRIVATE_TRADE 批量撤单
	PrivateRestTradeCancelAllOrders:   "/sapi/v1/trade/cancel_all_orders",   // POST PRIVATE_TRADE 全部撤单
	PrivateRestTradePosition:          "/sapi/v1/trade/position",            // POST PRIVATE_TRADE 市价全平
	PrivateRestTradePositionAll:       "/sapi/v1/trade/position_all",        // POST PRIVATE_TRADE 一键全平
	PrivateRestTradeOrderOpens:        "/sapi/v1/trade/order/opens",         // GET PRIVATE_READ 查询当前委托
	PrivateRestTradeOrderDetails:      "/sapi/v1/trade/order/details",       // GET PRIVATE_READ 查询成交明细
	PrivateRestTradeOrderHistory:      "/sapi/v1/trade/order/history",       // GET PRIVATE_READ 查询历史委托
	PrivateRestTradeOrderGet:          "/sapi/v1/trade/order",               // GET PRIVATE_READ 查询订单信息
	PrivateRestTradeOrderLimit:        "/sapi/v1/trade/order_limit",         // GET PRIVATE_READ 查询用户当前的下单量限制

	// Positon
	PrivateRestTradePositionOpens:    "/sapi/v1/trade/position/opens",    // GET PRIVATE_READ 查询当前持仓
	PrivateRestPositionLeverGet:      "/sapi/v1/position/lever",          // GET PRIVATE_READ 查询杠杆等级列表
	PrivateRestPositionLeverPost:     "/sapi/v1/position/lever",          // POST PRIVATE_TRADE 设置杠杆等级
	PrivateRestPositionModeGet:       "/sapi/v1/position/mode",           // GET PRIVATE_READ 查询持仓模式
	PrivateRestPositionModePost:      "/sapi/v1/position/mode",           // POST PRIVATE_TRADE 设置持仓模式
	PrivateRestPositionRiskLimit:     "/sapi/v1/position/risk/limit",     // GET PRIVATE_READ 查询持仓风险限额
	PrivateRestPositionPositionLimit: "/sapi/v1/position/position_limit", // POST PRIVATE_READ 用户持仓量限制的查询
}
