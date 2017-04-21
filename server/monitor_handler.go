package server

import (
	"fmt"

	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/valyala/fasthttp"
)

func (app *Api) Monitor(ctx *fasthttp.RequestCtx) {
	err := app.DbAdapter.DB.DB().Ping()
	fmt.Fprintf(ctx, "api server version v%s\n", settings.Version)
	fmt.Fprintf(ctx, "%s\n", settings.GoVersion)
	fmt.Fprintf(ctx, "BuildDhash%s\n", settings.BuildDhash)
	if err != nil {
		fmt.Fprintf(ctx, "db ping error")
	} else {
		fmt.Fprintf(ctx, "db ping ok")
	}
}
