package main

import (
	"devstream-rest-api/internal/config"
	"devstream-rest-api/internal/database"
	"devstream-rest-api/internal/handler"
	"devstream-rest-api/internal/repository"
	"devstream-rest-api/internal/router"
	"devstream-rest-api/internal/util/appauth"
	"devstream-rest-api/internal/util/applogger"
	"devstream-rest-api/internal/util/idgen"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config := config.Load()

	appauth.Init(config.AUTH_SECRET)

	database := database.NewPgDatabase(&config)
	database.Connect()

	r := chi.NewRouter()

	r.Use(applogger.AppLoggerMiddleware)

	idGen := idgen.NewSnowflakeIdGen()

	userRepo := repository.NewUserRepositoryPg(database.DB, idGen)
	streamKeyRepo := repository.NewStreamKeyRepositoryPG(database.DB, idGen)

	authHandler := handler.NewAuthHandler(&userRepo)
	streamKeysHandler := handler.NewStreamKeyHandler(&streamKeyRepo)

	router.AuthRoutes(r, &authHandler)
	router.StreamKeyRoutes(r, &streamKeysHandler)

	http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), r)
}
