package market

/**
 * Represents the best price/qty on the order book for a given symbol.
نمایش بهترین قیمت ارز
*/

type BookTicker struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}
