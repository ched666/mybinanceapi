package mybinanceapi

type PmMApi int

const (

	//杠杆账户接口
	PmMAccount               PmMApi = iota //GET接口 统一账户信息 (USER_DATA)
	PmMAccountBalance                      //GET接口 统一账户余额 (USER_DATA)	PmMMarginAccount                       //GET接口 查询全仓杠杆账户详情 (USER_DATA)
	PmMMarginMaxBorrowable                 //GET接口 查询账户最大可借贷额度(USER_DATA)
	PmMMarginMaxTransferable               //GET接口 查询最大可转出额 (USER_DATA)
	PmMMarginInterestHistory               //GET接口 获取利息历史 (USER_DATA)
	//杠杆订单接口
	PmMMarginOrderGet    //GET接口 查询杠杆账户订单 (USER_DATA)
	PmMMarginOrderPost   //POST接口 杠杆账户下单 (TRADE)
	PmMMarginOrderDelete //DELETE接口 撤销订单 (TRADE)
	PmMMarginAllOrders   //GET接口 查询杠杆账户所有订单 (USER_DATA)
	PmMMarginOpenOrders  //GET接口 查询杠杆账户挂单记录 (USER_DATA)

	PmMMarginOCOOrderGet    //GET接口 查询杠杆账户OCO订单 (USER_DATA)
	PmMMarginOCOOrderPost   //POST接口 杠杆账户OCO订单 (TRADE)
	PmMMarginOCOOrderDelete //DELETE接口 撤销OCO订单 (TRADE)
	PmMMarginOCOAllOrders   //GET接口 查询特定杠杆账户所有 OCO (USER_DATA)
	PmMMarginOCOOpenOrders  //GET接口 查询杠杆账户 OCO 挂单 (USER_DATA)

	// PmMMarginTransfer //POST接口 全仓杠杆账户划转 (MARGIN)
	// PmMMarginLoan     //POST接口 杠杆账户借贷 (MARGIN) 支持逐仓和全仓
	// PmMMarginRepay    //POST接口 杠杆账户归还借贷 (MARGIN) 支持逐仓和全仓

	//通用接口
	PmMPing         //GET接口 测试服务器连通性
	PmMServerTime   //GET接口 获取服务器时间
	PmMExchangeInfo //GET接口 获取交易规则和交易对信息。

	//行情接口
	PmMKlines           //GET接口 K线数据
	PmMTickerPrice      //GET接口 获取交易对最新价格
	PmMDepth            //GET接口 获取深度信息
	PmMTrades           //GET接口 近期成交列表
	PmMHistoricalTrades //GET接口 历史成交记录
	PmMAggTrades        //GET接口 近期成交(归集)
	PmMAvgPrice         //GET接口 当前平均价格
	PmMUiKlines         //GET接口 UIK线数据
	PmMTicker24hr       //GET接口 24hr 价格变动情况
	PmMTickerBookTicker //GET接口 当前最优挂单
	PmMTicker           //GET接口 滚动窗口价格变动统计

	//Ws账户推送相关

	PmMListenKeyPost   //POST接口   生成listenKey(USER_STREAM)
	PmMListenKeyPut    //PUT接口    延长listenKey有效期(USER_STREAM)
	PmMListenKeyDelete //DELETE接口 关闭listenKey(USER_STREAM)

)

var PmMApiMap = map[PmMApi]string{

	//杠杆账户接口
	PmMAccount:               "/papi/v1/account",                      //GET接口 账户信息V2 (USER_DATA)
	PmMAccountBalance:        "/papi/v1/balance",                      //GET接口 统一账户余额 (USER_DATA)
	PmMMarginMaxBorrowable:   "/papi/v1/margin/maxBorrowable",         //GET 查询账户最大可借贷额度(USER_DATA)
	PmMMarginMaxTransferable: "/papi/v1/margin/maxWithdraw",           //GET 查询最大可转出额 (USER_DATA)
	PmMMarginInterestHistory: "/papi/v1/margin/marginInterestHistory", //GET 获取利息历史 (USER_DATA)

	//杠杆订单接口
	PmMMarginOrderGet:    "/papi/v1/margin/order",      //GET 查询杠杆账户订单 (USER_DATA)
	PmMMarginOrderPost:   "/papi/v1/margin/order",      // POST /sapi/v1/margin/order (HMAC SHA256) 杠杆账户下单 (TRADE)
	PmMMarginOrderDelete: "/papi/v1/margin/order",      //DELETE 撤销杠杆账户订单 (TRADE)
	PmMMarginAllOrders:   "/papi/v1/margin/allOrders",  //GET 查询杠杆账户所有订单 (USER_DATA)
	PmMMarginOpenOrders:  "/papi/v1/margin/openOrders", //GET 查询杠杆账户挂单记录 (USER_DATA)

	PmMMarginOCOOrderGet:    "/papi/v1/margin/orderList",     //GET接口 查询杠杆账户OCO订单 (USER_DATA)
	PmMMarginOCOOrderPost:   "/papi/v1/margin/order/oco",     //POST接口 杠杆账户OCO订单 (TRADE)
	PmMMarginOCOOrderDelete: "/papi/v1/margin/orderList",     //DELETE接口 撤销OCO订单 (TRADE)
	PmMMarginOCOAllOrders:   "/papi/v1/margin/allOrderList",  //GET接口 查询特定杠杆账户所有 OCO (USER_DATA)
	PmMMarginOCOOpenOrders:  "/papi/v1/margin/openOrderList", //GET接口 查询杠杆账户 OCO 挂单 (USER_DATA)

	// PmMMarginTransfer: "/sapi/v1/margin/transfer", //POST  全仓杠杆账户划转 (MARGIN)
	// PmMMarginLoan:     "/sapi/v1/margin/loan",     //POST  杠杆账户借贷 (MARGIN) 支持逐仓和全仓
	// PmMMarginRepay:    "/sapi/v1/margin/repay",    //POST  杠杆账户归还借贷 (MARGIN) 支持逐仓和全仓

	//通用接口
	PmMPing:         "/api/v3/ping",         //GET接口 测试连通性
	PmMServerTime:   "/api/v3/time",         //GET接口 获取服务器时间
	PmMExchangeInfo: "/api/v3/exchangeInfo", //GET接口 获取交易规范

	//行情接口
	PmMTickerPrice:      "/api/v3/ticker/price",      //GET接口 获取交易对最新价格
	PmMKlines:           "/api/v3/klines",            //GET接口 获取K线数据
	PmMDepth:            "/api/v3/depth",             //GET接口 获取深度信息
	PmMTrades:           "/api/v3/trades",            //GET接口 近期成交列表
	PmMHistoricalTrades: "/api/v3/historicalTrades",  //GET接口 历史成交记录
	PmMAggTrades:        "/api/v3/aggTrades",         //GET接口 近期成交(归集)
	PmMAvgPrice:         "/api/v3/avgPrice",          //GET接口 当前平均价格
	PmMUiKlines:         "/api/v3/uiKlines",          //GET接口 UIK线数据
	PmMTicker24hr:       "/api/v3/ticker/24hr",       //GET接口 24hr 价格变动情况
	PmMTickerBookTicker: "/api/v3/ticker/bookTicker", //GET接口 当前最优挂单
	PmMTicker:           "/api/v3/ticker",            //GET接口 滚动窗口价格变动统计

	//Ws账户推送相关

	PmMListenKeyPost:   "/papi/v1/listenKey", //POST接口 生成listenKey(USER_STREAM)
	PmMListenKeyPut:    "/papi/v1/listenKey", //PUT接口 延长listenKey有效期(USER_STREAM)
	PmMListenKeyDelete: "/papi/v1/listenKey", //DELETE接口 关闭listenKey(USER_STREAM)
}
