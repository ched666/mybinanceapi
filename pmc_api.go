package mybinanceapi

type PmCApi int

const (
	//账户接口
	PmCAccount        PmCApi = iota //GET接口 统一账户信息 (USER_DATA)
	PmCAccountBalance               //GET接口 统一账户余额 (USER_DATA)
	PmCAccountForC                  //GET接口 CM账户信息v2 (USER_DATA)
	PmCCommissionRate               //GET接口 (HMAC SHA256)查询用户CM手续费率 (USER_DATA)

	//交易接口
	PmCOpenOrders                 //GET接口 (HMAC SHA256)查看当前全部CM挂单(USER_DATA)
	PmCAllOrders                  //GET接口 (HMAC SHA256)查询所有CM订单(包括历史订单)(USER_DATA)
	PmCOrderPost                  //POST接口 (HMAC SHA256)CM下单(TRADE)
	PmCOrderPut                   //PUT接口 (HMAC SHA256)修改CM订单(TRADE)
	PmCOrderGet                   //GET接口 (HMAC SHA256)查询CM订单(USER_DATA)
	PmCOrderDelete                //DELETE接口 (HMAC SHA256)撤销CM订单 (TRADE)
	PmCAllOrderDelete             //DELETE接口 (HMAC SHA256)撤销全部CM订单 (TRADE)
	PmCConditionalOrderPost       //POST接口 (HMAC SHA256)CM条件单下单(TRADE)
	PmCConditionalOrderDelete     //DELETE接口 (HMAC SHA256)取消CM条件订单(TRADE)
	PmCAllConditionnalOrderDelete //DELETE接口 (HMAC SHA256)撤销全部CM条件订单(TRADE)
	PmCUserTrades                 //GET接口 (HMAC SHA256)CM账户成交历史 (USER_DATA)

	//通用接口
	PmCPing         //GET接口 测试服务器连通性
	PmCServerTime   //GET接口 获取服务器时间
	PmCExchangeInfo //GET接口 交易规则和交易对信息

	//行情接口
	PmCKlines           //K线数据
	PmCDepth            //深度信息
	PmCTrades           //最新成交
	PmCHistoricalTrades //历史成交
	PmCAggTrades        //近期成交(归集)
	PmCPremiumIndex     //最新标记价格和资金费率
	PmCFundingRate      //查询资金费率历史
	PmCFundingInfo      //查询资金费率信息
	PmCTicker24hr       //24hr价格变动情况
	PmCTickerPrice      //最新价格
	PmCTickerBookTicker //当前最优挂单
	PmCDataBasis        //基差数据

	//Ws账户推送相关接口
	PmCListenKeyPost   //生成listenKey (USER_STREAM)
	PmCListenKeyPut    //延长listenKey有效期 (USER_STREAM)
	PmCListenKeyDelete //关闭listenKey (USER_STREAM)
)

var PmCApiMap = map[PmCApi]string{

	//账户接口
	PmCAccount:        "/papi/v1/account",           //GET接口 账户信息V2 (USER_DATA)
	PmCAccountBalance: "/papi/v1/balance",           //GET接口 统一账户余额 (USER_DATA)
	PmCAccountForC:    "/papi/v2/cm/account",        //GET接口 CM账户信息v2 (USER_DATA)
	PmCCommissionRate: "/papi/v1/cm/commissionRate", //GET接口 (HMAC SHA256)查询用户CM手续费率 (USER_DATA)

	//交易接口
	PmCOpenOrders:                 "/papi/v1/cm/openOrders",                //GET接口 (HMAC SHA256)查看当前全部CM挂单(USER_DATA)
	PmCAllOrders:                  "/papi/v1/cm/allOrders",                 //GET接口 (HMAC SHA256)查询所有CM订单(包括历史订单)(USER_DATA)
	PmCOrderPost:                  "/papi/v1/cm/order",                     //POST接口 (HMAC SHA256)CM下单(TRADE)
	PmCOrderPut:                   "/papi/v1/cm/order",                     //PUT接口 (HMAC SHA256)修改CM订单(TRADE)
	PmCOrderGet:                   "/papi/v1/cm/order",                     //GET接口 (HMAC SHA256)查询CM订单(USER_DATA)
	PmCOrderDelete:                "/papi/v1/cm/order",                     //DELETE接口 (HMAC SHA256)撤销CM订单 (TRADE)
	PmCAllOrderDelete:             "/papi/v1/cm/allOpenOrders",             //DELETE接口 (HMAC SHA256)撤销全部CM订单 (TRADE)
	PmCConditionalOrderPost:       "/papi/v1/cm/conditional/order",         //POST接口 (HMAC SHA256)CM条件单下单(TRADE)
	PmCConditionalOrderDelete:     "/papi/v1/cm/conditional/order",         //DELETE接口 (HMAC SHA256)取消CM条件订单(TRADE)
	PmCAllConditionnalOrderDelete: "/papi/v1/cm/conditional/allOpenOrders", //DELETE接口 (HMAC SHA256)撤销全部CM条件订单(TRADE)
	PmCUserTrades:                 "/papi/v1/cm/userTrades",                //GET接口 (HMAC SHA256)CM账户成交历史 (USER_DATA)

	//通用接口
	PmCPing:         "/fapi/v1/ping",         //GET接口 测试服务器连通性
	PmCServerTime:   "/fapi/v1/time",         //GET接口 获取服务器时间
	PmCExchangeInfo: "/fapi/v1/exchangeInfo", //GET接口 交易规则和交易对信息

	//行情接口
	PmCKlines:           "/fapi/v1/klines",            //GET接口 K线数据
	PmCDepth:            "/fapi/v1/depth",             //GET接口 深度信息
	PmCTrades:           "/fapi/v1/trades",            //GET接口 最新成交
	PmCHistoricalTrades: "/fapi/v1/historicalTrades",  //GET接口 查询历史成交(MARKET_DATA)
	PmCAggTrades:        "/fapi/v1/aggTrades",         //GET接口 近期成交(归集)
	PmCPremiumIndex:     "/fapi/v1/premiumIndex",      //GET接口 最新标记价格和资金费率
	PmCFundingRate:      "/fapi/v1/fundingRate",       //GET接口 查询资金费率历史
	PmCFundingInfo:      "/fapi/v1/fundingInfo",       //GET接口 查询资金费率信息
	PmCTicker24hr:       "/fapi/v1/ticker/24hr",       //GET接口 24hr价格变动情况
	PmCTickerPrice:      "/fapi/v1/ticker/price",      //GET接口 最新价格
	PmCTickerBookTicker: "/fapi/v1/ticker/bookTicker", //GET接口 当前最优挂单
	PmCDataBasis:        "/futures/data/basis",        //GET接口 基差数据

	//Ws账户推送相关接口
	PmCListenKeyPost:   "/papi/v1/listenKey", //POST接口 生成listenKey (USER_STREAM)
	PmCListenKeyPut:    "/papi/v1/listenKey", //PUT接口 延长listenKey有效期 (USER_STREAM)
	PmCListenKeyDelete: "/papi/v1/listenKey", //DELETE接口 关闭listenKey (USER_STREAM)

}
