package mybinanceapi

import "time"

// 账户接口
// binance PMC  PmCAccount 统一账户信息 (USER_DATA)
func (client *PmCRestClient) NewPmCAccount() *PmCAccountApi {
	return &PmCAccountApi{
		client: client,
		req:    &PmCAccountReq{},
	}
}
func (api *PmCAccountApi) Do() (*PmCAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAccountRes](api.client.c, url, GET)
}

// binance PMC  PmCAccountBalance 统一账户余额 (USER_DATA)
func (client *PmCRestClient) NewPmCAccountBalance() *PmCAccountBalanceApi {
	return &PmCAccountBalanceApi{
		client: client,
		req:    &PmCAccountBalanceReq{},
	}
}
func (api *PmCAccountBalanceApi) Do() (*PmCAccountBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAccountBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAccountBalanceRes](api.client.c, url, GET)
}

// binance PMC PmCAccountForC UM账户信息v2 (USER_DATA)
func (client *PmCRestClient) NewPmCAccountForC() *PmCAccountForCApi {
	return &PmCAccountForCApi{
		client: client,
		req:    &PmCAccountForCReq{},
	}
}
func (api *PmCAccountForCApi) Do() (*PmCAccountForCRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCAccountForC], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCAccountForCRes](api.client.c, url, GET)
}

func (client *PmCRestClient) NewPmCCommissionRate() *PmCCommissionRateApi {
	return &PmCCommissionRateApi{
		client: client,
		req:    &PmCCommissionRateReq{},
	}
}
func (api *PmCCommissionRateApi) Do() (*PmCCommissionRateRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMC, api.req, PmCApiMap[PmCCommissionRate], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmCCommissionRateRes](api.client.c, url, GET)
}
