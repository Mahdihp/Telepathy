package account

/**
 * An asset balance in an Account.
 *  دارایی حساب
 * @see Account
 */

type AssetBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
