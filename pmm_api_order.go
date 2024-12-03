package mybinanceapi

import "time"

// 交易接口
// binance PMM PmMOpenOrders rest查询当前挂单 (USER_DATA)
func (client *PmMRestClient) NewPmMMarginOrderGet() *PmMMarginOrderGetApi {
	return &PmMMarginOrderGetApi{
		client: client,
		req:    &PmMMarginOrderGetReq{},
	}
}
func (api *PmMMarginOrderGetApi) Do() (*PmMMarginOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOrderGetRes](api.client.c, url, GET)
}

func (client *PmMRestClient) NewPmMMarginOrderPost() *PmMMarginOrderPostApi {
	return &PmMMarginOrderPostApi{
		client: client,
		req:    &PmMMarginOrderPostReq{},
	}
}
func (api *PmMMarginOrderPostApi) Do() (*PmMMarginOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOrderPostRes](api.client.c, url, POST)
}
func (client *PmMRestClient) NewPmMMarginOrderDelete() *PmMMarginOrderDeleteApi {
	return &PmMMarginOrderDeleteApi{
		client: client,
		req:    &PmMMarginOrderDeleteReq{},
	}
}
func (api *PmMMarginOrderDeleteApi) Do() (*PmMMarginOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmMRestClient) NewPmMMarginAllOrders() *PmMMarginAllOrdersApi {
	return &PmMMarginAllOrdersApi{
		client: client,
		req:    &PmMMarginAllOrdersReq{},
	}
}
func (api *PmMMarginAllOrdersApi) Do() (*PmMMarginAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginAllOrdersRes](api.client.c, url, GET)
}
func (client *PmMRestClient) NewPmMMarginOpenOrders() *PmMMarginOpenOrdersApi {
	return &PmMMarginOpenOrdersApi{
		client: client,
		req:    &PmMMarginOpenOrdersReq{},
	}
}
func (api *PmMMarginOpenOrdersApi) Do() (*PmMMarginOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOpenOrdersRes](api.client.c, url, GET)
}

func (client *PmMRestClient) NewPmMMarginOCOOrderGet() *PmMMarginOCOOrderGetApi {
	return &PmMMarginOCOOrderGetApi{
		client: client,
		req:    &PmMMarginOCOOrderGetReq{},
	}
}
func (api *PmMMarginOCOOrderGetApi) Do() (*PmMMarginOCOOrderGetRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOCOOrderGet], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOCOOrderGetRes](api.client.c, url, GET)
}
func (client *PmMRestClient) NewPmMMarginOCOOrderPost() *PmMMarginOCOOrderPostApi {
	return &PmMMarginOCOOrderPostApi{
		client: client,
		req:    &PmMMarginOCOOrderPostReq{},
	}
}
func (api *PmMMarginOCOOrderPostApi) Do() (*PmMMarginOCOOrderPostRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOCOOrderPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOCOOrderPostRes](api.client.c, url, POST)
}
func (client *PmMRestClient) NewPmMMarginOCOOrderDelete() *PmMMarginOCOOrderDeleteApi {
	return &PmMMarginOCOOrderDeleteApi{
		client: client,
		req:    &PmMMarginOCOOrderDeleteReq{},
	}
}
func (api *PmMMarginOCOOrderDeleteApi) Do() (*PmMMarginOCOOrderDeleteRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOCOOrderDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOCOOrderDeleteRes](api.client.c, url, DELETE)
}
func (client *PmMRestClient) NewPmMMarginOCOAllOrders() *PmMMarginOCOAllOrdersApi {
	return &PmMMarginOCOAllOrdersApi{
		client: client,
		req:    &PmMMarginOCOAllOrdersReq{},
	}
}
func (api *PmMMarginOCOAllOrdersApi) Do() (*PmMMarginOCOAllOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOCOAllOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOCOAllOrdersRes](api.client.c, url, GET)
}
func (client *PmMRestClient) NewPmMMarginOCOOpenOrders() *PmMMarginOCOOpenOrdersApi {
	return &PmMMarginOCOOpenOrdersApi{
		client: client,
		req:    &PmMMarginOCOOpenOrdersReq{},
	}
}
func (api *PmMMarginOCOOpenOrdersApi) Do() (*PmMMarginOCOOpenOrdersRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginOCOOpenOrders], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginOCOOpenOrdersRes](api.client.c, url, GET)
}
