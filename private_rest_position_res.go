package mysunxapi

type PrivateRestTradePositionOpensResRow struct {
	ContractCode      string `json:"contract_code"`      // 合约代码
	PositionSide      string `json:"position_side"`      // 仓位方向 持仓方向 买卖模式下默认值both， 开平仓模式下： long：平多
	Direction         string `json:"direction"`          // 订单方向 "buy":买， "sell":卖
	MarginMode        string `json:"margin_mode"`        // 保证金模式  cross：全仓
	OpenAvgPrice      string `json:"open_avg_price"`     // 开仓均价
	Volume            string `json:"volume"`             // 持仓量（张）
	Available         string `json:"available"`          // 可平仓数量（张）
	LeverRate         int64  `json:"lever_rate"`         // 杠杠倍数
	AdlRiskPercent    int    `json:"adl_risk_percent"`   // 自动减仓指示  (1、2、3、4、5)，1档风险最低，5档风险最高
	LiquidationPrice  string `json:"liquidation_price"`  // 预估强平价
	InitialMargin     string `json:"initial_margin"`     // 初始保证金，仅适用于全仓
	MaintenanceMargin string `json:"maintenance_margin"` // 维持保证金
	ProfitUnreal      string `json:"profit_unreal"`      // 未实现盈亏
	ProfitRate        string `json:"profit_rate"`        // 未实现收益率
	MarginRate        string `json:"margin_rate"`        // 保证金率-对标调整信息
	MarginCurrency    string `json:"margin_currency"`    // 保证金币种（计价币种）
	LastPrice         string `json:"last_price"`         // 最新价格
	MarkPrice         string `json:"mark_price"`         // 标记价格
	ContractType      string `json:"contract_type"`      // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	CreatedTime       string `json:"created_time"`       // 创建时间
	UpdatedTime       string `json:"updated_time"`       // 更新时间
}
type PrivateRestTradePositionOpensRes []PrivateRestTradePositionOpensResRow

type PrivateRestPositionLeverResRow struct {
	ContractCode   string `json:"contract_code"`   // 合约代码
	ContractType   string `json:"contract_type"`   // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	MarginMode     string `json:"margin_mode"`     // 保证金模式 cross：全仓
	LeverRate      int64  `json:"lever_rate"`      // 杠杆等级
	AvailableLever []int  `json:"available_lever"` // 可用杠杆等级
}
type PrivateRestPositionLeverGetRes []PrivateRestPositionLeverResRow

type PrivateRestPositionLeverPostRes struct {
	ContractCode string `json:"contract_code"` // 合约代码
	MarginMode   string `json:"margin_mode"`   // 保证金模式
	LeverRate    int    `json:"lever_rate"`    // 杠杆等级
}

type PrivateRestPositionModeGetRes struct {
	PositionMode string `json:"position_mode"` // 持仓模式 single_side：单向持仓；dual_side：双向持仓
}

type PrivateRestPositionModePostRes struct {
	PositionMode string `json:"position_mode"` // 持仓模式 single_side：单向持仓；dual_side：双向持仓
}

type PrivateRestPositionRiskLimitResRow struct {
	ContractCode          string `json:"contract_code"`           // 合约代码
	MarginMode            string `json:"margin_mode"`             // 保证金模式
	PositionSide          string `json:"position_side"`           // 仓位方向
	MaxLever              string `json:"max_lever"`               // 该层最大杠杆等级
	MaintenanceMarginRate string `json:"maintenance_margin_rate"` // 该层维持保证率
	MaxVolume             string `json:"max_volume"`              // 该层最高持有张数
	MinVolume             string `json:"min_volume"`              // 该层最低持有张数
}
type PrivateRestPositionRiskLimitRes []PrivateRestPositionRiskLimitResRow

type PrivateRestPositionPositionLimitResRow struct {
	Symbol         string  `json:"symbol"`           // 品种代码 BTC,"ETH" ...
	ContractCode   string  `json:"contract_code"`    // 合约代码 永续："BTC-USDT"... ，交割：”BTC-USDT-210625“
	MarginMode     string  `json:"margin_mode"`      // 保证金模式 cross：全仓模式；
	BuyLimit       int     `json:"buy_limit"`        // 合约多仓持仓的最大值，单位为张
	SellLimit      int     `json:"sell_limit"`       // 合约空仓持仓的最大值，单位为张
	ContractType   string  `json:"contract_type"`    // 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	Pair           string  `json:"pair"`             // 合约对 如：“BTC-USDT”
	BusinessType   string  `json:"business_type"`    // 业务类型 futures：交割、swap：永续
	LeverRate      int     `json:"lever_rate"`       // 用户当前品种杠杆倍数
	BuyLimitValue  float64 `json:"buy_limit_value"`  // 合约多仓持仓价值上限，单位USDT
	SellLimitValue float64 `json:"sell_limit_value"` // 合约空仓持仓价值上限，单位USDT
	MarkPrice      float64 `json:"mark_price"`       // 当前品种标记价格（以该价格用于计算持仓张数）
	// TradePartition string  `json:"trade_partition"`  // 参数已废弃
}
type PrivateRestPositionPositionLimitRes []PrivateRestPositionPositionLimitResRow
