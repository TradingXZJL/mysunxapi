package mysunxapi

type PublicRestAPI int

const (
	// Basic
	PublicRestPublicContractInfo       PublicRestAPI = iota // GET PUBLIC_READ 获取合约信息
	PublicRestPublicIndex                                   // GET PUBLIC_READ 获取合约指数信息
	PublicRestPublicRiskLimit                               // GET PRIVATE_READ 获取合约最高限价和最低限价
	PublicRestPublicFundingRate                             // GET PUBLIC_READ 获取合约的资金费率
	PublicRestPublicFundingRateHistory                      // GET PUBLIC_READ 获取合约的历史资金费率
	PublicRestPublicPriceLimit                              // GET PUBLIC_READ 获取合约最高限价和最低限价
	PublicRestPublicMultiAssetsMargin                       // GET PRIVATE_READ 查询联合保证金支持币种

	// Market
	PublicRestMarketDepth        // GET PUBLIC_READ 获取行情深度数据
	PublicRestMarketHistoryKline // GET PUBLIC_READ 获取K线数据
	PublicRestMarketDetailMerged // GET PUBLIC_READ 获取聚合行情
	PublicRestMarketTrade        // GET PUBLIC_READ 获取最新成交
	PublicRestMarketBBO          // GET PUBLIC_READ 获取市场最优挂单
	PublicRestMarketHistoryTrade // GET PUBLIC_READ 批量获取最近的交易记录
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	// Public
	PublicRestPublicContractInfo:       "/sapi/v1/public/contract_info",        // GET PUBLIC_READ 获取合约信息
	PublicRestPublicIndex:              "/sapi/v1/public/index",                // GET PUBLIC_READ 获取合约指数信息
	PublicRestPublicRiskLimit:          "/sapi/v1/public/risk/limit",           // GET PRIVATE_READ 获取合约最高限价和最低限价
	PublicRestPublicFundingRate:        "/sapi/v1/public/funding_rate",         // GET PUBLIC_READ 获取合约的资金费率
	PublicRestPublicFundingRateHistory: "/sapi/v1/public/funding_rate_history", // GET PUBLIC_READ 获取合约的历史资金费率
	PublicRestPublicPriceLimit:         "/sapi/v1/public/price_limit",          // GET PUBLIC_READ 获取合约最高限价和最低限价
	PublicRestPublicMultiAssetsMargin:  "/sapi/v1/public/multi_assets_margin",  // GET PRIVATE_READ 查询联合保证金支持币种

	// Market
	PublicRestMarketDepth:        "/sapi/v1/market/depth",         // GET PUBLIC_READ 获取行情深度数据
	PublicRestMarketHistoryKline: "/sapi/v1/market/history/kline", // GET PUBLIC_READ 获取K线数据
	PublicRestMarketDetailMerged: "/sapi/v1/market/detail/merged", // GET PUBLIC_READ 获取聚合行情
	PublicRestMarketTrade:        "/sapi/v1/market/trade",         // GET PUBLIC_READ 获取最新成交
	PublicRestMarketBBO:          "/sapi/v1/market/bbo",           // GET PUBLIC_READ 获取市场最优挂单
	PublicRestMarketHistoryTrade: "/sapi/v1/market/history/trade", // GET PUBLIC_READ 批量获取最近的交易记录
}
