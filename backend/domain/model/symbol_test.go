package model_test

import (
	"automatic-trade/backend/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		tests := map[string]struct {
			symbol   model.Symbol
			expected string
		}{
			"return BTCUSD":  {symbol: model.BTCUSD, expected: "BTCUSD"},
			"return unknown": {symbol: model.Symbol(99), expected: "unknown"},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				assert.Equal(t, tt.expected, tt.symbol.String())
			})
		}
	})
}
