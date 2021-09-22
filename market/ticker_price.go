package market

/**
Wraps a symbol and its corresponding latest price.
قیمت لحطه ای ارز
*/
type TickerPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
