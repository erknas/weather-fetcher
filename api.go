package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type JSONAPIServer struct {
	port string
	svc  WeatherFetcher
}

func NewJSONAPIServer(port string, svc WeatherFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		port: port,
		svc:  svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandler(s.handleFetchWeather))

	log.Fatal(http.ListenAndServe(s.port, nil))
}

func (s *JSONAPIServer) handleFetchWeather(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	city := r.URL.Query().Get("city")

	resp, err := s.svc.FetchWeather(ctx, city)
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, resp)
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func makeHTTPHandler(fn APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, requestIDKey{}, uuid.New().String())
		if err := fn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}
