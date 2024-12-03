package mybinanceapi

// listenKey相关
type PmCListenKeyPostReq struct{}

type PmCListenKeyPostApi struct {
	client *PmCRestClient
	req    *PmCListenKeyPostReq
}

type PmCListenKeyPutReq struct{}

type PmCListenKeyPutApi struct {
	client *PmCRestClient
	req    *PmCListenKeyPutReq
}

type PmCListenKeyDeleteReq struct{}

type PmCListenKeyDeleteApi struct {
	client *PmCRestClient
	req    *PmCListenKeyDeleteReq
}
