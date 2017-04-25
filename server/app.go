package server

import (
	"net/http"

	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/labstack/echo"
)

type App struct {
	DbAdapter *orm.GormAdapter
	Config    *settings.Config
	echofw    *echo.Echo
}

func AppNew(config *settings.Config) *App {
	app := &App{
		DbAdapter: orm.NewGormAdapter(config),
		Config:    config,
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.echofw = e

	return app
}

func (app *App) Listen() error {
	return app.echofw.Start(app.Config.AppServer.Port)
}
