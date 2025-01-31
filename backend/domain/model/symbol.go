package model

type Symbol int

const (
	BTCUSD Symbol = iota + 1
)

func (s Symbol) String() string {
	switch s {
	case BTCUSD:
		return "BTCUSD"
	default:
		return "unknown"
	}
}
