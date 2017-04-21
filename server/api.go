package server

import (
	"github.com/KoganezawaRyouta/augehorus/orm"
	"github.com/KoganezawaRyouta/uppercut"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type Api struct {
	DbAdapter *orm.GormAdapter
	Uppercut  *uppercut.Uppercut
}

func ApiNew(configName string) *Api {
	app := &Api{
		DbAdapter: orm.NewGormAdapter(configName),
	}

	// set Router
	router := fasthttprouter.New()
	router.GET("/tickers", app.Tickers)
	router.GET("/trades", app.Trades)
	router.GET("/monitor", app.Monitor)

	// loggerMiddleware := NewLoudLoggerMiddleware()
	// counters := []uppercut.Counter{loggerMiddleware}
	app.Uppercut = uppercut.NewUppercut(router.Handler)
	app.Uppercut.AddCounters(NewLoudLoggerMiddleware())
	app.Uppercut.AddBeforeCounters(PanicWrapMiddleware)
	return app
}

func (app *Api) Listen() error {
	return fasthttp.ListenAndServe(":8080", app.Uppercut.Handler)
}
