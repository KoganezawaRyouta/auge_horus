package api

import (
	"encoding/json"

	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/serializer"
	"github.com/valyala/fasthttp"
)

// tickers handler
func (app *Api) Tickers(ctx *fasthttp.RequestCtx) {
	tickers := []model.Ticker{}
	app.DbAdapter.DB.Find(&tickers)

	if len(tickers) > 0 {
		tickersJSON := serializer.TickersParse(tickers)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("application/json")
		enc := json.NewEncoder(ctx)
		enc.Encode(tickersJSON)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.SetBodyString("{\"message\":\"data not found\"}")
	}
}
