package model

import "errors"

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

func NewSymbolFromString(s string) (Symbol, error) {
	switch s {
	case "BTCUSD":
		return BTCUSD, nil
	default:
		return 0, errors.New("unknown Symbol: " + s)
	}
}
