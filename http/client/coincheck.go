package client

import "encoding/json"

// Cticker Structure of ticker by api result
type Cticker struct {
	Last      int64
	Bid       int64
	Ask       int64
	High      int64
	Low       int64
	Volume    string
	Timestamp int64
}

// Ctrade Structure of trades by api result
type Ctrade struct {
	ID        int64 `json:"id"`
	Amount    string
	Rate      int64
	OrderType string `json:"order_type"`
	CreatedAt string `json:"created_at"`
}

// CoinCheckTicker get ticker from coincheck.jp
// https://coincheck.jp/documents/exchange/api?locale=ja#ticker
func CoinCheckTicker() Cticker {
	url := "https://coincheck.jp/api/ticker"
	byteArray := get(url)
	var t Cticker
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}

// CoinCheckTrades get trades from coincheck.jp
// https://coincheck.jp/documents/exchange/api?locale=ja#trades
func CoinCheckTrades() []Ctrade {
	url := "https://coincheck.jp/api/trades"
	byteArray := get(url)

	var t []Ctrade
	json.Unmarshal([]byte(string(byteArray)), &t)
	return t
}
