package mybinanceapi

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type PmUOpenOrdersReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmUOpenOrdersApi struct {
	client *PmURestClient
	req    *PmUOpenOrdersReq
}

func (api *PmUOpenOrdersApi) Symbol(Symbol string) *PmUOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUOpenOrdersApi) RecvWindow(RecvWindow int64) *PmUOpenOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUOpenOrdersApi) Timestamp(Timestamp int64) *PmUOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUAllOrdersReq struct {
	Symbol     *string `json:"symbol"`    //No	交易对
	OrderId    *int64  `json:"orderId"`   //No	只返回此orderID及之后的订单，缺省返回最近的订单
	StartTime  *int64  `json:"startTime"` //No	起始时间
	EndTime    *int64  `json:"endTime"`   //No	结束时间
	Limit      *int64  `json:"limit"`     //No	返回的结果集数量 默认值:500 最大值:1000
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmUAllOrdersApi struct {
	client *PmURestClient
	req    *PmUAllOrdersReq
}

func (api *PmUAllOrdersApi) Symbol(Symbol string) *PmUAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUAllOrdersApi) OrderId(OrderId int64) *PmUAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmUAllOrdersApi) StartTime(StartTime int64) *PmUAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *PmUAllOrdersApi) EndTime(EndTime int64) *PmUAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *PmUAllOrdersApi) Limit(Limit int64) *PmUAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *PmUAllOrdersApi) RecvWindow(RecvWindow int64) *PmUAllOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAllOrdersApi) Timestamp(Timestamp int64) *PmUAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                            //Yes	交易对
	Side                    *string          `json:"side"`                              //Yes	买卖方向 SELL, BUY
	PositionSide            *string          `json:"positionSide,omitempty"`            //No	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type                    *string          `json:"type"`                              //Yes	订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	TimeInForce             *string          `json:"timeInForce,omitempty"`             //No	有效方法
	Quantity                *decimal.Decimal `json:"quantity,omitempty"`                //No	下单数量,使用closePosition不支持此参数。
	ReduceOnly              *string          `json:"reduceOnly,omitempty"`              //No	true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	Price                   *decimal.Decimal `json:"price,omitempty"`                   //No	委托价格
	NewClientOrderId        *string          `json:"newClientOrderId,omitempty"`        //No	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	NewOrderRespType        *string          `json:"newOrderRespType,omitempty"`        //No	"ACK", "RESULT", 默认 "ACK"
	PriceMatch              *string          `json:"priceMatch,omitempty"`              //No	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode,omitempty"` //No	"GTC", "IOC", "FOK", "GTX"
	GoodTillDate            *int64           `json:"goodTillDate,omitempty"`            //No	订单有效期时间戳，单位毫秒
	RecvWindow              *int64           `json:"recvWindow,omitempty"`              //No
	Timestamp               *int64           `json:"timestamp,omitempty"`               //Yes
}
type PmUOrderPostApi struct {
	client *PmURestClient
	req    *PmUOrderPostReq
}

func (api *PmUOrderPostApi) Symbol(Symbol string) *PmUOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUOrderPostApi) Side(Side string) *PmUOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmUOrderPostApi) PositionSide(PositionSide string) *PmUOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *PmUOrderPostApi) Type(Type string) *PmUOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *PmUOrderPostApi) TimeInForce(TimeInForce string) *PmUOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *PmUOrderPostApi) Quantity(Quantity decimal.Decimal) *PmUOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmUOrderPostApi) ReduceOnly(ReduceOnly string) *PmUOrderPostApi {
	api.req.ReduceOnly = GetPointer(ReduceOnly)
	return api
}
func (api *PmUOrderPostApi) Price(Price decimal.Decimal) *PmUOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmUOrderPostApi) NewClientOrderId(NewClientOrderId string) *PmUOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *PmUOrderPostApi) NewOrderRespType(NewOrderRespType string) *PmUOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *PmUOrderPostApi) PriceMatch(PriceMatch string) *PmUOrderPostApi {
	api.req.PriceMatch = GetPointer(PriceMatch)
	return api
}
func (api *PmUOrderPostApi) SelfTradePreventionMode(SelfTradePreventionMode string) *PmUOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *PmUOrderPostApi) GoodTillDate(GoodTillDate int64) *PmUOrderPostApi {
	api.req.GoodTillDate = GetPointer(GoodTillDate)
	return api
}
func (api *PmUOrderPostApi) RecvWindow(RecvWindow int64) *PmUOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUOrderPostApi) Timestamp(Timestamp int64) *PmUOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUOrderPutReq struct {
	OrderId           *string          `json:"orderId,omitempty"`           //NO	系统订单号
	OrigClientOrderId *string          `json:"origClientOrderId,omitempty"` //NO	用户自定义的订单号
	Symbol            *string          `json:"symbol"`                      //YES	交易对
	Side              *string          `json:"side"`                        //YES	买卖方向 SELL, BUY; side需要和原订单相同
	Quantity          *decimal.Decimal `json:"quantity"`                    //YES	下单数量,使用closePosition不支持此参数。
	Price             *decimal.Decimal `json:"price"`                       //YES	委托价格
	PriceMatch        *string          `json:"priceMatch,omitempty"`        //NO	OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	RecvWindow        *int64           `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64           `json:"timestamp,omitempty"`         //YES
}
type PmUOrderPutApi struct {
	client *PmURestClient
	req    *PmUOrderPutReq
}

func (api *PmUOrderPutApi) Symbol(Symbol string) *PmUOrderPutApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUOrderPutApi) OrderId(OrderId int64) *PmUOrderPutApi {
	orderIdStr := strconv.FormatInt(OrderId, BIT_BASE_10)
	api.req.OrderId = GetPointer(orderIdStr)
	return api
}
func (api *PmUOrderPutApi) OrigClientOrderId(OrigClientOrderId string) *PmUOrderPutApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmUOrderPutApi) Side(Side string) *PmUOrderPutApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmUOrderPutApi) Quantity(Quantity decimal.Decimal) *PmUOrderPutApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmUOrderPutApi) Price(Price decimal.Decimal) *PmUOrderPutApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmUOrderPutApi) PriceMatch(PriceMatch string) *PmUOrderPutApi {
	api.req.PriceMatch = GetPointer(PriceMatch)
	return api
}
func (api *PmUOrderPutApi) RecvWindow(RecvWindow int64) *PmUOrderPutApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUOrderPutApi) Timestamp(Timestamp int64) *PmUOrderPutApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUOrderGetReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
}
type PmUOrderGetApi struct {
	client *PmURestClient
	req    *PmUOrderGetReq
}

func (api *PmUOrderGetApi) Symbol(Symbol string) *PmUOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUOrderGetApi) OrderId(OrderId int64) *PmUOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmUOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *PmUOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmUOrderGetApi) RecvWindow(RecvWindow int64) *PmUOrderGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUOrderGetApi) Timestamp(Timestamp int64) *PmUOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
}
type PmUOrderDeleteApi struct {
	client *PmURestClient
	req    *PmUOrderDeleteReq
}

func (api *PmUOrderDeleteApi) Symbol(Symbol string) *PmUOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUOrderDeleteApi) OrderId(OrderId int64) *PmUOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmUOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *PmUOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmUOrderDeleteApi) RecvWindow(RecvWindow int64) *PmUOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUOrderDeleteApi) Timestamp(Timestamp int64) *PmUOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUAllOrderDeleteReq struct {
	Symbol     *string `json:"symbol"`               //YES 交易对
	RecvWindow *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp  *int64  `json:"timestamp"`            //YES
}
type PmUAllOrderDeleteApi struct {
	client *PmURestClient
	req    *PmUAllOrderDeleteReq
}

func (api *PmUAllOrderDeleteApi) Symbol(Symbol string) *PmUAllOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUAllOrderDeleteApi) RecvWindow(RecvWindow int64) *PmUAllOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAllOrderDeleteApi) Timestamp(Timestamp int64) *PmUAllOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUConditionalOrderPostReq struct {
	Symbol                  *string          `json:"symbol"`                  //YES	交易对
	Side                    *string          `json:"side"`                    //YES	买卖方向
	PositionSide            *string          `json:"positionSide"`            //NO	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	StrategyType            *string          `json:"strategyType"`            //YES	条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	TimeInForce             *string          `json:"timeInForce"`             //NO	订单有效期
	Quantity                *decimal.Decimal `json:"quantity"`                //NO	下单数量
	ReduceOnly              *string          `json:"reduceOnly"`              //NO	true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price                   *decimal.Decimal `json:"price"`                   //NO	下单价格
	WorkingType             *string          `json:"workingType"`             //NO	stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect            *string          `json:"priceProtect"`            //NO	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewClientStrategyId     *string          `json:"newClientStrategyId"`     //NO	不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StopPrice               *decimal.Decimal `json:"stopPrice"`               //NO	Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice         *decimal.Decimal `json:"activationPrice"`         //NO	TRAILING_STOP_MARKET 单使用，默认标记价格
	CallbackRate            *decimal.Decimal `json:"callbackRate"`            //NO	TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
	PriceMatch              *string          `json:"priceMatch"`              //NO	OPPONENT/ OPPONENT_5/ OPPONENT_10/ OPPONENT_20/QUEUE/ QUEUE_5/ QUEUE_10/ QUEUE_20；不能与price同时传
	SelfTradePreventionMode *string          `json:"selfTradePreventionMode"` //NO	NONE / EXPIRE_TAKER/ EXPIRE_MAKER/ EXPIRE_BOTH； 默认NONE
	GoodTillDate            *int64           `json:"goodTillDate"`            //NO	TIF为GTD时订单的自动取消时间， 当timeInforce为GTD时必传；传入的时间戳仅保留秒级精度，毫秒级部分会被自动忽略，时间戳需大于当前时间+600s且小于253402300799000
	RecvWindow              *int64           `json:"recvWindow"`              //NO
	Timestamp               *int64           `json:"timestamp"`               //YES
}

type PmUConditionalOrderPostApi struct {
	client *PmURestClient
	req    *PmUConditionalOrderPostReq
}

func (api *PmUConditionalOrderPostApi) Symbol(Symbol string) *PmUConditionalOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

func (api *PmUConditionalOrderPostApi) Side(Side string) *PmUConditionalOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmUConditionalOrderPostApi) PositionSide(PositionSide string) *PmUConditionalOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *PmUConditionalOrderPostApi) StrategyType(StrategyType string) *PmUConditionalOrderPostApi {
	api.req.StrategyType = GetPointer(StrategyType)
	return api
}
func (api *PmUConditionalOrderPostApi) TimeInForce(TimeInForce string) *PmUConditionalOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *PmUConditionalOrderPostApi) Quantity(Quantity decimal.Decimal) *PmUConditionalOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmUConditionalOrderPostApi) Price(Price decimal.Decimal) *PmUConditionalOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmUConditionalOrderPostApi) WorkingType(WorkingType string) *PmUConditionalOrderPostApi {
	api.req.WorkingType = GetPointer(WorkingType)
	return api
}
func (api *PmUConditionalOrderPostApi) PriceProtect(PriceProtect string) *PmUConditionalOrderPostApi {
	api.req.PriceProtect = GetPointer(PriceProtect)
	return api
}
func (api *PmUConditionalOrderPostApi) NewClientStrategyId(NewClientStrategyId string) *PmUConditionalOrderPostApi {
	api.req.NewClientStrategyId = GetPointer(NewClientStrategyId)
	return api
}
func (api *PmUConditionalOrderPostApi) StopPrice(StopPrice decimal.Decimal) *PmUConditionalOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *PmUConditionalOrderPostApi) ActivationPrice(ActivationPrice decimal.Decimal) *PmUConditionalOrderPostApi {
	api.req.ActivationPrice = GetPointer(ActivationPrice)
	return api
}
func (api *PmUConditionalOrderPostApi) CallbackRate(CallbackRate decimal.Decimal) *PmUConditionalOrderPostApi {
	api.req.CallbackRate = GetPointer(CallbackRate)
	return api
}
func (api *PmUConditionalOrderPostApi) PriceMatch(PriceMatch string) *PmUConditionalOrderPostApi {
	api.req.PriceMatch = GetPointer(PriceMatch)
	return api
}
func (api *PmUConditionalOrderPostApi) SelfTradePreventionMode(SelfTradePreventionMode string) *PmUConditionalOrderPostApi {
	api.req.SelfTradePreventionMode = GetPointer(SelfTradePreventionMode)
	return api
}
func (api *PmUConditionalOrderPostApi) GoodTillDate(GoodTillDate int64) *PmUConditionalOrderPostApi {
	api.req.GoodTillDate = GetPointer(GoodTillDate)
	return api
}
func (api *PmUConditionalOrderPostApi) RecvWindow(RecvWindow int64) *PmUConditionalOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUConditionalOrderPostApi) Timestamp(Timestamp int64) *PmUConditionalOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUConditionalOrderDeleteReq struct {
	Symbol              *string `json:"symbol"`               //YES	交易对
	StrategyId          *string `json:"strategyId"`           //NO strategyId 与 newClientStrategyId 之一必须发送
	NewClientStrategyId *string `json:"newClientStrategyId"`  //NO
	RecvWindow          *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp           *int64  `json:"timestamp"`            //YES
}
type PmUConditionalOrderDeleteApi struct {
	client *PmURestClient
	req    *PmUConditionalOrderDeleteReq
}

func (api *PmUConditionalOrderDeleteApi) Symbol(Symbol string) *PmUConditionalOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUConditionalOrderDeleteApi) StrategyId(StrategyId string) *PmUConditionalOrderDeleteApi {
	api.req.StrategyId = GetPointer(StrategyId)
	return api
}
func (api *PmUConditionalOrderDeleteApi) NewClientStrategyId(NewClientStrategyId string) *PmUConditionalOrderDeleteApi {
	api.req.NewClientStrategyId = GetPointer(NewClientStrategyId)
	return api
}
func (api *PmUConditionalOrderDeleteApi) RecvWindow(RecvWindow int64) *PmUConditionalOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUConditionalOrderDeleteApi) Timestamp(Timestamp int64) *PmUConditionalOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUAllConditionnalOrderDeleteReq struct {
	Symbol     *string `json:"symbol"`               //YES	交易对
	RecvWindow *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp  *int64  `json:"timestamp"`            //YES
}
type PmUAllConditionnalOrderDeleteApi struct {
	client *PmURestClient
	req    *PmUAllConditionnalOrderDeleteReq
}

func (api *PmUAllConditionnalOrderDeleteApi) Symbol(Symbol string) *PmUAllConditionnalOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUAllConditionnalOrderDeleteApi) RecvWindow(RecvWindow int64) *PmUAllConditionnalOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUAllConditionnalOrderDeleteApi) Timestamp(Timestamp int64) *PmUAllConditionnalOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmUUserTradesReq struct {
	Symbol     *string `json:"symbol"`     //YES	交易对
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	FromId     *int64  `json:"fromId"`     //NO 返回该fromId及之后的成交，缺省返回最近的成交
	Limit      *int64  `json:"limit"`      //NO 返回的结果集数量 默认值:50 最大值:1000
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type PmUUserTradesApi struct {
	client *PmURestClient
	req    *PmUUserTradesReq
}

func (api *PmUUserTradesApi) Symbol(Symbol string) *PmUUserTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmUUserTradesApi) StartTime(StartTime int64) *PmUUserTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *PmUUserTradesApi) EndTime(EndTime int64) *PmUUserTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *PmUUserTradesApi) FromId(FromId int64) *PmUUserTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *PmUUserTradesApi) Limit(Limit int64) *PmUUserTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *PmUUserTradesApi) RecvWindow(RecvWindow int64) *PmUUserTradesApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmUUserTradesApi) Timestamp(Timestamp int64) *PmUUserTradesApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
