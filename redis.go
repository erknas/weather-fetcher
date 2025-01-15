package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeze322/weather-fetcher/types"
)

const ttl = time.Minute * 90

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(client *redis.Client) *RedisClient {
	return &RedisClient{
		client: client,
	}
}

func (c *RedisClient) Get(ctx context.Context, key string) (types.WeatherResponse, bool) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return types.WeatherResponse{}, false
	}

	var weather types.WeatherResponse

	if err := json.Unmarshal([]byte(val), &weather); err != nil {
		return types.WeatherResponse{}, false
	}

	return weather, true
}

func (c *RedisClient) Set(ctx context.Context, key string, val types.WeatherResponse) error {
	weather, err := json.Marshal(val)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, weather, ttl).Err()
}
