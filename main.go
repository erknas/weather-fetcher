package main

func main() {
	svc := NewLogger(&weatherFetcher{})
	jsonSrv := NewJSONAPIServer(":7000", svc)
	jsonSrv.Run()
}
