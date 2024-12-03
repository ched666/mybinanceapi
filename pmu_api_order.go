package mybinanceapi

import "time"

// 交易接口
// binance PMU PmUOpenOrders rest查询当前挂单 (USER_DATA)
func (client *PmURestClient) NewOpenOrders() *PmUOpenOrdersApi {
	return &PmUOpenOrdersApi{
		client: client,
		req:    &PmUOpenOrdersReq{},
	}
}
func (api *PmUOpenOrdersApi) Do() (*PmUOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUOpenOrdersRes](api.client.c, url, GET)
}

// binance PMU PmUAllOrders rest查询所有订单 (USER_DATA)
func (client *PmURestClient) NewAllOrders() *PmUAllOrdersApi {
	return &PmUAllOrdersApi{
		client: client,
		req:    &PmUAllOrdersReq{},
	}
}
func (api *PmUAllOrdersApi) Do() (*PmUAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAllOrdersRes](api.client.c, url, GET)
}

// binance PMU PmUOrderPost rest下单 (TRADE)
func (client *PmURestClient) NewPmUOrderPost() *PmUOrderPostApi {
	return &PmUOrderPostApi{
		client: client,
		req:    &PmUOrderPostReq{},
	}
}
func (api *PmUOrderPostApi) Do() (*PmUOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUOrderPostRes](api.client.c, url, POST)
}

// binance PMU PmUOrderPut rest修改订单 (TRADE)
func (client *PmURestClient) NewPmUOrderPut() *PmUOrderPutApi {
	return &PmUOrderPutApi{
		client: client,
		req:    &PmUOrderPutReq{},
	}
}
func (api *PmUOrderPutApi) Do() (*PmUOrderPutRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUOrderPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUOrderPutRes](api.client.c, url, PUT)
}

// binance PMU PmUOrderGet  rest查询订单 (USER_DATA)
func (client *PmURestClient) NewPmUOrderGet() *PmUOrderGetApi {
	return &PmUOrderGetApi{
		client: client,
		req:    &PmUOrderGetReq{},
	}
}
func (api *PmUOrderGetApi) Do() (*PmUOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUOrderGetRes](api.client.c, url, GET)
}

// binance PMU PmUOrderDelete rest撤销订单 (TRADE)
func (client *PmURestClient) NewPmUOrderDelete() *PmUOrderDeleteApi {
	return &PmUOrderDeleteApi{
		client: client,
		req:    &PmUOrderDeleteReq{},
	}
}
func (api *PmUOrderDeleteApi) Do() (*PmUOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUOrderDeleteRes](api.client.c, url, DELETE)
}

// binance PMU PmUAllOrderDelete rest撤销全部订单 (TRADE)
func (client *PmURestClient) NewPmUAllOrderDelete() *PmUAllOrderDeleteApi {
	return &PmUAllOrderDeleteApi{
		client: client,
		req:    &PmUAllOrderDeleteReq{},
	}
}
func (api *PmUAllOrderDeleteApi) Do() (*PmUAllOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAllOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAllOrderDeleteRes](api.client.c, url, DELETE)
}

func (client *PmURestClient) NewPmUConditionalOrderPost() *PmUConditionalOrderPostApi {
	return &PmUConditionalOrderPostApi{
		client: client,
		req:    &PmUConditionalOrderPostReq{},
	}
}
func (api *PmUConditionalOrderPostApi) Do() (*PmUConditionalOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUConditionalOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUConditionalOrderPostRes](api.client.c, url, POST)
}
func (client *PmURestClient) NewPmUConditionalOrderDelete() *PmUConditionalOrderDeleteApi {
	return &PmUConditionalOrderDeleteApi{
		client: client,
		req:    &PmUConditionalOrderDeleteReq{},
	}
}
func (api *PmUConditionalOrderDeleteApi) Do() (*PmUConditionalOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUConditionalOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUConditionalOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmURestClient) NewPmUAllConditionnalOrderDelete() *PmUAllConditionnalOrderDeleteApi {
	return &PmUAllConditionnalOrderDeleteApi{
		client: client,
		req:    &PmUAllConditionnalOrderDeleteReq{},
	}
}
func (api *PmUAllConditionnalOrderDeleteApi) Do() (*PmUAllConditionnalOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAllConditionnalOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAllConditionnalOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmURestClient) NewPmUUserTrades() *PmUUserTradesApi {
	return &PmUUserTradesApi{
		client: client,
		req:    &PmUUserTradesReq{},
	}
}
func (api *PmUUserTradesApi) Do() (*PmUUserTradesRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUUserTrades], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUUserTradesRes](api.client.c, url, GET)
}
