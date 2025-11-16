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
