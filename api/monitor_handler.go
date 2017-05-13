package api

import (
	"fmt"

	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/valyala/fasthttp"
)

func (app *Api) Monitor(ctx *fasthttp.RequestCtx) {
	err := app.DbAdapter.DB.DB().Ping()
	if err != nil {
		fmt.Fprintf(ctx, "db ping: error!!")
	} else {
		fmt.Fprintf(ctx, "db ping: ok!!\n")
	}

	fmt.Fprintf(ctx, "server version: v%s\n", settings.Version)
	fmt.Fprintf(ctx, "%s\n", settings.GoVersion)
	fmt.Fprintf(ctx, "BuildDhash: %s\n", settings.BuildDhash)
	fmt.Printf("=================")
	fmt.Fprintf(ctx, "ProcessName: %s\n", settings.ProcessName)
	fmt.Fprintf(ctx, "PID: %d\n", settings.PID)
	fmt.Fprintf(ctx, "ParentProcessName: %s\n", settings.ParentProcessName)
	fmt.Fprintf(ctx, "PPID: %d\n", settings.PPID)
}
