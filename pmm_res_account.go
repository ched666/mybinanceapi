package mybinanceapi

type PmMAccountRes struct {
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

type PmMAccountBalanceRow struct {
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

type PmMAccountBalanceRes []PmMAccountBalanceRow

type PmMMarginMaxBorrowableRes struct {
	Amount      string `json:"amount"`      //系统可借充足情况下用户账户当前最大可借额度
	BorrowLimit string `json:"borrowLimit"` //平台限制的用户当前等级可以借的额度
}

type PmMMarginMaxTransferableRes struct {
	Anount string `json:"anount"`
}

type PmMMarginInterestHistoryRow struct {
	TxId                int64  `json:"txId"`
	InterestAccuredTime int64  `json:"interestAccuredTime"`
	RawAsset            string `json:"rawAsset"`
	Principal           string `json:"principal"`
	Interest            string `json:"interest"`
	InterestRate        string `json:"interestRate"`
	Type                string `json:"type"`
}

type PmMMarginInterestHistoryRes struct {
	Rows  []PmMMarginInterestHistoryRow `json:"rows"`
	Total int                           `json:"total"`
}
