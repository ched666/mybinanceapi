package mybinanceapi

// listenKey相关
type PmUListenKeyPostReq struct{}

type PmUListenKeyPostApi struct {
	client *PmURestClient
	req    *PmUListenKeyPostReq
}

type PmUListenKeyPutReq struct{}

type PmUListenKeyPutApi struct {
	client *PmURestClient
	req    *PmUListenKeyPutReq
}

type PmUListenKeyDeleteReq struct{}

type PmUListenKeyDeleteApi struct {
	client *PmURestClient
	req    *PmUListenKeyDeleteReq
}
