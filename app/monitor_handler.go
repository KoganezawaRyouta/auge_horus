package app

import (
	"github.com/KoganezawaRyouta/augehorus/config"
	"github.com/labstack/echo"
	"net/http"
)

type MonitorContext struct {
	echo.Context
}

func (c *MonitorContext) Push() {
	println("server version: v%s\n", config.Version)
	println("%s\n", config.GoVersion)
	println("BuildDhash: %s\n", config.BuildDhash)
	println("=================")
	println("ProcessName: %s\n", config.ProcessName)
	println("PID: %d\n", config.PID)
	println("ParentProcessName: %s\n", config.ParentProcessName)
	println("PPID: %d\n", config.PPID)
}

// monito handler
func (app *App) Monitor(ctx echo.Context) error {
	cc := ctx.(*MonitorContext)
	cc.Push()
	return cc.String(http.StatusOK, "StatusOK!")
}
