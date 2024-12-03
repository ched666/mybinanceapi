package mybinanceapi

import "time"

// 账户接口
// binance PMU  PmUAccount 统一账户信息 (USER_DATA)
func (client *PmURestClient) NewPmUAccount() *PmUAccountApi {
	return &PmUAccountApi{
		client: client,
		req:    &PmUAccountReq{},
	}
}
func (api *PmUAccountApi) Do() (*PmUAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAccountRes](api.client.c, url, GET)
}

// binance PMU  PmUAccountBalance 统一账户余额 (USER_DATA)
func (client *PmURestClient) NewPmUAccountBalance() *PmUAccountBalanceApi {
	return &PmUAccountBalanceApi{
		client: client,
		req:    &PmUAccountBalanceReq{},
	}
}
func (api *PmUAccountBalanceApi) Do() (*PmUAccountBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAccountBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAccountBalanceRes](api.client.c, url, GET)
}

// binance PMU PmUAccountForU UM账户信息v2 (USER_DATA)
func (client *PmURestClient) NewPmUAccountForU() *PmUAccountForUApi {
	return &PmUAccountForUApi{
		client: client,
		req:    &PmUAccountForUReq{},
	}
}
func (api *PmUAccountForUApi) Do() (*PmUAccountForURes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUAccountForU], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUAccountForURes](api.client.c, url, GET)
}

func (client *PmURestClient) NewPmUCommissionRate() *PmUCommissionRateApi {
	return &PmUCommissionRateApi{
		client: client,
		req:    &PmUCommissionRateReq{},
	}
}
func (api *PmUCommissionRateApi) Do() (*PmUCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMU, api.req, PmUApiMap[PmUCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmUCommissionRateRes](api.client.c, url, GET)
}
