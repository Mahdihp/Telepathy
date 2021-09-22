package market

/**
 * Order book of a symbol in Binance.
قیمت ارز و تعداد
*/

type OrderBookEntry struct {
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
