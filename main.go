package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/zeze322/weather-fetcher/client"
	"github.com/zeze322/weather-fetcher/proto"
)

func main() {
	var (
		svc  = NewLogger(&weatherFetcher{})
		ctx  = context.Background()
		city = flag.String("city", "", "city name")
		port = flag.String("port", ":6000", "gprc port")
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

		fmt.Printf("%v\n", weather)
	}()

	time.Sleep(2 * time.Second)
}
