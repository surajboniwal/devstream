package database

import (
	"devstream-rest-api/internal/config"
	"devstream-rest-api/internal/util/applogger"
)

type Database interface {
	Connect(config.Config)
}

var logger applogger.Logger = applogger.New("database")
