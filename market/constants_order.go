package market

type OrderType int

const (
	MARKET OrderType = 0
	LIMIT  OrderType = 1
)

type OrderSide int

const (
	BUY  OrderSide = 0
	SELL OrderSide = 1
)

type OrderStatus int

const (
	NEW              OrderStatus = 0
	PARTIALLY_FILLED OrderStatus = 1
	FILLED           OrderStatus = 2
	CANCELED         OrderStatus = 3
	PENDING_CANCEL   OrderStatus = 4
	REJECTED         OrderStatus = 5
	EXPIRED          OrderStatus = 6
)

type TimeInForce int

const (
	GTC TimeInForce = 0
	IOC TimeInForce = 1
)
