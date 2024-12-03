package mybinanceapi

// Ws账户推送相关
// binance PMM PmMListenKeyPost rest生成listenKey (USER_STREAM)
func (client *PmMRestClient) NewPmMListenKeyPost() *PmMListenKeyPostApi {
	return &PmMListenKeyPostApi{
		client: client,
		req:    &PmMListenKeyPostReq{},
	}
}
func (api *PmMListenKeyPostApi) Do() (*PmMListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMListenKeyPostRes](api.client.c, url, POST)
}

// binance PMM PmMListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *PmMRestClient) NewPmMListenKeyPut() *PmMListenKeyPutApi {
	return &PmMListenKeyPutApi{
		client: client,
		req:    &PmMListenKeyPutReq{},
	}
}
func (api *PmMListenKeyPutApi) Do() (*PmMListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMListenKeyPutRes](api.client.c, url, PUT)
}

// binance PMM PmMListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *PmMRestClient) NewPmMListenKeyDelete() *PmMListenKeyDeleteApi {
	return &PmMListenKeyDeleteApi{
		client: client,
		req:    &PmMListenKeyDeleteReq{},
	}
}
func (api *PmMListenKeyDeleteApi) Do() (*PmMListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMListenKeyDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMListenKeyDeleteRes](api.client.c, url, DELETE)
}
