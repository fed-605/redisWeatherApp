package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fed-605/weatherApi/internal/cache"
	redisCache "github.com/fed-605/weatherApi/internal/cache/redis"
)

func main() {
	red, err := redisCache.NewRedisCache("localhost:6379", "")
	if err != nil {
		log.Fatal(err)
	}
	key := "hello"
	test := cache.WeatherResponse{
		ResolvedAddress: "Moscow222haha",
		Timezone:        "git log init",
	}
	if err := red.Set(context.Background(), key, &test, time.Minute); err != nil {
		log.Fatal(err)
	}
	ans, err := red.Get(context.Background(), key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*ans)

}

// package redisweather

// func main() {
// 	if err := env.Loadenv(); err != nil {
// 		log.Fatal(err)
// 	}
// 	cache, err := redis.NewRedisCache(env.GetEnvString("REDIS_ADDR", ""), env.GetEnvString("REDIS_PASS", ""))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	app := api.NewApplication(cache)
// 	srv := server.NewServer(env.GetEnvString("PORT", "8999"), app.Routes())
// 	srv.Run()
// }
