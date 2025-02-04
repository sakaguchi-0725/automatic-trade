package dto

import "automatic-trade/backend/domain/model"

type Position struct {
	OrderID     string `gorm:"primaryKey"`
	Symbol      string
	Side        string `gorm:"type: enum('buy', 'sell'); not null"`
	Price       float64
	OrderStatus string  `gorm:"type: enum('open', 'close'); default: 'open'; not null"`
	Quantity    float64 `gorm:"not null"`
}

func (p Position) ToModel() (model.Position, error) {
	symbol, err := model.NewSymbolFromString(p.Symbol)
	if err != nil {
		return model.Position{}, err
	}

	return model.Position{
		OrderID:     p.OrderID,
		Symbol:      symbol,
		Side:        model.TradeSide(p.Side),
		Price:       p.Price,
		OrderStatus: model.OrderStatus(p.OrderStatus),
		Quantity:    p.Quantity,
	}, nil
}

func NewPosition(model model.Position) Position {
	return Position{
		OrderID:     model.OrderID,
		Symbol:      model.Symbol.String(),
		Side:        model.Side.String(),
		Price:       model.Price,
		OrderStatus: model.OrderStatus.String(),
		Quantity:    model.Quantity,
	}
}
