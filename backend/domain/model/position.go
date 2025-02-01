package model

type OrderStatus string

const (
	Open  OrderStatus = "open"
	Close OrderStatus = "close"
)

type Position struct {
	Symbol      Symbol
	Side        TradeSide
	Price       float64
	OrderStatus OrderStatus
	Quantity    float64
}
