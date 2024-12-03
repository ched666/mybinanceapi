package mybinanceapi

import "time"

// 交易接口
// binance PMC PmCOpenOrders rest查询当前挂单 (USER_DATA)
func (client *PmCRestClient) NewOpenOrders() *PmCOpenOrdersApi {
	return &PmCOpenOrdersApi{
		client: client,
		req:    &PmCOpenOrdersReq{},
	}
}
func (api *PmCOpenOrdersApi) Do() (*PmCOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCOpenOrdersRes](api.client.c, url, GET)
}

// binance PMC PmCAllOrders rest查询所有订单 (USER_DATA)
func (client *PmCRestClient) NewAllOrders() *PmCAllOrdersApi {
	return &PmCAllOrdersApi{
		client: client,
		req:    &PmCAllOrdersReq{},
	}
}
func (api *PmCAllOrdersApi) Do() (*PmCAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAllOrdersRes](api.client.c, url, GET)
}

// binance PMC PmCOrderPost rest下单 (TRADE)
func (client *PmCRestClient) NewPmCOrderPost() *PmCOrderPostApi {
	return &PmCOrderPostApi{
		client: client,
		req:    &PmCOrderPostReq{},
	}
}
func (api *PmCOrderPostApi) Do() (*PmCOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCOrderPostRes](api.client.c, url, POST)
}

// binance PMC PmCOrderPut rest修改订单 (TRADE)
func (client *PmCRestClient) NewPmCOrderPut() *PmCOrderPutApi {
	return &PmCOrderPutApi{
		client: client,
		req:    &PmCOrderPutReq{},
	}
}
func (api *PmCOrderPutApi) Do() (*PmCOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCOrderPutRes](api.client.c, url, PUT)
}

// binance PMC PmCOrderGet  rest查询订单 (USER_DATA)
func (client *PmCRestClient) NewPmCOrderGet() *PmCOrderGetApi {
	return &PmCOrderGetApi{
		client: client,
		req:    &PmCOrderGetReq{},
	}
}
func (api *PmCOrderGetApi) Do() (*PmCOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCOrderGetRes](api.client.c, url, GET)
}

// binance PMC PmCOrderDelete rest撤销订单 (TRADE)
func (client *PmCRestClient) NewPmCOrderDelete() *PmCOrderDeleteApi {
	return &PmCOrderDeleteApi{
		client: client,
		req:    &PmCOrderDeleteReq{},
	}
}
func (api *PmCOrderDeleteApi) Do() (*PmCOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCOrderDeleteRes](api.client.c, url, DELETE)
}

// binance PMC PmCAllOrderDelete rest撤销全部订单 (TRADE)
func (client *PmCRestClient) NewPmCAllOrderDelete() *PmCAllOrderDeleteApi {
	return &PmCAllOrderDeleteApi{
		client: client,
		req:    &PmCAllOrderDeleteReq{},
	}
}
func (api *PmCAllOrderDeleteApi) Do() (*PmCAllOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAllOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAllOrderDeleteRes](api.client.c, url, DELETE)
}

func (client *PmCRestClient) NewPmCConditionalOrderPost() *PmCConditionalOrderPostApi {
	return &PmCConditionalOrderPostApi{
		client: client,
		req:    &PmCConditionalOrderPostReq{},
	}
}
func (api *PmCConditionalOrderPostApi) Do() (*PmCConditionalOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCConditionalOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCConditionalOrderPostRes](api.client.c, url, POST)
}
func (client *PmCRestClient) NewPmCConditionalOrderDelete() *PmCConditionalOrderDeleteApi {
	return &PmCConditionalOrderDeleteApi{
		client: client,
		req:    &PmCConditionalOrderDeleteReq{},
	}
}
func (api *PmCConditionalOrderDeleteApi) Do() (*PmCConditionalOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCConditionalOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCConditionalOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmCRestClient) NewPmCAllConditionnalOrderDelete() *PmCAllConditionnalOrderDeleteApi {
	return &PmCAllConditionnalOrderDeleteApi{
		client: client,
		req:    &PmCAllConditionnalOrderDeleteReq{},
	}
}
func (api *PmCAllConditionnalOrderDeleteApi) Do() (*PmCAllConditionnalOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAllConditionnalOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAllConditionnalOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmCRestClient) NewPmCUserTrades() *PmCUserTradesApi {
	return &PmCUserTradesApi{
		client: client,
		req:    &PmCUserTradesReq{},
	}
}
func (api *PmCUserTradesApi) Do() (*PmCUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCUserTrades], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCUserTradesRes](api.client.c, url, GET)
}
