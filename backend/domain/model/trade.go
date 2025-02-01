package model

import (
	"errors"
	"math"
)

type Trade struct {
	Symbol               Symbol
	Side                 TradeSide
	HigherTimeFrameRates Rates
	LowerTimeFrameRates  Rates
	Quantity             float64
}

func NewTrade(symbol Symbol, higherTimeFrameRates, lowerTimeFrameRates Rates) (*Trade, error) {
	if len(higherTimeFrameRates) == 0 {
		return nil, errors.New("higher timeframe rates cannot be empty")
	}
	if len(lowerTimeFrameRates) == 0 {
		return nil, errors.New("lower timeframe rates cannot be empty")
	}

	return &Trade{
		Symbol:               symbol,
		HigherTimeFrameRates: higherTimeFrameRates,
		LowerTimeFrameRates:  lowerTimeFrameRates,
	}, nil
}

func (trade *Trade) Tradable() bool {
	if trade.isTrending() && trade.isPossibleEntry() {
		return true
	}

	return false
}

// 上位足の相場を見てトレンドが発生しているか判定する
func (trade *Trade) isTrending() bool {
	// TODO: implements
	trade.Side = Sell
	return false
}

// 下位足の相場を見てエントリー可能か判定する
func (trade *Trade) isPossibleEntry() bool {
	// TODO: implements
	return false
}

// 総資産から取引額の計算
func (trade *Trade) CalculateQuantity(totalWallet float64) error {
	if totalWallet < 100.0 {
		return errors.New("totalWallet must be greater than 100.0USD")
	}

	trade.Quantity = truncateToOneDecimal(totalWallet * 0.1)

	return nil
}

func (trade *Trade) MakePosition(price float64) Position {
	return Position{
		Symbol:      trade.Symbol,
		Side:        trade.Side,
		Price:       price,
		OrderStatus: Open,
		Quantity:    trade.Quantity,
	}
}

func truncateToOneDecimal(value float64) float64 {
	return math.Floor(value*10) / 10
}
