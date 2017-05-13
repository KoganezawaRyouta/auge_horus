package middleware

import (
	"fmt"
	"os"

	"github.com/KoganezawaRyouta/augehorus/settings"
	"github.com/go-kit/kit/log"

	"github.com/valyala/fasthttp"
)

var logger log.Logger

type LoudLoggerMiddleware struct {
	logger log.Logger
}

// NewLoudLoggerMiddleware ttt
func NewLoudLoggerMiddleware(config *settings.Config) *LoudLoggerMiddleware {
	logfile, err := os.OpenFile(config.ApiServer.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot open " + config.ApiServer.LogFile + err.Error())
	}

	logger = log.NewJSONLogger(log.NewSyncWriter(logfile))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "api server", log.DefaultCaller)
	return &LoudLoggerMiddleware{
		logger: logger,
	}
}

func (l *LoudLoggerMiddleware) Call(ctx *fasthttp.RequestCtx) {
	l.logger.Log("request: ", fmt.Sprintf("Host: %s, Path: %s, Method: %s", ctx.Host(), ctx.Path(), ctx.Method()))
}
