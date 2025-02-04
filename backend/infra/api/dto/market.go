package dto

import (
	"automatic-trade/backend/core/util"
	"automatic-trade/backend/domain/model"
)

type MarketResponse struct {
	Symbol string
	List   [][]string
}

func (m MarketResponse) SetModel(market *model.Market) error {
	rates := make(model.Rates, len(m.List))
	for i, l := range m.List {
		unixInt, err := util.StringToInt64(l[0])
		if err != nil {
			return err
		}

		closePrice, err := util.StringToFloat64(l[5])
		if err != nil {
			return err
		}

		rates[i] = model.Rate{
			DateTime: util.UnixToJST(unixInt, false),
			Price:    closePrice,
		}
	}

	market.Rates = rates
	return nil
}
