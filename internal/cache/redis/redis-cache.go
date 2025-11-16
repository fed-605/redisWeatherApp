package redisCache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fed-605/weatherApi/internal/cache"
	"github.com/go-redis/redis/v8"
)

type redisCache struct {
	client *redis.Client
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
		client: client,
	}, nil

}

func (s *redisCache) Get(ctx context.Context, key string) (*cache.WeatherResponse, error) {

	resp, err := s.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}
	var wearherResp cache.WeatherResponse
	if err := json.Unmarshal([]byte(resp), &wearherResp); err != nil {
		return nil, err
	}
	return &wearherResp, nil

}

func (s *redisCache) Set(ctx context.Context, key string, value *cache.WeatherResponse, ttl time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = s.client.Set(ctx, key, string(v), ttl).Err()
	if err != nil {
		return err
	}
	return nil
}
