package api

import (
	"github.com/KoganezawaRyouta/augehorus/middleware"
	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/KoganezawaRyouta/uppercut"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Api struct {
	DbAdapter *orm.GormAdapter
	Uppercut  *uppercut.Uppercut
	Config    *settings.Config
}

func ApiNew(config *settings.Config) *Api {
	app := &Api{
		DbAdapter: orm.NewGormAdapter(config),
		Config:    config,
	}

	// set Router
	router := fasthttprouter.New()
	router.GET("/tickers", app.Tickers)
	router.GET("/trades", app.Trades)
	router.GET("/monitor", app.Monitor)

	app.Uppercut = uppercut.NewUppercut(router.Handler)
	app.Uppercut.AddCounters(middleware.NewLoudLoggerMiddleware(config))
	app.Uppercut.AddHooks(middleware.CrashMiddleware)
	return app
}

func (app *Api) Listen() error {
	return fasthttp.ListenAndServe(app.Config.ApiServer.Port, app.Uppercut.Handler)
}
