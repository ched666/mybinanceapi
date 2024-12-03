package mybinanceapi

import "sync"

type SpotWsStreamClient struct {
	WsStreamClient
	spotWsType     SpotWsType
	isolatedSymbol string
	client         *SpotRestClient
}
type FutureWsStreamClient struct {
	WsStreamClient
	client *FutureRestClient
}
type SwapWsStreamClient struct {
	WsStreamClient
	client *SwapRestClient
}

//zsk修改
type PmUWsStreamClient struct {
	WsStreamClient
	client *PmURestClient
}
type PmCWsStreamClient struct {
	WsStreamClient
	client *PmCRestClient
}
type PmMWsStreamClient struct {
	WsStreamClient
	client *PmMRestClient
}

func (*MyBinance) NewSpotWsStreamClient() *SpotWsStreamClient {
	ws := &SpotWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SPOT,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewFutureWsStreamClient() *FutureWsStreamClient {
	ws := &FutureWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       FUTURE,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewSwapWsStreamClient() *SwapWsStreamClient {
	ws := &SwapWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       SWAP,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}

//zsk修改
func (*MyBinance) NewPmUWsStreamClient() *PmUWsStreamClient {
	ws := &PmUWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       PMU,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewPmCWsStreamClient() *PmCWsStreamClient {
	ws := &PmCWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       PMC,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (*MyBinance) NewPmMWsStreamClient() *PmMWsStreamClient {
	ws := &PmMWsStreamClient{
		WsStreamClient: WsStreamClient{
			apiType:       PMM,
			isGzip:        false,
			reSubscribeMu: &sync.Mutex{},
			wsStreamPath:  WS_STREAM_PATH,
		},
	}
	ws.initStructs()
	return ws
}
func (ws *SpotWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *FutureWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *SwapWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}

//zsk修改
func (ws *PmUWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *PmCWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
func (ws *PmMWsStreamClient) Close() error {
	if ws.isListenWs {
		err := ws.listenKeyDelete()
		if err != nil {
			return err
		}
		*ws.listenKeyRefreshStopChan <- struct{}{}
	}
	return ws.WsStreamClient.Close()
}
