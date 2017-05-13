package app

import (
	"net/http"
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/labstack/echo"
)

type App struct {
	Config    *settings.Config
	echofw    *echo.Echo
}

func AppNew(config *settings.Config) *App {
	app := &App{Config: config}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &MonitorContext{c}
			return h(cc)
		}
	})

	e.GET("/monitor", app.Monitor)
	app.echofw = e

	return app
}

func (app *App) Listen() error {
	return app.echofw.Start(app.Config.AppServer.Port)
}
