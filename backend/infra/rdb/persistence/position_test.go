package persistence_test

import (
	"automatic-trade/backend/core/apperr"
	"automatic-trade/backend/domain/model"
	"automatic-trade/backend/infra/rdb/dto"
	"automatic-trade/backend/infra/rdb/persistence"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPosition(t *testing.T) {
	positionRepo := persistence.NewPositionRepository(testDB)

	t.Run("delete", func(t *testing.T) {
		defer cleanupTestDB()

		tests := map[string]struct {
			input       string
			expectedErr error
		}{
			"returns no error when success":         {input: "order-1", expectedErr: nil},
			"returns error when not exists orderID": {input: "order-99", expectedErr: errors.New("position with orderID order-99 not found")},
		}

		t.Run("create position", func(t *testing.T) {
			position := dto.Position{
				OrderID:     "order-1",
				Symbol:      model.BTCUSD.String(),
				Side:        model.Buy.String(),
				Price:       1002.2,
				OrderStatus: model.Open.String(),
				Quantity:    200,
			}

			err := testDB.Create(&position).Error
			require.NoError(t, err)
		})

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := positionRepo.Delete(tt.input)
				assert.Equal(t, tt.expectedErr, err)
			})
		}
	})

	t.Run("get", func(t *testing.T) {
		defer cleanupTestDB()

		tests := map[string]struct {
			input       string
			expected    model.Position
			expectedErr error
		}{
			"returns model.Position when success": {
				input: "order-1",
				expected: model.Position{
					OrderID:     "order-1",
					Symbol:      model.BTCUSD,
					Side:        model.Sell,
					Price:       1000,
					OrderStatus: model.Open,
					Quantity:    300,
				},
				expectedErr: nil,
			},
			"returns error when not exists orderID": {
				input:       "order-99",
				expected:    model.Position{},
				expectedErr: apperr.ErrDataNotFound,
			},
			"returns error when invalid data": {
				input:       "order-2",
				expected:    model.Position{},
				expectedErr: errors.New("unknown Symbol: HOGE"),
			},
		}

		t.Run("create position", func(t *testing.T) {
			position := []dto.Position{
				{
					OrderID:     "order-1",
					Symbol:      model.BTCUSD.String(),
					Side:        model.Sell.String(),
					Price:       1000,
					OrderStatus: model.Open.String(),
					Quantity:    300,
				}, {
					OrderID:     "order-2",
					Symbol:      "HOGE", // invalid symbol
					Side:        model.Sell.String(),
					Price:       1000,
					OrderStatus: model.Open.String(),
					Quantity:    300,
				},
			}

			err := testDB.Create(&position).Error
			require.NoError(t, err)
		})

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				actual, err := positionRepo.Get(tt.input)

				assert.Equal(t, tt.expected, actual)
				assert.Equal(t, tt.expectedErr, err)
			})
		}
	})

	t.Run("store", func(t *testing.T) {
		defer cleanupTestDB()

		tests := map[string]struct {
			input       model.Position
			expectedErr error
		}{
			"returns no error when success": {
				input: model.Position{
					OrderID:     "order-1",
					Symbol:      model.BTCUSD,
					Side:        model.Buy,
					Price:       1000,
					OrderStatus: model.Open,
					Quantity:    200,
				},
				expectedErr: nil,
			},
		}

		for name, tt := range tests {
			t.Run(name, func(t *testing.T) {
				err := positionRepo.Store(tt.input)
				assert.Equal(t, tt.expectedErr, err)
			})
		}
	})
}
