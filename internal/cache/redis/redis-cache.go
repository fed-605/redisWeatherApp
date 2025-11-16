package redis

import (
	"context"
	"time"

	"github.com/fed-605/weatherApi/internal/cache"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	redis *redis.Client
}

func NewRedisCache(addr string, password string) (*redisCache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &redisCache{
		redis: client,
	}, nil

}

func (s *redisCache) Get(ctx context.Context, key string) (*cache.WeatherResponse, error) {
	return nil, nil
}

func (s *redisCache) Set(ctx context.Context, key string, value *cache.WeatherResponse, ttl time.Duration) error {
	return nil
}
