package repository

import "automatic-trade/backend/domain/model"

type Position interface {
	Store(position model.Position) error
	Get(orderID string) (model.Position, error)
	Delete(orderID string) error
}
