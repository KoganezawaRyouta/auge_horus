package model

import (
	"time"

	httpClient "github.com/KoganezawaRyouta/augehorus/http/client"
)

type Trade struct {
	ID        int64 `gorm:"primary_key"`
	TradeID   int64
	Amount    string
	Rate      int64
	OrderType string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTradeForCtrade(chTrade httpClient.Ctrade) Trade {
	return Trade{
		TradeID:   chTrade.ID,
		Amount:    chTrade.Amount,
		Rate:      chTrade.Rate,
		OrderType: chTrade.OrderType,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
