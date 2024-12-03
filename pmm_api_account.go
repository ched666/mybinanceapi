package mybinanceapi

import "time"

// 账户接口
// binance PMM  PmMAccount 统一账户信息 (USER_DATA)
func (client *PmMRestClient) NewPmMAccount() *PmMAccountApi {
	return &PmMAccountApi{
		client: client,
		req:    &PmMAccountReq{},
	}
}
func (api *PmMAccountApi) Do() (*PmMAccountRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMAccount], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMAccountRes](api.client.c, url, GET)
}

// binance PMM  PmMAccountBalance 统一账户余额 (USER_DATA)
func (client *PmMRestClient) NewPmMAccountBalance() *PmMAccountBalanceApi {
	return &PmMAccountBalanceApi{
		client: client,
		req:    &PmMAccountBalanceReq{},
	}
}
func (api *PmMAccountBalanceApi) Do() (*PmMAccountBalanceRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMAccountBalance], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMAccountBalanceRes](api.client.c, url, GET)
}

func (client *PmMRestClient) NewPmMMarginMaxBorrowable() *PmMMarginMaxBorrowableApi {
	return &PmMMarginMaxBorrowableApi{
		client: client,
		req:    &PmMMarginMaxBorrowableReq{},
	}
}
func (api *PmMMarginMaxBorrowableApi) Do() (*PmMMarginMaxBorrowableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginMaxBorrowable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginMaxBorrowableRes](api.client.c, url, GET)
}
func (client *PmMRestClient) NewPmMMarginMaxTransferable() *PmMMarginMaxTransferableApi {
	return &PmMMarginMaxTransferableApi{
		client: client,
		req:    &PmMMarginMaxTransferableReq{},
	}
}
func (api *PmMMarginMaxTransferableApi) Do() (*PmMMarginMaxTransferableRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginMaxTransferable], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginMaxTransferableRes](api.client.c, url, GET)
}
func (client *PmMRestClient) NewPmMMarginInterestHistory() *PmMMarginInterestHistoryApi {
	return &PmMMarginInterestHistoryApi{
		client: client,
		req:    &PmMMarginInterestHistoryReq{},
	}
}
func (api *PmMMarginInterestHistoryApi) Do() (*PmMMarginInterestHistoryRes, error) {
	api.Timestamp(time.Now().UnixMilli() + serverTimeDelta)
	url := binanceHandlerRequestApiWithSecret(PMM, api.req, PmMApiMap[PmMMarginInterestHistory], api.client.c.ApiSecret)
	return binanceCallApiWithSecret[PmMMarginInterestHistoryRes](api.client.c, url, GET)
}
