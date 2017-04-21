package server

import (
	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type App struct {
	DbAdapter *orm.GormAdapter
	Router    *fasthttprouter.Router
}

func AppNew(configName string) *App {
	app := &App{
		DbAdapter: orm.NewGormAdapter(configName),
		Router:    fasthttprouter.New(),
	}

	// set Router
	{
		// app.Router.GET("/tickers", app.Tickers)
		// app.Router.GET("/trades", app.Trades)
	}

	return app
}

func (app *App) Listen() error {
	return fasthttp.ListenAndServe(":9090", app.Router.Handler)
}
