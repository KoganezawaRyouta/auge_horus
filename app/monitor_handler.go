package app

import (
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/labstack/echo"
	"net/http"
)

type MonitorContext struct {
	echo.Context
}

func (c *MonitorContext) Push() {
	println("server version: v%s\n", settings.Version)
	println("%s\n", settings.GoVersion)
	println("BuildDhash: %s\n", settings.BuildDhash)
	println("=================")
	println("ProcessName: %s\n", settings.ProcessName)
	println("PID: %d\n", settings.PID)
	println("ParentProcessName: %s\n", settings.ParentProcessName)
	println("PPID: %d\n", settings.PPID)
}

func (app *App) Monitor(ctx echo.Context) error {
	cc := ctx.(*MonitorContext)
	cc.Push()
	return cc.String(http.StatusOK, "StatusOK!")
}
