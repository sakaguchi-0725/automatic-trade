package model

type TradeSide string

const (
	Sell TradeSide = "sell"
	Buy  TradeSide = "buy"
)

func (ts TradeSide) String() string {
	return string(ts)
}
