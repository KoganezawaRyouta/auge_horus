package middleware

import (
	"encoding/json"

	"github.com/KoganezawaRyouta/uppercut"
	"github.com/valyala/fasthttp"
)

var CrashMiddleware = uppercut.CounterFunc(func(ctx *fasthttp.RequestCtx) {
	defer func() {
		err := recover()

		if err != nil {
			message := ""
			statusCode := fasthttp.StatusServiceUnavailable
			attrs := map[string]interface{}{}

			switch e := err.(type) {
			case string:
				message = e
			case error:
				message = e.Error()
			}
			attrs["message"] = message

			ctx.SetStatusCode(statusCode)
			ctx.SetContentType("application/json")
			w := json.NewEncoder(ctx)
			w.Encode(attrs)
		}
	}()
})
