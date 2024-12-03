package mybinanceapi

import (
	"strconv"

	"github.com/shopspring/decimal"
)

type PmCOpenOrdersReq struct {
	Symbol     *string `json:"symbol"` //No	交易对
	Pair       *string `json:"pair"`   //No
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmCOpenOrdersApi struct {
	client *PmCRestClient
	req    *PmCOpenOrdersReq
}

func (api *PmCOpenOrdersApi) Symbol(Symbol string) *PmCOpenOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCOpenOrdersApi) Pair(Pair string) *PmCOpenOrdersApi {
	api.req.Pair = GetPointer(Pair)
	return api
}
func (api *PmCOpenOrdersApi) RecvWindow(RecvWindow int64) *PmCOpenOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCOpenOrdersApi) Timestamp(Timestamp int64) *PmCOpenOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCAllOrdersReq struct {
	Symbol     *string `json:"symbol"`    //No	交易对
	Pair       *string `json:"pair"`      //No
	OrderId    *int64  `json:"orderId"`   //No	只返回此orderID及之后的订单，缺省返回最近的订单
	StartTime  *int64  `json:"startTime"` //No	起始时间
	EndTime    *int64  `json:"endTime"`   //No	结束时间
	Limit      *int64  `json:"limit"`     //No	返回的结果集数量 默认值:500 最大值:1000
	RecvWindow *int64  `json:"recvWindow"`
	Timestamp  *int64  `json:"timestamp"`
}
type PmCAllOrdersApi struct {
	client *PmCRestClient
	req    *PmCAllOrdersReq
}

func (api *PmCAllOrdersApi) Symbol(Symbol string) *PmCAllOrdersApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCAllOrdersApi) Pair(Pair string) *PmCAllOrdersApi {
	api.req.Pair = GetPointer(Pair)
	return api
}
func (api *PmCAllOrdersApi) OrderId(OrderId int64) *PmCAllOrdersApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmCAllOrdersApi) StartTime(StartTime int64) *PmCAllOrdersApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *PmCAllOrdersApi) EndTime(EndTime int64) *PmCAllOrdersApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *PmCAllOrdersApi) Limit(Limit int64) *PmCAllOrdersApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *PmCAllOrdersApi) RecvWindow(RecvWindow int64) *PmCAllOrdersApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAllOrdersApi) Timestamp(Timestamp int64) *PmCAllOrdersApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCOrderPostReq struct {
	Symbol           *string          `json:"symbol"`                     //Yes	交易对
	Side             *string          `json:"side"`                       //Yes	买卖方向 SELL, BUY
	PositionSide     *string          `json:"positionSide,omitempty"`     //No	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	Type             *string          `json:"type"`                       //Yes	订单类型 LIMIT, MARKET, STOP, TAKE_PROFIT, STOP_MARKET, TAKE_PROFIT_MARKET, TRAILING_STOP_MARKET
	TimeInForce      *string          `json:"timeInForce,omitempty"`      //No	有效方法
	Quantity         *decimal.Decimal `json:"quantity,omitempty"`         //No	下单数量,使用closePosition不支持此参数。
	ReduceOnly       *string          `json:"reduceOnly,omitempty"`       //No	true, false; 非双开模式下默认false；双开模式下不接受此参数； 使用closePosition不支持此参数。
	Price            *decimal.Decimal `json:"price,omitempty"`            //No	委托价格
	NewClientOrderId *string          `json:"newClientOrderId,omitempty"` //No	用户自定义的订单号，不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则 ^[\.A-Z\:/a-z0-9_-]{1,36}$
	NewOrderRespType *string          `json:"newOrderRespType,omitempty"` //No	"ACK", "RESULT", 默认 "ACK"
	RecvWindow       *int64           `json:"recvWindow,omitempty"`       //No
	Timestamp        *int64           `json:"timestamp,omitempty"`        //Yes
}
type PmCOrderPostApi struct {
	client *PmCRestClient
	req    *PmCOrderPostReq
}

func (api *PmCOrderPostApi) Symbol(Symbol string) *PmCOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCOrderPostApi) Side(Side string) *PmCOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmCOrderPostApi) PositionSide(PositionSide string) *PmCOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *PmCOrderPostApi) Type(Type string) *PmCOrderPostApi {
	api.req.Type = GetPointer(Type)
	return api
}
func (api *PmCOrderPostApi) TimeInForce(TimeInForce string) *PmCOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *PmCOrderPostApi) Quantity(Quantity decimal.Decimal) *PmCOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmCOrderPostApi) ReduceOnly(ReduceOnly string) *PmCOrderPostApi {
	api.req.ReduceOnly = GetPointer(ReduceOnly)
	return api
}
func (api *PmCOrderPostApi) Price(Price decimal.Decimal) *PmCOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmCOrderPostApi) NewClientOrderId(NewClientOrderId string) *PmCOrderPostApi {
	api.req.NewClientOrderId = GetPointer(NewClientOrderId)
	return api
}
func (api *PmCOrderPostApi) NewOrderRespType(NewOrderRespType string) *PmCOrderPostApi {
	api.req.NewOrderRespType = GetPointer(NewOrderRespType)
	return api
}
func (api *PmCOrderPostApi) RecvWindow(RecvWindow int64) *PmCOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCOrderPostApi) Timestamp(Timestamp int64) *PmCOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCOrderPutReq struct {
	OrderId           *string          `json:"orderId,omitempty"`           //NO	系统订单号
	OrigClientOrderId *string          `json:"origClientOrderId,omitempty"` //NO	用户自定义的订单号
	Symbol            *string          `json:"symbol"`                      //YES	交易对
	Side              *string          `json:"side"`                        //YES	买卖方向 SELL, BUY; side需要和原订单相同
	Quantity          *decimal.Decimal `json:"quantity"`                    //YES	下单数量,使用closePosition不支持此参数。
	Price             *decimal.Decimal `json:"price"`                       //YES	委托价格
	RecvWindow        *int64           `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64           `json:"timestamp,omitempty"`         //YES
}
type PmCOrderPutApi struct {
	client *PmCRestClient
	req    *PmCOrderPutReq
}

func (api *PmCOrderPutApi) Symbol(Symbol string) *PmCOrderPutApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCOrderPutApi) OrderId(OrderId int64) *PmCOrderPutApi {
	orderIdStr := strconv.FormatInt(OrderId, BIT_BASE_10)
	api.req.OrderId = GetPointer(orderIdStr)
	return api
}
func (api *PmCOrderPutApi) OrigClientOrderId(OrigClientOrderId string) *PmCOrderPutApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmCOrderPutApi) Side(Side string) *PmCOrderPutApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmCOrderPutApi) Quantity(Quantity decimal.Decimal) *PmCOrderPutApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmCOrderPutApi) Price(Price decimal.Decimal) *PmCOrderPutApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmCOrderPutApi) RecvWindow(RecvWindow int64) *PmCOrderPutApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCOrderPutApi) Timestamp(Timestamp int64) *PmCOrderPutApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCOrderGetReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
}
type PmCOrderGetApi struct {
	client *PmCRestClient
	req    *PmCOrderGetReq
}

func (api *PmCOrderGetApi) Symbol(Symbol string) *PmCOrderGetApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCOrderGetApi) OrderId(OrderId int64) *PmCOrderGetApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmCOrderGetApi) OrigClientOrderId(OrigClientOrderId string) *PmCOrderGetApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmCOrderGetApi) RecvWindow(RecvWindow int64) *PmCOrderGetApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCOrderGetApi) Timestamp(Timestamp int64) *PmCOrderGetApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCOrderDeleteReq struct {
	Symbol            *string `json:"symbol"`                      //YES 交易对
	OrderId           *int64  `json:"orderId,omitempty"`           //NO 系统订单号
	OrigClientOrderId *string `json:"origClientOrderId,omitempty"` //NO 用户自定义的订单号
	RecvWindow        *int64  `json:"recvWindow,omitempty"`        //NO
	Timestamp         *int64  `json:"timestamp"`                   //YES
}
type PmCOrderDeleteApi struct {
	client *PmCRestClient
	req    *PmCOrderDeleteReq
}

func (api *PmCOrderDeleteApi) Symbol(Symbol string) *PmCOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCOrderDeleteApi) OrderId(OrderId int64) *PmCOrderDeleteApi {
	api.req.OrderId = GetPointer(OrderId)
	return api
}
func (api *PmCOrderDeleteApi) OrigClientOrderId(OrigClientOrderId string) *PmCOrderDeleteApi {
	api.req.OrigClientOrderId = GetPointer(OrigClientOrderId)
	return api
}
func (api *PmCOrderDeleteApi) RecvWindow(RecvWindow int64) *PmCOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCOrderDeleteApi) Timestamp(Timestamp int64) *PmCOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCAllOrderDeleteReq struct {
	Symbol     *string `json:"symbol"`               //YES 交易对
	RecvWindow *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp  *int64  `json:"timestamp"`            //YES
}
type PmCAllOrderDeleteApi struct {
	client *PmCRestClient
	req    *PmCAllOrderDeleteReq
}

func (api *PmCAllOrderDeleteApi) Symbol(Symbol string) *PmCAllOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCAllOrderDeleteApi) RecvWindow(RecvWindow int64) *PmCAllOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAllOrderDeleteApi) Timestamp(Timestamp int64) *PmCAllOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCConditionalOrderPostReq struct {
	Symbol              *string          `json:"symbol"`              //YES	交易对
	Side                *string          `json:"side"`                //YES	买卖方向
	PositionSide        *string          `json:"positionSide"`        //NO	持仓方向，单向持仓模式下非必填，默认且仅可填BOTH;在双向持仓模式下必填,且仅可选择 LONG 或 SHORT
	StrategyType        *string          `json:"strategyType"`        //YES	条件单类型"STOP", "STOP_MARKET", "TAKE_PROFIT", "TAKE_PROFIT_MARKET"或"TRAILING_STOP_MARKET"
	TimeInForce         *string          `json:"timeInForce"`         //NO	订单有效期
	Quantity            *decimal.Decimal `json:"quantity"`            //NO	下单数量
	ReduceOnly          *string          `json:"reduceOnly"`          //NO	true或false; 非双开模式下默认false；双开模式下不接受此参数
	Price               *decimal.Decimal `json:"price"`               //NO	下单价格
	WorkingType         *string          `json:"workingType"`         //NO	stopPrice 触发类型: MARK_PRICE(标记价格), CONTRACT_PRICE(合约最新价). 默认 CONTRACT_PRICE
	PriceProtect        *string          `json:"priceProtect"`        //NO	条件单触发保护："TRUE","FALSE", 默认"FALSE". 仅 STOP, STOP_MARKET, TAKE_PROFIT, TAKE_PROFIT_MARKET 需要此参数
	NewClientStrategyId *string          `json:"newClientStrategyId"` //NO	不可以重复出现在挂单中。如空缺系统会自动赋值。必须满足正则规则: ^[\.A-Z\:/a-z0-9_-]{1,32}$
	StopPrice           *decimal.Decimal `json:"stopPrice"`           //NO	Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	ActivationPrice     *decimal.Decimal `json:"activationPrice"`     //NO	TRAILING_STOP_MARKET 单使用，默认标记价格
	CallbackRate        *decimal.Decimal `json:"callbackRate"`        //NO	TRAILING_STOP_MARKET 单使用, 最小0.1, 最大5，1代表1%
	RecvWindow          *int64           `json:"recvWindow"`          //NO
	Timestamp           *int64           `json:"timestamp"`           //YES
}

type PmCConditionalOrderPostApi struct {
	client *PmCRestClient
	req    *PmCConditionalOrderPostReq
}

func (api *PmCConditionalOrderPostApi) Symbol(Symbol string) *PmCConditionalOrderPostApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}

func (api *PmCConditionalOrderPostApi) Side(Side string) *PmCConditionalOrderPostApi {
	api.req.Side = GetPointer(Side)
	return api
}
func (api *PmCConditionalOrderPostApi) PositionSide(PositionSide string) *PmCConditionalOrderPostApi {
	api.req.PositionSide = GetPointer(PositionSide)
	return api
}
func (api *PmCConditionalOrderPostApi) StrategyType(StrategyType string) *PmCConditionalOrderPostApi {
	api.req.StrategyType = GetPointer(StrategyType)
	return api
}
func (api *PmCConditionalOrderPostApi) TimeInForce(TimeInForce string) *PmCConditionalOrderPostApi {
	api.req.TimeInForce = GetPointer(TimeInForce)
	return api
}
func (api *PmCConditionalOrderPostApi) Quantity(Quantity decimal.Decimal) *PmCConditionalOrderPostApi {
	api.req.Quantity = GetPointer(Quantity)
	return api
}
func (api *PmCConditionalOrderPostApi) Price(Price decimal.Decimal) *PmCConditionalOrderPostApi {
	api.req.Price = GetPointer(Price)
	return api
}
func (api *PmCConditionalOrderPostApi) WorkingType(WorkingType string) *PmCConditionalOrderPostApi {
	api.req.WorkingType = GetPointer(WorkingType)
	return api
}
func (api *PmCConditionalOrderPostApi) PriceProtect(PriceProtect string) *PmCConditionalOrderPostApi {
	api.req.PriceProtect = GetPointer(PriceProtect)
	return api
}
func (api *PmCConditionalOrderPostApi) NewClientStrategyId(NewClientStrategyId string) *PmCConditionalOrderPostApi {
	api.req.NewClientStrategyId = GetPointer(NewClientStrategyId)
	return api
}
func (api *PmCConditionalOrderPostApi) StopPrice(StopPrice decimal.Decimal) *PmCConditionalOrderPostApi {
	api.req.StopPrice = GetPointer(StopPrice)
	return api
}
func (api *PmCConditionalOrderPostApi) ActivationPrice(ActivationPrice decimal.Decimal) *PmCConditionalOrderPostApi {
	api.req.ActivationPrice = GetPointer(ActivationPrice)
	return api
}
func (api *PmCConditionalOrderPostApi) CallbackRate(CallbackRate decimal.Decimal) *PmCConditionalOrderPostApi {
	api.req.CallbackRate = GetPointer(CallbackRate)
	return api
}
func (api *PmCConditionalOrderPostApi) RecvWindow(RecvWindow int64) *PmCConditionalOrderPostApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCConditionalOrderPostApi) Timestamp(Timestamp int64) *PmCConditionalOrderPostApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCConditionalOrderDeleteReq struct {
	Symbol              *string `json:"symbol"`               //YES	交易对
	StrategyId          *string `json:"strategyId"`           //NO strategyId 与 newClientStrategyId 之一必须发送
	NewClientStrategyId *string `json:"newClientStrategyId"`  //NO
	RecvWindow          *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp           *int64  `json:"timestamp"`            //YES
}
type PmCConditionalOrderDeleteApi struct {
	client *PmCRestClient
	req    *PmCConditionalOrderDeleteReq
}

func (api *PmCConditionalOrderDeleteApi) Symbol(Symbol string) *PmCConditionalOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCConditionalOrderDeleteApi) StrategyId(StrategyId string) *PmCConditionalOrderDeleteApi {
	api.req.StrategyId = GetPointer(StrategyId)
	return api
}
func (api *PmCConditionalOrderDeleteApi) NewClientStrategyId(NewClientStrategyId string) *PmCConditionalOrderDeleteApi {
	api.req.NewClientStrategyId = GetPointer(NewClientStrategyId)
	return api
}
func (api *PmCConditionalOrderDeleteApi) RecvWindow(RecvWindow int64) *PmCConditionalOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCConditionalOrderDeleteApi) Timestamp(Timestamp int64) *PmCConditionalOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCAllConditionnalOrderDeleteReq struct {
	Symbol     *string `json:"symbol"`               //YES	交易对
	RecvWindow *int64  `json:"recvWindow,omitempty"` //NO
	Timestamp  *int64  `json:"timestamp"`            //YES
}
type PmCAllConditionnalOrderDeleteApi struct {
	client *PmCRestClient
	req    *PmCAllConditionnalOrderDeleteReq
}

func (api *PmCAllConditionnalOrderDeleteApi) Symbol(Symbol string) *PmCAllConditionnalOrderDeleteApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCAllConditionnalOrderDeleteApi) RecvWindow(RecvWindow int64) *PmCAllConditionnalOrderDeleteApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCAllConditionnalOrderDeleteApi) Timestamp(Timestamp int64) *PmCAllConditionnalOrderDeleteApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}

type PmCUserTradesReq struct {
	Symbol     *string `json:"symbol"`     //YES	交易对
	Pair       *string `json:"pair"`       //NO
	StartTime  *int64  `json:"startTime"`  //NO
	EndTime    *int64  `json:"endTime"`    //NO
	FromId     *int64  `json:"fromId"`     //NO 返回该fromId及之后的成交，缺省返回最近的成交
	Limit      *int64  `json:"limit"`      //NO 返回的结果集数量 默认值:50 最大值:1000
	RecvWindow *int64  `json:"recvWindow"` //NO
	Timestamp  *int64  `json:"timestamp"`  //YES
}
type PmCUserTradesApi struct {
	client *PmCRestClient
	req    *PmCUserTradesReq
}

func (api *PmCUserTradesApi) Symbol(Symbol string) *PmCUserTradesApi {
	api.req.Symbol = GetPointer(Symbol)
	return api
}
func (api *PmCUserTradesApi) Pair(Pair string) *PmCUserTradesApi {
	api.req.Pair = GetPointer(Pair)
	return api
}
func (api *PmCUserTradesApi) StartTime(StartTime int64) *PmCUserTradesApi {
	api.req.StartTime = GetPointer(StartTime)
	return api
}
func (api *PmCUserTradesApi) EndTime(EndTime int64) *PmCUserTradesApi {
	api.req.EndTime = GetPointer(EndTime)
	return api
}
func (api *PmCUserTradesApi) FromId(FromId int64) *PmCUserTradesApi {
	api.req.FromId = GetPointer(FromId)
	return api
}
func (api *PmCUserTradesApi) Limit(Limit int64) *PmCUserTradesApi {
	api.req.Limit = GetPointer(Limit)
	return api
}
func (api *PmCUserTradesApi) RecvWindow(RecvWindow int64) *PmCUserTradesApi {
	api.req.RecvWindow = GetPointer(RecvWindow)
	return api
}
func (api *PmCUserTradesApi) Timestamp(Timestamp int64) *PmCUserTradesApi {
	api.req.Timestamp = GetPointer(Timestamp)
	return api
}
