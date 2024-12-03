package mybinanceapi

// Ws账户推送相关
// binance PMC PmCListenKeyPost rest生成listenKey (USER_STREAM)
func (client *PmCRestClient) NewPmCListenKeyPost() *PmCListenKeyPostApi {
	return &PmCListenKeyPostApi{
		client: client,
		req:    &PmCListenKeyPostReq{},
	}
}
func (api *PmCListenKeyPostApi) Do() (*PmCListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCListenKeyPostRes](api.client.c, url, POST)
}

// binance PMC PmCListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *PmCRestClient) NewPmCListenKeyPut() *PmCListenKeyPutApi {
	return &PmCListenKeyPutApi{
		client: client,
		req:    &PmCListenKeyPutReq{},
	}
}
func (api *PmCListenKeyPutApi) Do() (*PmCListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCListenKeyPutRes](api.client.c, url, PUT)
}

// binance PMC PmCListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *PmCRestClient) NewPmCListenKeyDelete() *PmCListenKeyDeleteApi {
	return &PmCListenKeyDeleteApi{
		client: client,
		req:    &PmCListenKeyDeleteReq{},
	}
}
func (api *PmCListenKeyDeleteApi) Do() (*PmCListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCListenKeyDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCListenKeyDeleteRes](api.client.c, url, DELETE)
}
