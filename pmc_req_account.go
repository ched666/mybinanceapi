package mybinanceapi

// PmCAccount
type PmCAccountReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type PmCAccountApi struct {
	client *PmCRestClient
	req    *PmCAccountReq
}

func (api *PmCAccountApi) RecvWindow(RecvWindow int64) *PmCAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAccountApi) Timestamp(Timestamp int64) *PmCAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// PmCAccountForU
type PmCAccountForCReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type PmCAccountForCApi struct {
	client *PmCRestClient
	req    *PmCAccountForCReq
}

func (api *PmCAccountForCApi) RecvWindow(RecvWindow int64) *PmCAccountForCApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAccountForCApi) Timestamp(Timestamp int64) *PmCAccountForCApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCAccountBalanceReq struct {
	Asset      *string `json:"asset"` //NO	资产名称
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmCAccountBalanceApi struct {
	client *PmCRestClient
	req    *PmCAccountBalanceReq
}

func (api *PmCAccountBalanceApi) Asset(Asset string) *PmCAccountBalanceApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmCAccountBalanceApi) RecvWindow(RecvWindow int64) *PmCAccountBalanceApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAccountBalanceApi) Timestamp(Timestamp int64) *PmCAccountBalanceApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCCommissionRateReq struct {
	Symbol     *string `json:"symbol"` //Yses	交易对名称
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmCCommissionRateApi struct {
	client *PmCRestClient
	req    *PmCCommissionRateReq
}

func (api *PmCCommissionRateApi) Symbol(Symbol string) *PmCCommissionRateApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCCommissionRateApi) RecvWindow(RecvWindow int64) *PmCCommissionRateApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCCommissionRateApi) Timestamp(Timestamp int64) *PmCCommissionRateApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
