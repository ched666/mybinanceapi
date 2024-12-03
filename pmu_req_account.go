package mybinanceapi

// PmUAccount
type PmUAccountReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type PmUAccountApi struct {
	client *PmURestClient
	req    *PmUAccountReq
}

func (api *PmUAccountApi) RecvWindow(RecvWindow int64) *PmUAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAccountApi) Timestamp(Timestamp int64) *PmUAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

// PmUAccountForU
type PmUAccountForUReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type PmUAccountForUApi struct {
	client *PmURestClient
	req    *PmUAccountForUReq
}

func (api *PmUAccountForUApi) RecvWindow(RecvWindow int64) *PmUAccountForUApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAccountForUApi) Timestamp(Timestamp int64) *PmUAccountForUApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUAccountBalanceReq struct {
	Asset      *string `json:"asset"` //NO	资产名称
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmUAccountBalanceApi struct {
	client *PmURestClient
	req    *PmUAccountBalanceReq
}

func (api *PmUAccountBalanceApi) Asset(Asset string) *PmUAccountBalanceApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmUAccountBalanceApi) RecvWindow(RecvWindow int64) *PmUAccountBalanceApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAccountBalanceApi) Timestamp(Timestamp int64) *PmUAccountBalanceApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUCommissionRateReq struct {
	Symbol     *string `json:"symbol"` //Yses	交易对名称
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmUCommissionRateApi struct {
	client *PmURestClient
	req    *PmUCommissionRateReq
}

func (api *PmUCommissionRateApi) Symbol(Symbol string) *PmUCommissionRateApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUCommissionRateApi) RecvWindow(RecvWindow int64) *PmUCommissionRateApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUCommissionRateApi) Timestamp(Timestamp int64) *PmUCommissionRateApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
