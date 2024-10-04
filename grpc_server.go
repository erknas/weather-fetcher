package main

import (
	"context"
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/zeze322/weather-fetcher/proto"
	"google.golang.org/grpc"
)

type GRPCWeatherFetcherServer struct {
	port string
	svc  WeatherFetcher
	proto.UnimplementedWeatherFetcherServer
}

func NewGRPCWeatherFetcherServer(port string, svc WeatherFetcher) *GRPCWeatherFetcherServer {
	return &GRPCWeatherFetcherServer{
		port: port,
		svc:  svc,
	}
}

func (s *GRPCWeatherFetcherServer) Run() error {
	ln, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	fmt.Printf("listening port %s\n", s.port)

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)
	srv := NewGRPCWeatherFetcherServer(s.port, s.svc)
	proto.RegisterWeatherFetcherServer(server, srv)

	return server.Serve(ln)
}

func (s *GRPCWeatherFetcherServer) FetchWeather(ctx context.Context, req *proto.CityRequest) (*proto.WeatherResponse, error) {
	ctx = context.WithValue(ctx, RequestIDKey{}, uuid.New().String())

	weather, err := s.svc.FetchWeather(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	protoLocation := &proto.Location{
		Name:      weather.Location.Name,
		Region:    weather.Location.Region,
		Country:   weather.Location.Country,
		Latitude:  weather.Location.Latitude,
		Longitude: weather.Location.Longitude,
		Localtime: weather.Location.LocalTime,
	}

	protoCurrent := &proto.Current{
		Temperature: weather.Current.Temperature,
		Wind:        weather.Current.Wind,
		Pressure:    weather.Current.Pressure,
		Humidity:    weather.Current.Humidity,
		Feelslike:   weather.Current.Feelslike,
		Condition: &proto.Condition{
			Name: weather.Current.Condition.Name,
		},
	}

	resp := &proto.WeatherResponse{
		Location: protoLocation,
		Current:  protoCurrent,
	}

	return resp, nil
}
