package mysunxapi

type PrivateRestTradeOrderResCommon struct {
	OrderId       string `json:"order_id"`        // 订单ID
	ClientOrderId string `json:"client_order_id"` // 用户下单时填写的客户端订单ID，没填则不返回
}

type PrivateRestTradeCancelOrderRes PrivateRestTradeOrderResCommon

type PrivateRestTradeBatchOrdersResCommon struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	PrivateRestTradeOrderResCommon
}
type PrivateRestTradeBatchOrdersRes []PrivateRestTradeBatchOrdersResCommon

type PrivateRestTradeCancelBatchOrdersRes []PrivateRestTradeBatchOrdersResCommon

type PrivateRestTradeCancelAllOrdersRes PrivateRestTradeCancelBatchOrdersRes

type PrivateRestTradePositionRes struct {
	OrderId       int64  `json:"order_id"`        // 订单ID
	ClientOrderId string `json:"client_order_id"` // 用户下单时填写的客户端订单ID，没填则不返回
}

type PrivateRestTradePositionAllRes []PrivateRestTradeBatchOrdersResCommon

type PrivateRestTradeOrderOpensResRow struct {
	ContractCode       string `json:"contract_code"`         // 合约代码 支持大小写，示例："BTC-USDT" ...
	Side               string `json:"side"`                  // 订单方向 "buy":买， "sell":卖
	PositionSide       string `json:"position_side"`         // 仓位方向 "long":多 "short":空 “both”:单向持仓，开平模式必填，买卖模式默认为both。
	Type               string `json:"type"`                  // 订单类型；枚举 "market": 市价，"limit":限价, "post_only":只做maker
	PriceMatch         string `json:"price_match"`           // 最优档位，和price互斥 opponent-对手价、opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
	OrderId            string `json:"order_id"`              // 订单ID
	ClientOrderId      string `json:"client_order_id"`       // 用户自定义订单ID [1-9223372036854775807]
	MarginMode         string `json:"margin_mode"`           // 保证金模式 cross：全仓
	Price              string `json:"price"`                 // 价格，仅当限价单的时候有效，市价无需输入。
	Volume             string `json:"volume"`                // 委托张数。
	LeverRate          int64  `json:"lever_rate"`            // 杠杆等级
	State              string `json:"state"`                 // 订单状态 new, partially_filled, filled, partially_canceled, canceled,rejected
	OrderSource        string `json:"order_source"`          // 订单来源 api：接口下单, （system:系统、web:用户网页、api:用户API、m:用户M站、risk:风控系统、settlement:交割结算、ios：ios客户端、android：安卓客户端、windows：windows客户端、mac：mac客户端、trigger：计划委托触发、tpsl:止盈止损触发、ADL: adl订单）
	ReduceOnly         bool   `json:"reduce_only"`           // 只减仓
	TimeInForce        string `json:"time_in_force"`         // 枚举fok, ioc, gtc，非必填，默认是gtc
	TpTriggerPrice     string `json:"tp_trigger_price"`      // 止盈触发价格
	TpOrderPrice       string `json:"tp_order_price"`        // 止盈委托价格（最优N档委托类型时无需填写价格）
	TpType             string `json:"tp_type"`               // 止盈委托类型,不填默认为market；市价：market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	TpTriggerPriceType string `json:"tp_trigger_price_type"` // 止盈价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	SlTriggerPrice     string `json:"sl_trigger_price"`      // 止损触发价格
	SlOrderPrice       string `json:"sl_order_price"`        // 止损委托价格（最优N档委托类型时无需填写价格）
	SlType             string `json:"sl_type"`               // 止损委托类型,不填默认为market; 市价:market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	SlTriggerPriceType string `json:"sl_trigger_price_type"` // 止损价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	TradeAvgPrice      string `json:"trade_avg_price"`       // 成交均价
	TradeVolume        string `json:"trade_volume"`          // 成交数量
	TradeTurnover      string `json:"trade_turnover"`        // 成交总金额
	FeeCurrency        string `json:"fee_currency"`          // 手续费币种，多个则使用","分隔
	Fee                string `json:"fee"`                   // 总手续费(U)
	PriceProtect       bool   `json:"price_protect"`         // 价差保护，默认为false。仅当设置止盈/止损需要该参数。 false 或者 true
	Profit             string `json:"profit"`                // 平仓盈亏
	ContractType       string `json:"contract_type"`         // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	CreatedTime        string `json:"created_time"`          // 订单创建时间, UTC时间戳(MS)
	UpdatedTime        string `json:"updated_time"`          // 订单更新时间, UTC时间戳(MS)
	SelfMatchPrevent   string `json:"self_match_prevent"`    // 防自成交 cancel_maker：撤销maker单 cancel_both：撤销全部订单 默认值：cancel_taker
}
type PrivateRestTradeOrderOpensRes []PrivateRestTradeOrderOpensResRow

type PrivateRestTradeOrderDetailsResRow struct {
	Id            string `json:"id"`             // 查询id
	ContractCode  string `json:"contract_code"`  // 合约代码 支持大小写，示例："BTC-USDT" ...
	OrderId       string `json:"order_id"`       // 订单ID
	TradeId       string `json:"trade_id"`       // 成交id
	Side          string `json:"side"`           // 订单方向 "buy":买， "sell":卖
	PositionSide  string `json:"position_side"`  // 仓位方向 "long":多 "short":空 “both”:单向持仓，开平模式必填，买卖模式默认为both。
	OrderType     string `json:"order_type"`     // 成交类型 "1":"trade", "3":"liquidation" ,"4":"delivery","22":"adl"
	MarginMode    string `json:"margin_mode"`    // 保证金模式 cross：全仓
	Type          string `json:"type"`           // 订单类型；枚举 "market": 市价，"limit":限价, "post_only":只做maker
	Role          string `json:"role"`           // 订单角色 TAKER MAKER
	TradePrice    string `json:"trade_price"`    // 成交价格
	TradeVolume   string `json:"trade_volume"`   // 成交数量
	TradeTurnover string `json:"trade_turnover"` // 成交总金额
	CreatedTime   string `json:"created_time"`   // 订单创建时间, UTC时间戳(MS)
	UpdatedTime   string `json:"updated_time"`   // 订单更新时间, UTC时间戳(MS)
	OrderSource   string `json:"order_source"`   // 订单来源 api：接口下单, （system:系统、web:用户网页、api:用户API、m:用户M站、risk:风控系统、settlement:交割结算、ios：ios客户端、android：安卓客户端、windows：windows客户端、mac：mac客户端、trigger：计划委托触发、tpsl:止盈止损触发、ADL: adl订单）
	FeeCurrency   string `json:"fee_currency"`   // 手续费币种，多个则使用","分隔
	TradeFee      string `json:"trade_fee"`      // 总手续费(U)
	Profit        string `json:"profit"`         // 平仓盈亏
	ContractType  string `json:"contract_type"`  // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
}
type PrivateRestTradeOrderDetailsRes []PrivateRestTradeOrderDetailsResRow

type PrivateRestTradeOrderHistoryResRow struct {
	Id                 string `json:"id"`                    // 查询id
	ContractCode       string `json:"contract_code"`         // 合约代码
	Side               string `json:"side"`                  // 订单方向 "buy":买， "sell":卖
	PositionSide       string `json:"position_side"`         // 仓位方向 "long":多 "short":空 “both”:单向持仓，开平模式必填，买卖模式默认为both。
	Type               string `json:"type"`                  // 订单类型 "market": 市价，"limit":限价，"post_only":只做maker单
	PriceMatch         string `json:"price_match"`           // 最优档位 opponent-对手价、"optimal_5"：最优5档，"optimal_10"：最优10档，"optimal_20"：最优20档
	OrderId            string `json:"order_id"`              // 订单ID
	ClientOrderId      string `json:"client_order_id"`       // 用户自定义订单ID
	MarginMode         string `json:"margin_mode"`           // 保证金模式 cross：全仓
	Price              string `json:"price"`                 // 价格
	Volume             string `json:"volume"`                // 委托张数
	LeverRate          int64  `json:"lever_rate"`            // 杠杆等级
	State              string `json:"state"`                 // 订单状态 "filled":已成交, "partially_canceled":部分成交已撤销, "canceled":已撤销
	OrderSource        string `json:"order_source"`          // 订单来源 api：接口下单, （system:系统、web:用户网页、api:用户API、m:用户M站、risk:风控系统、settlement:交割结算、ios：ios客户端、android：安卓客户端、windows：windows客户端、mac：mac客户端、trigger：计划委托触发、tpsl:止盈止损触发、ADL: adl订单）
	ReduceOnly         bool   `json:"reduce_only"`           // 只减仓
	TimeInForce        string `json:"time_in_force"`         // 枚举fok, ioc, gtc，非必填，默认是gtc
	TpTriggerPrice     string `json:"tp_trigger_price"`      // 止盈触发价格
	TpOrderPrice       string `json:"tp_order_price"`        // 止盈委托价格（最优N档委托类型时无需填写价格）
	TpType             string `json:"tp_type"`               // 止盈委托类型,不填默认为market；市价：market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	TpTriggerPriceType string `json:"tp_trigger_price_type"` // 止盈价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	SlTriggerPrice     string `json:"sl_trigger_price"`      // 止损触发价格
	SlOrderPrice       string `json:"sl_order_price"`        // 止损委托价格（最优N档委托类型时无需填写价格）
	SlType             string `json:"sl_type"`               // 止损委托类型,不填默认为market; 市价:market，限价：limit ，最优5档： optimal_5，最优10档：optimal_10，最优20档：optimal_20
	SlTriggerPriceType string `json:"sl_trigger_price_type"` // 止损价格触发类型，默认是最新价 "last":最新价，"mark"：标记价格
	TradeAvgPrice      string `json:"trade_avg_price"`       // 成交均价
	TradeVolume        string `json:"trade_volume"`          // 成交数量
	TradeTurnover      string `json:"trade_turnover"`        // 成交总金额
	FeeCurrency        string `json:"fee_currency"`          // 手续费币种，多个则使用","分隔
	Fee                string `json:"fee"`                   // 总手续费(U)
	PriceProtect       bool   `json:"price_protect"`         // 价差保护，默认为false。仅当设置止盈/止损需要该参数。 false 或者 true
	Profit             string `json:"profit"`                // 平仓盈亏
	ContractType       string `json:"contract_type"`         // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	CancelReason       string `json:"cancel_reason"`         // 撤单原因 "Limit order cancelation by the client":"用户主动撤销普通限价订单","Order cancelation by system":"系统撤单","Market order circuit-breaker":"市价单熔断","Bankruptcy price of market order":"市价单破产价","Order cancelation due to no matching orders":"没有对手盘导致撤单"," Self trading prevention":"防自成交撤单","Number of maker orders matched with your taker orders exceeding limit":"成交笔数超过最大限制","Order cancelation due to API timeout":"API超时撤单"
	CreatedTime        string `json:"created_time"`          // 订单创建时间, UTC时间戳(MS)
	UpdatedTime        string `json:"updated_time"`          // 订单更新时间, UTC时间戳(MS)
	SelfMatchPrevent   string `json:"self_match_prevent"`    // 防自成交 cancel_maker：撤销maker单 cancel_both：撤销全部订单 默认值：cancel_taker
}
type PrivateRestTradeOrderHistoryRes []PrivateRestTradeOrderHistoryResRow

type PrivateRestTradeOrderGetRes PrivateRestTradeOrderHistoryResRow

type PrivateRestTradeOrderLimitRes struct {
	OrderPriceType string `json:"order_price_type"` // 订单报价类型 limit:限价，"opponent":对手价，"lightning":闪电平仓，"optimal_5":最优5档，"optimal_10":最优10档，"optimal_20":最优20档，"fok":FOK订单，"ioc":IOC订单,opponent_ioc"： 对手价-IOC下单，"lightning_ioc"：闪电平仓-IOC下单，"optimal_5_ioc"：最优5档-IOC下单，"optimal_10_ioc"：最优10档-IOC下单，"optimal_20_ioc"：最优20档-IOC下单,"opponent_fok"： 对手价-FOK下单，"lightning_fok"：闪电平仓-FOK下单，"optimal_5_fok"：最优5档-FOK下单，"optimal_10_fok"：最优10档-FOK下单，"optimal_20_fok"：最优20档-FOK下单
	List           []struct {
		Symbol       string  `json:"symbol"`        // 品种代码 BTC,"ETH" ...
		ContractCode string  `json:"contract_code"` // 合约代码
		OpenLimit    float64 `json:"open_limit"`    // 开仓下单量限制
		CloseLimit   float64 `json:"close_limit"`   // 平仓下单量限制
		ContractType string  `json:"contract_type"` // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
		Pair         string  `json:"pair"`          // 交易对 如：“BTC-USDT”
		BusinessType string  `json:"business_type"` // 业务类型 futures：交割、swap：永续、all：全部
	} `json:"list"`
}
