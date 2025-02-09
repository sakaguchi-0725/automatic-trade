package dto

import (
	"automatic-trade/backend/core/util"
	"automatic-trade/backend/domain/model"
)

type RateResponse struct {
	List [][]string
}

func (rate RateResponse) ToRate() (model.Rate, error) {
	latest := rate.List[0]

	unixTime, err := util.StringToInt64(latest[0])
	if err != nil {
		return model.Rate{}, err
	}

	openPrice, err := util.StringToFloat64(latest[1])
	if err != nil {
		return model.Rate{}, err
	}

	return model.Rate{
		DateTime: util.UnixToJST(unixTime, false),
		Price:    openPrice,
	}, nil
}

func (rate RateResponse) ToRates() (model.Rates, error) {
	results := make(model.Rates, len(rate.List))

	for i, v := range rate.List {
		unixTime, err := util.StringToInt64(v[0])
		if err != nil {
			return model.Rates{}, err
		}

		openPrice, err := util.StringToFloat64(v[1])
		if err != nil {
			return model.Rates{}, err
		}

		results[i] = model.Rate{
			DateTime: util.UnixToJST(unixTime, false),
			Price:    openPrice,
		}
	}

	return results, nil
}
