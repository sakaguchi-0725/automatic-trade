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
			higher      model.Market
			lower       model.Market
			expected    *model.Trade
			expectedErr error
		}{
			"returns *model.Trade on success": {
				higher: model.Market{
					Symbol: model.BTCUSD,
					Rates: model.Rates{
						{DateTime: timeFactory.At(10, 0), Price: 102.1},
						{DateTime: timeFactory.At(11, 0), Price: 103.0},
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				lower: model.Market{
					Symbol: model.BTCUSD,
					Rates: model.Rates{
						{DateTime: timeFactory.At(11, 50), Price: 102.3},
						{DateTime: timeFactory.At(11, 55), Price: 102.5},
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				expected: &model.Trade{
					HigherTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(10, 0), Price: 102.1},
							{DateTime: timeFactory.At(11, 0), Price: 103.0},
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
					},
					LowerTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(11, 50), Price: 102.3},
							{DateTime: timeFactory.At(11, 55), Price: 102.5},
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
					},
				},
				expectedErr: nil,
			},
			"returns error when higher and lower timeframe symbols mismatch": {
				higher: model.Market{
					Symbol: model.Symbol(99),
					Rates: model.Rates{
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				lower: model.Market{
					Symbol: model.BTCUSD,
					Rates: model.Rates{
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				expected:    nil,
				expectedErr: errors.New("higher and lower time frame symbols must match"),
			},
			"returns error when higher timeframe rates is empty": {
				higher: model.Market{Symbol: model.BTCUSD},
				lower: model.Market{
					Symbol: model.BTCUSD,
					Rates: model.Rates{
						{DateTime: timeFactory.At(11, 50), Price: 102.3},
						{DateTime: timeFactory.At(11, 55), Price: 102.5},
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				expected:    nil,
				expectedErr: errors.New("higher timeframe rates cannot be empty"),
			},
			"returns error when lower timeframe rates is empty": {
				higher: model.Market{
					Symbol: model.BTCUSD,
					Rates: model.Rates{
						{DateTime: timeFactory.At(10, 0), Price: 102.1},
						{DateTime: timeFactory.At(11, 0), Price: 103.0},
						{DateTime: timeFactory.At(12, 0), Price: 102.8},
					},
				},
				lower:       model.Market{Symbol: model.BTCUSD},
				expected:    nil,
				expectedErr: errors.New("lower timeframe rates cannot be empty"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				actual, err := model.NewTrade(tt.higher, tt.lower)

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
					HigherTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
					},
					LowerTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
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

	t.Run("make order", func(t *testing.T) {
		tests := map[string]struct {
			trade    model.Trade
			expected model.Order
		}{
			"returns model.Order on success": {
				trade: model.Trade{
					Side: model.Buy,
					HigherTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
					},
					LowerTimeFrame: model.Market{
						Symbol: model.BTCUSD,
						Rates: model.Rates{
							{DateTime: timeFactory.At(12, 0), Price: 102.8},
						},
					},
					Quantity: 120,
				},
				expected: model.Order{
					ID:        "",
					OrderType: model.MarketOrder,
					Side:      model.Buy,
					Symbol:    model.BTCUSD,
					Quantity:  120,
				},
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				actual := tt.trade.MakeOrder()

				assert.Equal(t, tt.expected, actual)
			})
		}
	})
}
