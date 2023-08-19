package applogger

import (
	"devstream-rest-api/internal/util/appenv"
	"net/http"
)

type Logger interface {
	I(any, ...any)
	E(any, ...any)
	D(any, ...any)
}

var env = appenv.AppEnv()

func New(name string) Logger {
	switch env {
	case "development":
		return newConsoleLogger(name)
	case "production":
		return newSentryLogger(name)
	default:
		return newConsoleLogger(name)
	}
}

func AppLoggerMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if env != "development" {
			h.ServeHTTP(w, r)
		} else {
			logger := New("REQUEST")
			logger.I(r.RequestURI)
			h.ServeHTTP(w, r)
		}
	})
}
