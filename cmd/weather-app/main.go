package main

import (
	"log"
	"time"

	"github.com/fed-605/weatherApi/env"
	"github.com/fed-605/weatherApi/internal/app"
	redisCache "github.com/fed-605/weatherApi/internal/cache/redis"
	"github.com/fed-605/weatherApi/internal/provider"
	"github.com/fed-605/weatherApi/internal/transport/api"
	"github.com/fed-605/weatherApi/server"
)

const (
	cacheTTL = time.Minute
)

func main() {
	//load env file
	if err := env.Loadenv(); err != nil {
		log.Fatal(err)
	}

	//init redis
	red, err := redisCache.NewRedisCache(env.GetEnvString("REDIS_ADDR", ""), "")
	if err != nil {
		log.Fatal(err)
	}

	//init provider
	prvdr := provider.NewvisualCrossingProvider(env.GetEnvString("PROVIDER_BASE_URL", ""), env.GetEnvString("API_KEY", ""))
	wService := app.NewService(red, prvdr, cacheTTL)

	//init app
	applicationAPI := api.NewApplication(wService)

	//init server
	srv := server.NewServer(env.GetEnvString("PORT", ""), applicationAPI.Routes())

	//start server
	srv.Run()
}
