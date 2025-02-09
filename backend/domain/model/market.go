package model

import (
	"sort"
)

type Market struct {
	Symbol   Symbol
	Interval Interval
	Rates    Rates
}

func (m *Market) Set(rates Rates) {
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].DateTime.After(rates[j].DateTime)
	})

	m.Rates = rates
}

func (m *Market) Add(rate Rate) {
	m.Rates = append(Rates{rate}, m.Rates...)
	m.Rates = m.Rates[:len(m.Rates)-1]
}
