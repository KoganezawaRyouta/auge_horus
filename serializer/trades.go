package serializer

import (
	"time"

	"github.com/KoganezawaRyouta/augehorus/model"
)

type TradesJSON struct {
	Trades []*TradeJSON `json:"trades"`
}

type TradeJSON struct {
	ID        int64     `json:"id"`
	TradeID   int64     `json:"trade_id"`
	Amount    string    `json:"amount"`
	Rate      int64     `json:"rate"`
	OrderType string    `json:"order_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewTradeJSON(t model.Trade) *TradeJSON {
	tradeJSON := &TradeJSON{
		ID:        t.ID,
		TradeID:   t.TradeID,
		Amount:    t.Amount,
		Rate:      t.Rate,
		OrderType: t.OrderType,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}

	return tradeJSON
}

func NewTradesJSON(trades []model.Trade) []*TradeJSON {
	jsons := make([]*TradeJSON, len(trades))
	for i, m := range trades {
		jsons[i] = NewTradeJSON(m)
	}

	return jsons
}

func TradesParse(trades []model.Trade) TradesJSON {
	return TradesJSON{Trades: NewTradesJSON(trades)}
}
