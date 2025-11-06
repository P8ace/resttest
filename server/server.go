package server

import (
	"net/http"
	controllers "resttest/server/controllers"
)

func NewServer() *http.Server {
	server := http.Server{
		Addr:    ":5432",
		Handler: registerControllers(),
	}

	return &server
}

func registerControllers() *http.ServeMux {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("GET /healthcheck", controllers.HandleHealthCheck)

	mux.HandleFunc("GET /api/v1/items", controllers.HandleGetItems)
	// mux.HandleFunc("POST /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("PUT /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("PATCH /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("DELETE /api/v1/items/", controllers.HandleHealthCheck)
	// mux.HandleFunc("HEAD /api/v1/items", controllers.HandleHealthCheck)
	// mux.HandleFunc("OPTIONS /api/v1/healthcheck", controllers.HandleHealthCheck)
	// mux.HandleFunc("TRACE /api/v1/healthcheck", controllers.HandleHealthCheck)
	// mux.HandleFunc("CONNECT /api/v1/healthcheck", controllers.HandleHealthCheck)

	return mux
}
