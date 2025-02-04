package model

type OrderType string

const (
	MarketOrder OrderType = "market"
	Limit       OrderType = "limit"
)

func (ot OrderType) String() string {
	return string(ot)
}

type Order struct {
	ID        string
	OrderType OrderType
	Side      TradeSide
	Symbol    Symbol
	Quantity  float64
}
