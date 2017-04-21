package serializer

import (
	"time"

	"github.com/KoganezawaRyouta/augehorus/model"
)

type TickersJSON struct {
	Tickers []*TickerJSON `json:"tickers"`
}

type TickerJSON struct {
	ID        int64     `json:"id"`
	Last      int64     `json:"last"`
	Bid       int64     `json:"bid"`
	Ask       int64     `json:"ask"`
	High      int64     `json:"high"`
	Low       int64     `json:"low"`
	Volume    string    `json:"volume"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewTickerJSON(t model.Ticker) *TickerJSON {
	tickerJSON := &TickerJSON{
		ID:        t.ID,
		Last:      t.Last,
		Bid:       t.Bid,
		Ask:       t.Ask,
		High:      t.High,
		Low:       t.Low,
		Volume:    t.Volume,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}

	return tickerJSON
}

func NewTickersJSON(tickers []model.Ticker) []*TickerJSON {
	jsons := make([]*TickerJSON, len(tickers))
	for i, m := range tickers {
		jsons[i] = NewTickerJSON(m)
	}

	return jsons
}

func TickersParse(tickers []model.Ticker) TickersJSON {
	return TickersJSON{Tickers: NewTickersJSON(tickers)}
}
