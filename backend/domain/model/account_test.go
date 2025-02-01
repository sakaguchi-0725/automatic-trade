package model_test

import (
	"automatic-trade/backend/domain/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("new account", func(t *testing.T) {
		tests := map[string]struct {
			equity        float64
			marginBalance float64
			expected      model.Account
			expectedErr   error
		}{
			"returns model.Account on success": {
				equity:        150,
				marginBalance: 150,
				expected: model.Account{
					TotalEquity:        150,
					TotalMarginBalance: 150,
				},
				expectedErr: nil,
			},
			"returns error when total equity is empty": {
				equity:        0,
				marginBalance: 110,
				expected:      model.Account{},
				expectedErr:   errors.New("total equity must be greater than 100.0USD"),
			},
			"returns error when total margin balance is empty": {
				equity:        100,
				marginBalance: 0,
				expected:      model.Account{},
				expectedErr:   errors.New("total margin balance must be greater than 100.0USD"),
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				actual, err := model.NewAccount(tt.equity, tt.marginBalance)

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
