package model_test

import (
	"automatic-trade/backend/core/testutil"
	"automatic-trade/backend/domain/model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMarket(t *testing.T) {
	timeFactory := testutil.NewFixedDateTimeFactory(2025, time.February, 1)

	t.Run("set", func(t *testing.T) {
		tests := map[string]struct {
			input    model.Rates
			expected model.Rates
		}{
			"sorted by date descending": {
				input: model.Rates{
					{DateTime: timeFactory.At(10, 20), Price: 1000},
					{DateTime: timeFactory.At(10, 15), Price: 1000.4},
					{DateTime: timeFactory.At(10, 5), Price: 1001},
					{DateTime: timeFactory.At(10, 10), Price: 1000.6},
				},
				expected: model.Rates{
					{DateTime: timeFactory.At(10, 20), Price: 1000},
					{DateTime: timeFactory.At(10, 15), Price: 1000.4},
					{DateTime: timeFactory.At(10, 10), Price: 1000.6},
					{DateTime: timeFactory.At(10, 5), Price: 1001},
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				market := model.Market{
					Symbol:   model.BTCUSD,
					Interval: model.Min5,
				}

				market.Set(tt.input)
				assert.Equal(t, tt.expected, market.Rates)
			})
		}
	})

	t.Run("add", func(t *testing.T) {
		tests := map[string]struct {
			market   model.Market
			input    model.Rate
			expected model.Rates
		}{
			"latest rate is added to index 0 and the oldest rate is deleted": {
				market: model.Market{
					Symbol:   model.BTCUSD,
					Interval: model.Min5,
					Rates: model.Rates{
						{DateTime: timeFactory.At(10, 25), Price: 1000.5},
						{DateTime: timeFactory.At(10, 20), Price: 1000.6},
						{DateTime: timeFactory.At(10, 15), Price: 1000.4},
						{DateTime: timeFactory.At(10, 10), Price: 1001},
						{DateTime: timeFactory.At(10, 5), Price: 1000.8},
					},
				},
				input: model.Rate{DateTime: timeFactory.At(10, 30), Price: 1001},
				expected: model.Rates{
					{DateTime: timeFactory.At(10, 30), Price: 1001},
					{DateTime: timeFactory.At(10, 25), Price: 1000.5},
					{DateTime: timeFactory.At(10, 20), Price: 1000.6},
					{DateTime: timeFactory.At(10, 15), Price: 1000.4},
					{DateTime: timeFactory.At(10, 10), Price: 1001},
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				tt.market.Add(tt.input)
				assert.Equal(t, tt.expected, tt.market.Rates)
			})
		}
	})
}
