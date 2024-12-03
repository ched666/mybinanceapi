package mybinanceapi

// listenKey相关
type PmMListenKeyPostReq struct{}

type PmMListenKeyPostApi struct {
	client *PmMRestClient
	req    *PmMListenKeyPostReq
}

type PmMListenKeyPutReq struct{}

type PmMListenKeyPutApi struct {
	client *PmMRestClient
	req    *PmMListenKeyPutReq
}

type PmMListenKeyDeleteReq struct{}

type PmMListenKeyDeleteApi struct {
	client *PmMRestClient
	req    *PmMListenKeyDeleteReq
}
