package mybinanceapi

type PmUApi int

const (
	//账户接口
	PmUAccount        PmUApi = iota //GET接口 统一账户信息 (USER_DATA)
	PmUAccountBalance               //GET接口 统一账户余额 (USER_DATA)
	PmUAccountForU                  //GET接口 UM账户信息v2 (USER_DATA)
	PmUCommissionRate               //GET接口 (HMAC SHA256)查询用户UM手续费率 (USER_DATA)

	//交易接口
	PmUOpenOrders                 //GET接口 (HMAC SHA256)查看当前全部UM挂单(USER_DATA)
	PmUAllOrders                  //GET接口 (HMAC SHA256)查询所有UM订单(包括历史订单)(USER_DATA)
	PmUOrderPost                  //POST接口 (HMAC SHA256)UM下单(TRADE)
	PmUOrderPut                   //PUT接口 (HMAC SHA256)修改UM订单(TRADE)
	PmUOrderGet                   //GET接口 (HMAC SHA256)查询UM订单(USER_DATA)
	PmUOrderDelete                //DELETE接口 (HMAC SHA256)撤销UM订单 (TRADE)
	PmUAllOrderDelete             //DELETE接口 (HMAC SHA256)撤销全部UM订单 (TRADE)
	PmUConditionalOrderPost       //POST接口 (HMAC SHA256)UM条件单下单(TRADE)
	PmUConditionalOrderDelete     //DELETE接口 (HMAC SHA256)取消UM条件订单(TRADE)
	PmUAllConditionnalOrderDelete //DELETE接口 (HMAC SHA256)撤销全部UM条件订单(TRADE)
	PmUUserTrades                 //GET接口 (HMAC SHA256)UM账户成交历史 (USER_DATA)

	//通用接口
	PmUPing         //GET接口 测试服务器连通性
	PmUServerTime   //GET接口 获取服务器时间
	PmUExchangeInfo //GET接口 交易规则和交易对信息

	//行情接口
	PmUKlines           //K线数据
	PmUDepth            //深度信息
	PmUTrades           //最新成交
	PmUHistoricalTrades //历史成交
	PmUAggTrades        //近期成交(归集)
	PmUPremiumIndex     //最新标记价格和资金费率
	PmUFundingRate      //查询资金费率历史
	PmUFundingInfo      //查询资金费率信息
	PmUTicker24hr       //24hr价格变动情况
	PmUTickerPrice      //最新价格
	PmUTickerBookTicker //当前最优挂单
	PmUDataBasis        //基差数据

	//Ws账户推送相关接口
	PmUListenKeyPost   //生成listenKey (USER_STREAM)
	PmUListenKeyPut    //延长listenKey有效期 (USER_STREAM)
	PmUListenKeyDelete //关闭listenKey (USER_STREAM)
)

var PmUApiMap = map[PmUApi]string{

	//账户接口
	PmUAccount:        "/papi/v1/account",           //GET接口 账户信息V2 (USER_DATA)
	PmUAccountBalance: "/papi/v1/balance",           //GET接口 统一账户余额 (USER_DATA)
	PmUAccountForU:    "/papi/v2/um/account",        //GET接口 UM账户信息v2 (USER_DATA)
	PmUCommissionRate: "/papi/v1/um/commissionRate", //GET接口 (HMAC SHA256)查询用户UM手续费率 (USER_DATA)

	//交易接口
	PmUOpenOrders:                 "/papi/v1/um/openOrders",                //GET接口 (HMAC SHA256)查看当前全部UM挂单(USER_DATA)
	PmUAllOrders:                  "/papi/v1/um/allOrders",                 //GET接口 (HMAC SHA256)查询所有UM订单(包括历史订单)(USER_DATA)
	PmUOrderPost:                  "/papi/v1/um/order",                     //POST接口 (HMAC SHA256)UM下单(TRADE)
	PmUOrderPut:                   "/papi/v1/um/order",                     //PUT接口 (HMAC SHA256)修改UM订单(TRADE)
	PmUOrderGet:                   "/papi/v1/um/order",                     //GET接口 (HMAC SHA256)查询UM订单(USER_DATA)
	PmUOrderDelete:                "/papi/v1/um/order",                     //DELETE接口 (HMAC SHA256)撤销UM订单 (TRADE)
	PmUAllOrderDelete:             "/papi/v1/um/allOpenOrders",             //DELETE接口 (HMAC SHA256)撤销全部UM订单 (TRADE)
	PmUConditionalOrderPost:       "/papi/v1/um/conditional/order",         //POST接口 (HMAC SHA256)UM条件单下单(TRADE)
	PmUConditionalOrderDelete:     "/papi/v1/um/conditional/order",         //DELETE接口 (HMAC SHA256)取消UM条件订单(TRADE)
	PmUAllConditionnalOrderDelete: "/papi/v1/um/conditional/allOpenOrders", //DELETE接口 (HMAC SHA256)撤销全部UM条件订单(TRADE)
	PmUUserTrades:                 "/papi/v1/um/userTrades",                //GET接口 (HMAC SHA256)UM账户成交历史 (USER_DATA)

	//通用接口
	PmUPing:         "/fapi/v1/ping",         //GET接口 测试服务器连通性
	PmUServerTime:   "/fapi/v1/time",         //GET接口 获取服务器时间
	PmUExchangeInfo: "/fapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	PmUKlines:           "/fapi/v1/klines",            //GET接口 K线数据
	PmUDepth:            "/fapi/v1/depth",             //GET接口 深度信息
	PmUTrades:           "/fapi/v1/trades",            //GET接口 最新成交
	PmUHistoricalTrades: "/fapi/v1/historicalTrades",  //GET接口 查询历史成交(MARKET_DATA)
	PmUAggTrades:        "/fapi/v1/aggTrades",         //GET接口 近期成交(归集)
	PmUPremiumIndex:     "/fapi/v1/premiumIndex",      //GET接口 最新标记价格和资金费率
	PmUFundingRate:      "/fapi/v1/fundingRate",       //GET接口 查询资金费率历史
	PmUFundingInfo:      "/fapi/v1/fundingInfo",       //GET接口 查询资金费率信息
	PmUTicker24hr:       "/fapi/v1/ticker/24hr",       //GET接口 24hr价格变动情况
	PmUTickerPrice:      "/fapi/v1/ticker/price",      //GET接口 最新价格
	PmUTickerBookTicker: "/fapi/v1/ticker/bookTicker", //GET接口 当前最优挂单
	PmUDataBasis:        "/futures/data/basis",        //GET接口 基差数据

	//Ws账户推送相关接口
	PmUListenKeyPost:   "/papi/v1/listenKey", //POST接口 生成listenKey (USER_STREAM)
	PmUListenKeyPut:    "/papi/v1/listenKey", //PUT接口 延长listenKey有效期 (USER_STREAM)
	PmUListenKeyDelete: "/papi/v1/listenKey", //DELETE接口 关闭listenKey (USER_STREAM)

}
