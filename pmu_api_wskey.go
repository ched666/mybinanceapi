package mybinanceapi

// Ws账户推送相关
// binance PMU PmUListenKeyPost rest生成listenKey (USER_STREAM)
func (client *PmURestClient) NewPmUListenKeyPost() *PmUListenKeyPostApi {
	return &PmUListenKeyPostApi{
		client: client,
		req:    &PmUListenKeyPostReq{},
	}
}
func (api *PmUListenKeyPostApi) Do() (*PmUListenKeyPostRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUListenKeyPost], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUListenKeyPostRes](api.client.c, url, POST)
}

// binance PMU PmUListenKeyPut rest延长listenKey有效期 (USER_STREAM)
func (client *PmURestClient) NewPmUListenKeyPut() *PmUListenKeyPutApi {
	return &PmUListenKeyPutApi{
		client: client,
		req:    &PmUListenKeyPutReq{},
	}
}
func (api *PmUListenKeyPutApi) Do() (*PmUListenKeyPutRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUListenKeyPut], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUListenKeyPutRes](api.client.c, url, PUT)
}

// binance PMU PmUListenKeyDelete rest关闭listenKey (USER_STREAM)
func (client *PmURestClient) NewPmUListenKeyDelete() *PmUListenKeyDeleteApi {
	return &PmUListenKeyDeleteApi{
		client: client,
		req:    &PmUListenKeyDeleteReq{},
	}
}
func (api *PmUListenKeyDeleteApi) Do() (*PmUListenKeyDeleteRes, error) {
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUListenKeyDelete], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUListenKeyDeleteRes](api.client.c, url, DELETE)
}
