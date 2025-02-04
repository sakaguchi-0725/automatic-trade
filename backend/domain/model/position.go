package model

type OrderStatus string

const (
	Open  OrderStatus = "open"
	Close OrderStatus = "close"
)

func (status OrderStatus) String() string {
	return string(status)
}

type Position struct {
	OrderID     string
	Symbol      Symbol
	Side        TradeSide
	Price       float64
	OrderStatus OrderStatus
	Quantity    float64
}
