package market

/**
 لیست قیمت ارز
Order book of a symbol in Binance.
*/
type OrderBook struct {
	LastUpdateId int64
	Bids         []OrderBookEntry
	Asks         []OrderBookEntry
}
