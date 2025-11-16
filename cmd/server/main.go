package main

import (
	"log"

	"github.com/fed-605/weatherApi/env"
	"github.com/fed-605/weatherApi/internal/api"
	"github.com/fed-605/weatherApi/internal/cache/redis"
	"github.com/fed-605/weatherApi/server"
)

func main() {
	if err := env.Loadenv(); err != nil {
		log.Fatal(err)
	}
	cache, err := redis.NewRedisCache(env.GetEnvString("REDIS_ADDR", ""), env.GetEnvString("REDIS_PASS", ""))
	if err != nil {
		log.Fatal(err)
	}
	app := api.NewApplication(cache)
	srv := server.NewServer(env.GetEnvString("PORT", "8999"), app.Routes())
	srv.Run()
}
