package repository

import "automatic-trade/backend/domain/model"

type Market interface {
	Get(limit int, market *model.Market) error
}
