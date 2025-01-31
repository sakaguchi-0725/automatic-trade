package model_test

import (
	"automatic-trade/backend/core/testutil"
	"automatic-trade/backend/domain/model"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRate(t *testing.T) {
	timeFactory := testutil.NewFixedDateTimeFactory(2025, time.February, 1)

	t.Run("Latest", func(t *testing.T) {
		tests := map[string]struct {
			rates       model.Rates
			expected    model.Rate
			expectedErr error
		}{
			"returns model.Rate on success": {
				rates: model.Rates{
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 0), Price: 10.5},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 5), Price: 10.8},
					{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 10), Price: 10.6},
				},
				expected:    model.Rate{Symbol: model.BTCUSD, DateTime: timeFactory.At(10, 10), Price: 10.6},
				expectedErr: nil,
			},
			"returns error when rates is empty": {
				rates:       model.Rates{},
				expected:    model.Rate{},
				expectedErr: errors.New("no rates available"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				result, err := tt.rates.Latest()

				assert.Equal(t, tt.expected, result)

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
