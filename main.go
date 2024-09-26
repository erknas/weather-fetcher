package main

func main() {
	service := NewLogger(&weatherFetcher{})

	server := NewJSONAPIServer(":7000", service)
	server.Run()
}
