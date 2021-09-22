package account

/**
* Account information.
اطلاعات حساب
*/

type Account struct {
	MakerCommission  int
	TakerCommission  int
	BuyerCommission  int
	SellerCommission int
	CanTrade         bool
	CanWithdraw      bool
	CanDeposit       bool
	balances         []AssetBalance
}
