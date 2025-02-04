package persistence

import (
	"automatic-trade/backend/core/config"
	"automatic-trade/backend/core/util"
	"automatic-trade/backend/domain/model"
	"automatic-trade/backend/domain/repository"
	"automatic-trade/backend/infra/api/dto"
	"context"

	bybit "github.com/wuhewuhe/bybit.go.api"
)

type marketRepository struct {
	client *bybit.Client
}

func (m *marketRepository) Get(limit int, market *model.Market) error {
	params := map[string]interface{}{
		"category": "inverse",
		"symbol":   market.Symbol.String(),
		"interval": market.Interval,
		"limit":    limit,
	}

	res, err := m.client.NewUtaBybitServiceWithParams(params).GetMarketKline(context.Background())
	if err != nil {
		return err
	}

	result, err := util.DecodeJSON[dto.ServerResponse[dto.MarketResponse]](res)
	if err != nil {
		return err
	}

	if err := result.Result.SetModel(market); err != nil {
		return err
	}

	return nil
}

func NewMarketRepository(cfg config.Bybit) repository.Market {
	return &marketRepository{
		client: bybit.NewBybitHttpClient(cfg.APIKey, cfg.SecretKey, bybit.WithBaseURL(cfg.BaseURL)),
	}
}
