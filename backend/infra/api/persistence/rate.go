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

type rateRepository struct {
	client *bybit.Client
}

func (repo *rateRepository) Fetch(ctx context.Context, limit int, symbol model.Symbol, interval model.Interval) (model.Rates, error) {
	params := map[string]interface{}{
		"category": "inverse",
		"symbol":   symbol.String(),
		"interval": interval.String(),
		"limit":    limit,
	}

	res, err := repo.client.NewUtaBybitServiceWithParams(params).GetMarketKline(ctx)
	if err != nil {
		return model.Rates{}, err
	}

	data, err := util.DecodeJSON[dto.ServerResponse[dto.RateResponse]](res)
	if err != nil {
		return model.Rates{}, err
	}

	return data.Result.ToRates()
}

func (repo *rateRepository) Get(ctx context.Context, symbol model.Symbol, interval model.Interval) (model.Rate, error) {
	params := map[string]interface{}{
		"category": "inverse",
		"symbol":   symbol.String(),
		"interval": interval.String(),
		"limit":    1,
	}

	res, err := repo.client.NewUtaBybitServiceWithParams(params).GetMarketKline(ctx)
	if err != nil {
		return model.Rate{}, err
	}

	data, err := util.DecodeJSON[dto.ServerResponse[dto.RateResponse]](res)
	if err != nil {
		return model.Rate{}, err
	}

	return data.Result.ToRate()
}

func NewRateRepository(cfg config.Bybit) repository.Rate {
	return &rateRepository{
		client: bybit.NewBybitHttpClient(cfg.APIKey, cfg.SecretKey, bybit.WithBaseURL(cfg.BaseURL)),
	}
}
