package main

import (
	"context"

	"github.com/zeze322/weather-fetcher/types"
)

type Cacher interface {
	Get(context.Context, string) (types.WeatherResponse, bool)
	Set(context.Context, string, types.WeatherResponse) error
}

type Cache struct{}

func (c *Cache) Get(context.Context, string) (types.WeatherResponse, bool) {
	return types.WeatherResponse{}, false
}

func (c *Cache) Set(context.Context, string, types.WeatherResponse) error {
	return nil
}
