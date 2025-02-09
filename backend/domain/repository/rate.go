package repository

import (
	"automatic-trade/backend/domain/model"
	"context"
)

type Rate interface {
	Get(ctx context.Context, symbol model.Symbol, interval model.Interval) (model.Rate, error)
	Fetch(ctx context.Context, limit int, symbol model.Symbol, interval model.Interval) (model.Rates, error)
}
