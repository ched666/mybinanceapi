package mybinanceapi

type PmMMarginOrder struct {
	ClientOrderId           string `json:"clientOrderId"`             // 用户自定义的订单号
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`       // 成交金额
	ExecutedQty             string `json:"executedQty"`               // 成交量
	IcebergQty              string `json:"icebergQty"`                // 冰山数量
	IsWorking               bool   `json:"isWorking"`                 // 是否生效
	OrderId                 int64  `gorm:"primaryKey" json:"orderId"` // 系统订单号
	OrigQty                 string `json:"origQty"`                   // 原始委托数量
	Price                   string `json:"price"`                     // 委托价格
	Side                    string `json:"side"`                      // 买卖方向
	Status                  string `json:"status"`                    // 订单状态
	StopPrice               string `json:"stopPrice"`                 // 触发价格
	Symbol                  string `json:"symbol"`                    // 交易对
	Time                    int64  `json:"time"`                      // 订单时间
	TimeInForce             string `json:"timeInForce"`               // 有效方法
	Type                    string `json:"type"`                      // 订单类型
	UpdateTime              int64  `json:"updateTime"`                // 更新时间
	AccountId               int64  `json:"accountId"`                 // 账户ID
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`   // 是否开启条件单触发保护
	PreventedMatchId        string `json:"preventedMatchId"`          // 触发条件单的订单号
	PreventedQuantity       string `json:"preventedQuantity"`         // 触发条件单的成交量
}

type PmMMarginOrderGetRes PmMMarginOrder

type PmMMarginOrderPostRow struct {
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
}

type PmMMarginOrderPostRes struct {
	Symbol                string                  `json:"symbol"`
	OrderId               int64                   `json:"orderId"`
	ClientOrderId         string                  `json:"clientOrderId"`
	TransactTime          int64                   `json:"transactTime"`
	Price                 string                  `json:"price"`
	OrigQty               string                  `json:"origQty"`
	ExecutedQty           string                  `json:"executedQty"`
	CummulativeQuoteQty   string                  `json:"cummulativeQuoteQty"`
	Status                string                  `json:"status"`
	TimeInForce           string                  `json:"timeInForce"`
	Type                  string                  `json:"type"`
	Side                  string                  `json:"side"`
	MarginBuyBorrowAmount string                  `json:"marginBuyBorrowAmount"`
	MarginBuyBorrowAsset  string                  `json:"marginBuyBorrowAsset"`
	PmMMarginOrderPostRow []PmMMarginOrderPostRow `json:"fills"`
}

type PmMMarginOrderDeleteRes struct {
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}
type PmMMarginAllOrdersRes []PmMMarginOrder

type PmMMarginOpenOrdersRes []PmMMarginOrder

type PmPOCOOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type PmMMarginOCOOrderGetRes struct {
	OrderListId       int64         `json:"orderListId"`
	ContingencyType   string        `json:"contingencyType"`
	ListStatusType    string        `json:"listStatusType"`
	ListOrderStatus   string        `json:"listOrderStatus"`
	ListClientOrderId string        `json:"listClientOrderId"`
	TransactionTime   int64         `json:"transactionTime"`
	Symbol            string        `json:"symbol"`
	Orders            []PmPOCOOrder `json:"orders"`
}

type PmPOCOOrderReport struct {
	Symbol              string `json:"symbol"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	TransactTime        int64  `json:"transactTime"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
}
type PmMMarginOCOOrderPostRes struct {
	OrderListId           int64               `json:"orderListId"`
	ContingencyType       string              `json:"contingencyType"`
	ListStatusType        string              `json:"listStatusType"`
	ListOrderStatus       string              `json:"listOrderStatus"`
	ListClientOrderId     string              `json:"listClientOrderId"`
	TransactionTime       int64               `json:"transactionTime"`
	Symbol                string              `json:"symbol"`
	MarginBuyBorrowAmount string              `json:"marginBuyBorrowAmount"` // 下单后没有发生借款则不返回该字段
	MarginBuyBorrowAsset  string              `json:"marginBuyBorrowAsset"`  // 下单后没有发生借款则不返回该字段
	Orders                []PmPOCOOrder       `json:"orders"`
	OrderReports          []PmPOCOOrderReport `json:"orderReports"`
}

type PmPOCODeleteOrderReport struct {
	Symbol              string `json:"symbol"`
	OrigClientOrderId   string `json:"origClientOrderId"`
	OrderId             int64  `json:"orderId"`
	OrderListId         int64  `json:"orderListId"`
	ClientOrderId       string `json:"clientOrderId"`
	Price               string `json:"price"`
	OrigQty             string `json:"origQty"`
	ExecutedQty         string `json:"executedQty"`
	CummulativeQuoteQty string `json:"cummulativeQuoteQty"`
	Status              string `json:"status"`
	TimeInForce         string `json:"timeInForce"`
	Type                string `json:"type"`
	Side                string `json:"side"`
	StopPrice           string `json:"stopPrice"`
}
type PmMMarginOCOOrderDeleteRes struct {
	OrderListId       int64                     `json:"orderListId"`
	ContingencyType   string                    `json:"contingencyType"`
	ListStatusType    string                    `json:"listStatusType"`
	ListOrderStatus   string                    `json:"listOrderStatus"`
	ListClientOrderId string                    `json:"listClientOrderId"`
	TransactionTime   int64                     `json:"transactionTime"`
	Symbol            string                    `json:"symbol"`
	Orders            []PmPOCOOrder             `json:"orders"`
	OrderReports      []PmPOCODeleteOrderReport `json:"orderReports"`
}

type PmMMarginOCOAllOrdersRes []PmMMarginOCOOrderGetRes

type PmMMarginOCOOpenOrdersRes []PmMMarginOCOOrderGetRes
