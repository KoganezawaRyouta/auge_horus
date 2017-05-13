package api

import (
	"encoding/json"

	"github.com/KoganezawaRyouta/augehorus/model"
	"github.com/KoganezawaRyouta/augehorus/serializer"
	"github.com/valyala/fasthttp"
)

func (app *Api) Trades(ctx *fasthttp.RequestCtx) {
	trades := []model.Trade{}
	app.DbAdapter.DB.Find(&trades)

	if len(trades) > 0 {
		tradesJSON := serializer.TradesParse(trades)
		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("application/json")
		enc := json.NewEncoder(ctx)
		enc.Encode(tradesJSON)
	} else {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetContentType("application/json")
		ctx.SetBodyString("{\"message\":\"data not found\"}")
	}
}
