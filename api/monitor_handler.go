package api

import (
	"fmt"

	"github.com/KoganezawaRyouta/augehorus/config"
	"github.com/valyala/fasthttp"
)

// monitor handler
func (app *Api) Monitor(ctx *fasthttp.RequestCtx) {
	err := app.DbAdapter.DB.DB().Ping()
	if err != nil {
		fmt.Fprintf(ctx, "db ping: error!!")
	} else {
		fmt.Fprintf(ctx, "db ping: ok!!\n")
	}

	fmt.Fprintf(ctx, "server version: v%s\n", config.Version)
	fmt.Fprintf(ctx, "%s\n", config.GoVersion)
	fmt.Fprintf(ctx, "BuildDhash: %s\n", config.BuildDhash)
	fmt.Printf("=================")
	fmt.Fprintf(ctx, "ProcessName: %s\n", config.ProcessName)
	fmt.Fprintf(ctx, "PID: %d\n", config.PID)
	fmt.Fprintf(ctx, "ParentProcessName: %s\n", config.ParentProcessName)
	fmt.Fprintf(ctx, "PPID: %d\n", config.PPID)
}
