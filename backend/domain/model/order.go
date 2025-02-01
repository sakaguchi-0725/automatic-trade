package model

type OrderType string

const (
	MarketOrder OrderType = "market"
	Limit       OrderType = "limit"
)

type Order struct {
	ID        string
	OrderType OrderType
	Side      TradeSide
	Symbol    Symbol
	Quantity  float64
}
