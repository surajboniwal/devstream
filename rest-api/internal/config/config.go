package config

import (
	"devstream-rest-api/internal/util/appenv"
	"devstream-rest-api/internal/util/applogger"
	"flag"
	"fmt"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	PORT        string `koanf:"PORT"`
	DB_URL      string `koanf:"DB_URL"`
	APP_ENV     string
	AUTH_SECRET string `koanf:"AUTH_SECRET"`
}

var env = appenv.AppEnv()

var k = koanf.New(".")

var logger = applogger.New("config")

func Load() Config {
	flag.Parse()
	var config Config = Config{
		APP_ENV: env,
	}

	if err := k.Load(file.Provider(fmt.Sprintf("./internal/config/%v.env", env)), dotenv.Parser()); err != nil {
		logger.E(err)
	}

	if err := k.Unmarshal("", &config); err != nil {
		logger.E(err)
	}

	return config
}
