package model

import "errors"

type Account struct {
	TotalEquity        float64
	TotalMarginBalance float64
}

func NewAccount(equity, marginBalance float64) (Account, error) {
	if equity < 100 {
		return Account{}, errors.New("total equity must be greater than 100.0USD")
	}
	if marginBalance < 100 {
		return Account{}, errors.New("total margin balance must be greater than 100.0USD")
	}

	return Account{
		TotalEquity:        equity,
		TotalMarginBalance: marginBalance,
	}, nil
}
