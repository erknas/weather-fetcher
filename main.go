package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/zeze322/weather-fetcher/client"
	"github.com/zeze322/weather-fetcher/proto"
)

func main() {
	var (
		ctx        = context.Background()
		svc        = NewLogger(&weatherFetcher{})
		redisCl    = redis.NewClient(&redis.Options{Addr: "localhost:6379"})
		redisCache = NewRedisClient(redisCl)
		city       = flag.String("city", "", "city name")
		port       = flag.String("port", ":6000", "gprc port")
	)
	flag.Parse()

	grpcSrv := NewGRPCWeatherFetcherServer(*port, svc)
	go grpcSrv.Run()

	grpcCl, err := client.GRPCClient(*port)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		weather, err := grpcCl.FetchWeather(ctx, &proto.CityRequest{Name: *city})
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Current weather: %v\n", weather)
	}()

	// time.Sleep(time.Second * 2)
	jsonSrv := NewJSONAPIServer(":3000", svc, redisCache)
	jsonSrv.Run()
}
