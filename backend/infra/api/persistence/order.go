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

type orderRepository struct {
	client *bybit.Client
}

func (o *orderRepository) Cancel(symbol model.Symbol, orderID string) error {
	params := map[string]interface{}{
		"category": "inverse",
		"symbol":   symbol.String(),
		"orderId":  orderID,
	}

	_, err := o.client.NewUtaBybitServiceWithParams(params).CancelOrder(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) Place(order *model.Order) error {
	params := map[string]interface{}{
		"category":  "inverse",
		"symbol":    order.Symbol.String(),
		"side":      order.Side.String(),
		"orderType": order.OrderType.String(),
		"qty":       order.Quantity,
	}

	res, err := o.client.NewUtaBybitServiceWithParams(params).PlaceOrder(context.Background())
	if err != nil {
		return err
	}

	result, err := util.DecodeJSON[dto.ServerResponse[dto.PlaceOrderResponse]](res)
	if err != nil {
		return err
	}

	order.ID = result.Result.OrderID
	return nil
}

func NewOrderRepository(cfg config.Bybit) repository.Order {
	return &orderRepository{
		client: bybit.NewBybitHttpClient(cfg.APIKey, cfg.SecretKey, bybit.WithBaseURL(cfg.BaseURL)),
	}
}
