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
	"devstream-rest-api/pkg/userpb"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	config := config.Load()

	appauth.Init(config.AUTH_SECRET)

	database := database.NewPgDatabase(&config)
	database.Connect()

	userService := userGrpcService()

	r := chi.NewRouter()

	r.Use(applogger.AppLoggerMiddleware)

	idGen := idgen.NewSnowflakeIdGen()

	userRepo := repository.NewUserRepositoryGrpc(database.DB, userService, idGen)
	streamKeyRepo := repository.NewStreamKeyRepositoryPG(database.DB, idGen)

	authHandler := handler.NewAuthHandler(&userRepo)
	streamKeysHandler := handler.NewStreamKeyHandler(&streamKeyRepo)

	router.AuthRoutes(r, &authHandler)
	router.StreamKeyRoutes(r, &streamKeysHandler)

	http.ListenAndServe(fmt.Sprintf(":%v", config.PORT), r)
}

func userGrpcService() userpb.UserServiceRpcClient {
	conn, err := grpc.Dial("data-service:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic("Unable to connect to data service")
	}

	userClient := userpb.NewUserServiceRpcClient(conn)

	return userClient
}
