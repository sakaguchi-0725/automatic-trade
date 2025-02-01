package model_test

import (
	"automatic-trade/backend/core/testutil"
	"automatic-trade/backend/domain/model"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTrade(t *testing.T) {
	timeFactory := testutil.NewFixedDateTimeFactory(2025, time.February, 2)

	t.Run("new trade", func(t *testing.T) {
		tests := map[string]struct {
			symbol      model.Symbol
			higherRates model.Rates
			lowerRates  model.Rates
			expected    *model.Trade
			expectedErr error
		}{
			"returns *model.Trade on success": {
				symbol: model.BTCUSD,
				higherRates: model.Rates{
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 0), Price: 102.1},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 0), Price: 103.0},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
				},
				lowerRates: model.Rates{
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 50), Price: 102.3},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 55), Price: 102.5},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
				},
				expected: &model.Trade{
					Symbol: model.BTCUSD,
					HigherTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 0), Price: 102.1},
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 0), Price: 103.0},
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
					LowerTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 50), Price: 102.3},
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 55), Price: 102.5},
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				expectedErr: nil,
			},
			"returns error when higher timeframe rates is empty": {
				symbol:      model.BTCUSD,
				higherRates: model.Rates{},
				lowerRates: model.Rates{
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 50), Price: 102.3},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 55), Price: 102.5},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
				},
				expected:    nil,
				expectedErr: errors.New("higher timeframe rates cannot be empty"),
			},
			"returns error when lower timeframe rates is empty": {
				symbol: model.BTCUSD,
				higherRates: model.Rates{
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 0), Price: 102.1},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(11, 0), Price: 103.0},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
				},
				lowerRates:  model.Rates{},
				expected:    nil,
				expectedErr: errors.New("lower timeframe rates cannot be empty"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				actual, err := model.NewTrade(tt.symbol, tt.higherRates, tt.lowerRates)

				assert.Equal(t, tt.expected, actual)
				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("calculate quantity", func(t *testing.T) {
		tests := map[string]struct {
			input       float64
			expected    float64
			expectedErr error
		}{
			"returns no error set quantity on success": {input: 101.0, expected: 10.1, expectedErr: nil},
			"returns error when less than 100USD":      {input: 99.9, expected: 0, expectedErr: errors.New("totalWallet must be greater than 100.0USD")},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				trade := &model.Trade{
					HigherTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
					LowerTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				}

				err := trade.CalculateQuantity(tt.input)

				assert.Equal(t, tt.expected, trade.Quantity)
				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})

	t.Run("make position", func(t *testing.T) {
		tests := map[string]struct {
			execPrice   float64
			expected    model.Position
			expectedErr error
		}{
			"returns position on success": {
				execPrice: 101.2,
				expected: model.Position{
					Symbol:      model.BTCUSD,
					Side:        model.Buy,
					Price:       101.2,
					OrderStatus: model.Open,
					Quantity:    120,
				},
				expectedErr: nil,
			},
			"returns error when invalid price": {
				execPrice:   0,
				expected:    model.Position{},
				expectedErr: errors.New("invalid execution price"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				trade := &model.Trade{
					Symbol: model.BTCUSD,
					Side:   model.Buy,
					HigherTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
					LowerTimeFrameRates: model.Rates{
						{Symbol: model.BTCUSD, DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
					Quantity: 120,
				}

				actual, err := trade.MakePosition(tt.execPrice)

				assert.Equal(t, tt.expected, actual)
				if tt.expectedErr != nil {
					assert.Error(t, err)
					assert.Equal(t, tt.expectedErr, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}
