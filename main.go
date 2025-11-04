package main

import (
	"fmt"
	"net/http"
)

func handleHealthCheck(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Serivce is healthy")
}

func restServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", handleHealthCheck)

	server := http.Server{
		Addr:    ":5432",
		Handler: mux,
	}

	return &server
}

func main() {
	fmt.Println("Hello from resttest")

	error := restServer().ListenAndServe()
	fmt.Printf("error: %v\n", error)
}
