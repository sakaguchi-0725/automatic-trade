package repository

import "automatic-trade/backend/domain/model"

type Order interface {
	Place(order *model.Order) error
	Cancel(symbol model.Symbol, orderID string) error
}
