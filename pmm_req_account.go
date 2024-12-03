package mybinanceapi

// PmMAccount
type PmMAccountReq struct {
	RecvWindow *int64 `json:"recvWindow"`
	Timestamp  *int64 `json:"timestamp"`
}
type PmMAccountApi struct {
	client *PmMRestClient
	req    *PmMAccountReq
}

func (api *PmMAccountApi) RecvWindow(RecvWindow int64) *PmMAccountApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmMAccountApi) Timestamp(Timestamp int64) *PmMAccountApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmMAccountBalanceReq struct {
	Asset      *string `json:"asset"` //NO	资产名称
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmMAccountBalanceApi struct {
	client *PmMRestClient
	req    *PmMAccountBalanceReq
}

func (api *PmMAccountBalanceApi) Asset(Asset string) *PmMAccountBalanceApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmMAccountBalanceApi) RecvWindow(RecvWindow int64) *PmMAccountBalanceApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmMAccountBalanceApi) Timestamp(Timestamp int64) *PmMAccountBalanceApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmMMarginMaxBorrowableReq struct {
	Asset      *string `json:"asset"` //YES
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmMMarginMaxBorrowableApi struct {
	client *PmMRestClient
	req    *PmMMarginMaxBorrowableReq
}

func (api *PmMMarginMaxBorrowableApi) Asset(Asset string) *PmMMarginMaxBorrowableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmMMarginMaxBorrowableApi) RecvWindow(RecvWindow int64) *PmMMarginMaxBorrowableApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmMMarginMaxBorrowableApi) Timestamp(Timestamp int64) *PmMMarginMaxBorrowableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmMMarginMaxTransferableReq struct {
	Asset      *string `json:"asset"` //YES
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmMMarginMaxTransferableApi struct {
	client *PmMRestClient
	req    *PmMMarginMaxTransferableReq
}

func (api *PmMMarginMaxTransferableApi) Asset(Asset string) *PmMMarginMaxTransferableApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmMMarginMaxTransferableApi) RecvWindow(RecvWindow int64) *PmMMarginMaxTransferableApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmMMarginMaxTransferableApi) Timestamp(Timestamp int64) *PmMMarginMaxTransferableApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmMMarginInterestHistoryReq struct {
	Asset      *string `json:"asset"` //NO
	StartTime  *int64  `json:"startTime"`
	EndTime    *int64  `json:"endTime"`
	Current    *int64  `json:"current"`  //当前查询页。 开始值 1. 默认:1
	Size       *int64  `json:"size"`     //默认:10 最大:100
	Archived   *string `json:"archived"` //默认: false. 查询6个月以前的数据，需要设为 true
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmMMarginInterestHistoryApi struct {
	client *PmMRestClient
	req    *PmMMarginInterestHistoryReq
}

func (api *PmMMarginInterestHistoryApi) Asset(Asset string) *PmMMarginInterestHistoryApi {
	api.req.Asset = GetPointer(Asset)
	return api
}
func (api *PmMMarginInterestHistoryApi) StartTime(StartTime int64) *PmMMarginInterestHistoryApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *PmMMarginInterestHistoryApi) EndTime(EndTime int64) *PmMMarginInterestHistoryApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *PmMMarginInterestHistoryApi) Current(Current int64) *PmMMarginInterestHistoryApi {
	api.req.Current = GetPointer(Current)
	return api
}
func (api *PmMMarginInterestHistoryApi) Size(Size int64) *PmMMarginInterestHistoryApi {
	api.req.Size = GetPointer(Size)
	return api
}
func (api *PmMMarginInterestHistoryApi) Archived(Archived string) *PmMMarginInterestHistoryApi {
	api.req.Archived = GetPointer(Archived)
	return api
}
func (api *PmMMarginInterestHistoryApi) RecvWindow(RecvWindow int64) *PmMMarginInterestHistoryApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmMMarginInterestHistoryApi) Timestamp(Timestamp int64) *PmMMarginInterestHistoryApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
