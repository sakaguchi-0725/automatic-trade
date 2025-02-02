package model

import (
	"errors"
	"math"
)

type Trade struct {
	Side            TradeSide
	HigherTimeFrame Market
	LowerTimeFrame  Market
	Quantity        float64
}

func NewTrade(higherTimeFrame, lowerTimeFrame Market) (*Trade, error) {
	if higherTimeFrame.Symbol != lowerTimeFrame.Symbol {
		return nil, errors.New("higher and lower time frame symbols must match")
	}
	if len(higherTimeFrame.Rates) == 0 {
		return nil, errors.New("higher timeframe rates cannot be empty")
	}
	if len(lowerTimeFrame.Rates) == 0 {
		return nil, errors.New("lower timeframe rates cannot be empty")
	}

	return &Trade{
		HigherTimeFrame: higherTimeFrame,
		LowerTimeFrame:  lowerTimeFrame,
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

func (trade *Trade) MakeOrder() Order {
	return Order{
		OrderType: MarketOrder,
		Side:      trade.Side,
		Symbol:    trade.HigherTimeFrame.Symbol,
		Quantity:  trade.Quantity,
	}
}

func truncateToOneDecimal(value float64) float64 {
	return math.Floor(value*10) / 10
}
