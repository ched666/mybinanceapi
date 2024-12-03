package mybinanceapi

type PmCOrderOrder struct {
	AvgPrice      string `json:"avgPrice"`                  // 平均成交价
	ClientOrderId string `json:"clientOrderId"`             // 用户自定义的订单号
	CumBase       string `json:"cumBase"`                   // 成交数量
	ExecutedQty   string `json:"executedQty"`               // 成交量
	OrderId       int64  `gorm:"primaryKey" json:"orderId"` // 系统订单号
	OrigQty       string `json:"origQty"`                   // 原始委托数量
	OrigType      string `json:"origType"`                  // 触发前订单类型
	Price         string `json:"price"`                     // 委托价格
	ReduceOnly    bool   `json:"reduceOnly"`                // 是否仅减仓
	Side          string `json:"side"`                      // 买卖方向
	PositionSide  string `json:"positionSide"`              // 持仓方向
	Status        string `json:"status"`                    // 订单状态
	Symbol        string `json:"symbol"`                    // 交易对
	Pair          string `json:"pair"`
	Time          int64  `json:"time"`        // 订单时间
	TimeInForce   string `json:"timeInForce"` // 有效方法
	Type          string `json:"type"`        // 订单类型
	UpdateTime    int64  `json:"updateTime"`  // 更新时间
}

type PmCOpenOrdersRes []PmCOrderOrder

type PmCAllOrdersRes []PmCOrderOrder

type PmCOrderPostRes struct {
	ClientOrderId string `json:"clientOrderId"` // 用户自定义的订单号
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"` // 成交数量
	ExecutedQty   string `json:"executedQty"`
	OrderId       int64  `json:"orderId"` // 系统订单号
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"` // 原始委托数量
	Price         string `json:"price"`   // 委托价格
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`         // 买卖方向
	PositionSide  string `json:"positionSide"` // 持仓方向
	Status        string `json:"status"`       // 订单状态
	Symbol        string `json:"symbol"`       // 交易对
	Pair          string `json:"pair"`
	TimeInForce   string `json:"timeInForce"` // 有效方法
	Type          string `json:"type"`        // 订单类型
	UpdateTime    int64  `json:"updateTime"`  // 更新时间
}

type PmCOrderPutRes struct {
	OrderId       int64  `json:"orderId"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	Status        string `json:"status"`
	ClientOrderId string `json:"clientOrderId"`
	Price         string `json:"price"`
	AvgPrice      string `json:"avgPrice"`
	OrigQty       string `json:"origQty"`
	ExecutedQty   string `json:"executedQty"`
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	OrigType      string `json:"origType"`
	UpdateTime    int64  `json:"updateTime"`
}

type PmCOrderGetRes PmCOrderOrder

type PmCOrderDeleteRes struct {
	AvgPrice      string `json:"avgPrice"`
	ClientOrderId string `json:"clientOrderId"`
	CumQty        string `json:"cumQty"`
	CumBase       string `json:"cumBase"`
	ExecutedQty   string `json:"executedQty"`
	OrderId       int64  `json:"orderId"`
	OrigQty       string `json:"origQty"`
	Price         string `json:"price"`
	ReduceOnly    bool   `json:"reduceOnly"`
	Side          string `json:"side"`
	PositionSide  string `json:"positionSide"`
	Status        string `json:"status"`
	Symbol        string `json:"symbol"`
	Pair          string `json:"pair"`
	TimeInForce   string `json:"timeInForce"`
	Type          string `json:"type"`
	UpdateTime    int64  `json:"updateTime"`
}

type PmCAllOrderDeleteRes BinanceErrorRes

type PmCConditionalOrderPostRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"`
	StrategyId          int64  `json:"strategyId"`
	StrategyStatus      string `json:"strategyStatus"`
	StrategyType        string `json:"strategyType"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	ReduceOnly          bool   `json:"reduceOnly"`
	Side                string `json:"side"`
	PositionSide        string `json:"positionSide"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	Pair                string `json:"pair"`
	TimeInForce         string `json:"timeInForce"`
	ActivatePrice       string `json:"activatePrice"`
	PriceRate           string `json:"priceRate"`
	BookTime            int64  `json:"bookTime"` //条件单下单时间
	UpdateTime          int64  `json:"updateTime"`
	WorkingType         string `json:"workingType"`
	PriceProtect        bool   `json:"priceProtect"`
}

type PmCConditionalOrderDeleteRes struct {
	NewClientStrategyId string `json:"newClientStrategyId"`
	StrategyId          int64  `json:"strategyId"`
	StrategyStatus      string `json:"strategyStatus"`
	StrategyType        string `json:"strategyType"`
	OrigQty             string `json:"origQty"`
	Price               string `json:"price"`
	ReduceOnly          bool   `json:"reduceOnly"`
	Side                string `json:"side"`
	PositionSide        string `json:"positionSide"`
	StopPrice           string `json:"stopPrice"`
	Symbol              string `json:"symbol"`
	TimeInForce         string `json:"timeInForce"`
	ActivatePrice       string `json:"activatePrice"`
	PriceRate           string `json:"priceRate"`
	UpdateTime          int64  `json:"updateTime"`
	WorkingType         string `json:"workingType"`
	PriceProtect        bool   `json:"priceProtect"`
}

type PmCAllConditionnalOrderDeleteRes BinanceErrorRes

type PmCUserTradesRes []PmCUserTradesOrder

type PmCUserTradesOrder struct {
	Symbol          string `json:"symbol"`
	Id              int64  `json:"id"`
	OrderId         int64  `json:"orderId"`
	Pair            string `json:"pair"`
	Side            string `json:"side"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	RealizedPnl     string `json:"realizedPnl"`
	MarginAsset     string `json:"marginAsset"`
	BaseQty         string `json:"baseQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	PositionSide    string `json:"positionSide"`
	Buyer           bool   `json:"buyer"`
	Maker           bool   `json:"maker"`
}
