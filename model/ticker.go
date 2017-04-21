package model

import (
	"time"

	httpClient "github.com/KoganezawaRyouta/augehorus/http/client"
)

type Ticker struct {
	ID        int64 `gorm:"primary_key"`
	Last      int64
	Bid       int64
	Ask       int64
	High      int64
	Low       int64
	Volume    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTickerForCticker(chTicker httpClient.Cticker) Ticker {
	return Ticker{
		Last:      chTicker.Last,
		Bid:       chTicker.Bid,
		Ask:       chTicker.Ask,
		High:      chTicker.High,
		Low:       chTicker.Low,
		Volume:    chTicker.Volume,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
