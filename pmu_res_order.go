package mybinanceapi

type PmUOrderOrder struct {
	AvgPrice                string `json:"avgPrice"`                  // 平均成交价
	ClientOrderId           string `json:"clientOrderId"`             // 用户自定义的订单号
	CumQuote                string `json:"cumQuote"`                  // 成交金额
	ExecutedQty             string `json:"executedQty"`               // 成交量
	OrderId                 int64  `gorm:"primaryKey" json:"orderId"` // 系统订单号
	OrigQty                 string `json:"origQty"`                   // 原始委托数量
	OrigType                string `json:"origType"`                  // 触发前订单类型
	Price                   string `json:"price"`                     // 委托价格
	ReduceOnly              bool   `json:"reduceOnly"`                // 是否仅减仓
	Side                    string `json:"side"`                      // 买卖方向
	PositionSide            string `json:"positionSide"`              // 持仓方向
	Status                  string `json:"status"`                    // 订单状态
	Symbol                  string `json:"symbol"`                    // 交易对
	Time                    int64  `json:"time"`                      // 订单时间
	TimeInForce             string `json:"timeInForce"`               // 有效方法
	Type                    string `json:"type"`                      // 订单类型
	UpdateTime              int64  `json:"updateTime"`                // 更新时间
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`   // 是否开启条件单触发保护
	GoodTillDate            int64  `json:"goodTillDate"`              // 订单TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`                // 触发条件
}

type PmUOpenOrdersRes []PmUOrderOrder

type PmUAllOrdersRes []PmUOrderOrder

type PmUOrderPostRes struct {
	ClientOrderId           string `json:"clientOrderId"` // 用户自定义的订单号
	CumQty                  string `json:"cumQty"`
	CumQuote                string `json:"cumQuote"` // 成交金额
	ExecutedQty             string `json:"executedQty"`
	OrderId                 int64  `json:"orderId"` // 系统订单号
	AvgPrice                string `json:"avgPrice"`
	OrigQty                 string `json:"origQty"` // 原始委托数量
	Price                   string `json:"price"`   // 委托价格
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`                    // 买卖方向
	PositionSide            string `json:"positionSide"`            // 持仓方向
	Status                  string `json:"status"`                  // 订单状态
	Symbol                  string `json:"symbol"`                  // 交易对
	TimeInForce             string `json:"timeInForce"`             // 有效方法
	Type                    string `json:"type"`                    // 订单类型
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` // 订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            // 订单TIF为GTD时的自动取消时间
	UpdateTime              int64  `json:"updateTime"`              // 更新时间
	PriceMatch              string `json:"priceMatch"`              // 触发条件
}

type PmUOrderPutRes struct {
	OrderId                 int64  `json:"orderId"`
	Symbol                  string `json:"symbol"`
	Pair                    string `json:"pair"`
	Status                  string `json:"status"`
	ClientOrderId           string `json:"clientOrderId"`
	Price                   string `json:"price"`
	AvgPrice                string `json:"avgPrice"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	CumQty                  string `json:"cumQty"`
	CumBase                 string `json:"cumBase"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	ReduceOnly              bool   `json:"reduceOnly"`
	ClosePosition           bool   `json:"closePosition"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	OrigType                string `json:"origType"`
	PriceMatch              string `json:"priceMatch"`              //盘口价格下单模式 与price只可同时传一个
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` //订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            //订单TIF为GTD时的自动取消时间
	UpdateTime              int64  `json:"updateTime"`
}

type PmUOrderGetRes PmUOrderOrder

type PmUOrderDeleteRes PmUOrderOrder

type PmUAllOrderDeleteRes BinanceErrorRes

type PmUConditionalOrderPostRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`
	StrategyId              int64  `json:"strategyId"`
	StrategyStatus          string `json:"strategyStatus"`
	StrategyType            string `json:"strategyType"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	Symbol                  string `json:"symbol"`
	TimeInForce             string `json:"timeInForce"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	BookTime                int64  `json:"bookTime"` //条件单下单时间
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` ////订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            //订单TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`
}

type PmUConditionalOrderDeleteRes struct {
	NewClientStrategyId     string `json:"newClientStrategyId"`
	StrategyId              int64  `json:"strategyId"`
	StrategyStatus          string `json:"strategyStatus"`
	StrategyType            string `json:"strategyType"`
	OrigQty                 string `json:"origQty"`
	Price                   string `json:"price"`
	ReduceOnly              bool   `json:"reduceOnly"`
	Side                    string `json:"side"`
	PositionSide            string `json:"positionSide"`
	StopPrice               string `json:"stopPrice"`
	Symbol                  string `json:"symbol"`
	TimeInForce             string `json:"timeInForce"`
	ActivatePrice           string `json:"activatePrice"`
	PriceRate               string `json:"priceRate"`
	BookTime                int64  `json:"bookTime"` //条件单下单时间
	UpdateTime              int64  `json:"updateTime"`
	WorkingType             string `json:"workingType"`
	PriceProtect            bool   `json:"priceProtect"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"` ////订单自成交保护模式
	GoodTillDate            int64  `json:"goodTillDate"`            //订单TIF为GTD时的自动取消时间
	PriceMatch              string `json:"priceMatch"`
}

type PmUAllConditionnalOrderDeleteRes BinanceErrorRes

type PmUUserTradesRes []PmUUserTradesOrder

type PmUUserTradesOrder struct {
	Symbol          string `json:"symbol"`
	Id              int64  `json:"id"`
	OrderId         int64  `json:"orderId"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
	PositionSide    string `json:"positionSide"`
}
