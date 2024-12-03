package mybinanceapi

type PmUAccountForUResAsset struct {
	Asset                  string `json:"asset"`                  //资产
	CrossWalletBalance     string `json:"crossWalletBalance"`     //全仓账户余额
	CrossUnPnl             string `json:"crossUnPnl"`             // 全仓持仓未实现盈亏
	MaintMargin            string `json:"maintMargin"`            // 维持保证金
	InitialMargin          string `json:"initialMargin"`          // 当前所需起始保证金
	PositionInitialMargin  string `json:"positionInitialMargin"`  // 持仓所需起始保证金(基于最新标记价格)
	OpenOrderInitialMargin string `json:"openOrderInitialMargin"` // 当前挂单所需起始保证金(基于最新标记价格)
	UpdateTime             int64  `json:"updateTime"`             //更新时间
}
type PmUAccountForUResPosition struct {
	//根据用户持仓模式展示持仓方向，即单向模式下只返回BOTH持仓情况，双向模式下只返回 LONG 和 SHORT 持仓情况
	Symbol           string `json:"symbol"`           // 交易对
	InitialMargin    string `json:"initialMargin"`    // 当前所需起始保证金(基于最新标记价格)
	MaintMargin      string `json:"maintMargin"`      //维持保证金
	UnrealizedProfit string `json:"unrealizedProfit"` // 持仓未实现盈亏
	PositionSide     string `json:"positionSide"`     // 持仓方向
	PositionAmt      string `json:"positionAmt"`      // 持仓数量
	UpdateTime       int64  `json:"updateTime"`       // 更新时间
	NotionalValue    string `json:"notional"`         // 净值
}
type PmUAccountForURes struct {
	Assets    []PmUAccountForUResAsset    `json:"assets"`
	Positions []PmUAccountForUResPosition `json:"positions"` // 头寸，将返回所有市场symbol。
}

type PmUAccountRes struct {
	UniMMR                   string `json:"uniMMR"`                   // 统一账户维持保证金率
	AccountEquity            string `json:"accountEquity"`            // 以USD计价的账户权益
	ActualEquity             string `json:"actualEquity"`             // 不考虑质押率的以USD计价账户权益
	AccountInitialMargin     string `json:"accountInitialMargin"`     // 当前所需起始保证金(基于最新标记价格)
	AccountMaintMargin       string `json:"accountMaintMargin"`       // 以USD计价统一账户维持保证金
	AccountStatus            string `json:"accountStatus"`            // 统一账户账户状态："NORMAL", "MARGIN_CALL", "SUPPLY_MARGIN", "REDUCE_ONLY", "ACTIVE_LIQUIDATION", "FORCE_LIQUIDATION", "BANKRUPTED"
	VirtualMaxWithdrawAmount string `json:"virtualMaxWithdrawAmount"` // 以USD计价的最大可转出
	TotalAvailableBalance    string `json:"totalAvailableBalance"`    // 以USD计价的可用余额
	TotalMarginOpenLoss      string `json:"totalMarginOpenLoss"`      // 以USD计价的总开仓保证金
	UpdateTime               int64  `json:"updateTime"`               // 更新时间
}

type PmUAccountBalanceRow struct {
	Asset               string `json:"asset"`               //资产
	TotalWalletBalance  string `json:"totalWalletBalance"`  // 钱包余额 =  全仓杠杆未锁定 + 全仓杠杆锁定 + u本位合约钱包余额 + 币本位合约钱包余额
	CrossMarginAsset    string `json:"crossMarginAsset"`    // 全仓资产 = 全仓杠杆未锁定 + 全仓杠杆锁定
	CrossMarginBorrowed string `json:"crossMarginBorrowed"` // 全仓杠杆借贷
	CrossMarginFree     string `json:"crossMarginFree"`     // 全仓杠杆未锁定
	CrossMarginInterest string `json:"crossMarginInterest"` // 全仓杠杆利息
	CrossMarginLocked   string `json:"crossMarginLocked"`   //全仓杠杆锁定
	UmWalletBalance     string `json:"umWalletBalance"`     // u本位合约钱包余额
	UmUnrealizedPNL     string `json:"umUnrealizedPNL"`     // u本位未实现盈亏
	CmWalletBalance     string `json:"cmWalletBalance"`     // 币本位合约钱包余额
	CmUnrealizedPNL     string `json:"cmUnrealizedPNL"`     // 币本位未实现盈亏
	UpdateTime          int64  `json:"updateTime"`          // 更新时间
	NegativeBalance     string `json:"negativeBalance"`     // 负资产余额
}

type PmUAccountBalanceRes []PmUAccountBalanceRow

type PmUCommissionRateRes struct {
	Symbol              string `json:"symbol"`
	MakerCommissionRate string `json:"makerCommissionRate"`
	TakerCommissionRate string `json:"takerCommissionRate"`
}
