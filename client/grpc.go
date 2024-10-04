package client

import (
	"github.com/zeze322/weather-fetcher/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GRPCClient(port string) (proto.WeatherFetcherClient, error) {
	conn, err := grpc.NewClient(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := proto.NewWeatherFetcherClient(conn)

	return client, nil
}
