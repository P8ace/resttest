package webserver

import (
	"net/http"
	controller "resttest/webcontrollers"
	middleware "resttest/webmiddlewares"
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

	// Register controllers and the desired middlewares on the router (mux)

	mux.HandleFunc("GET /healthcheck",
		middleware.ChainMiddleWares(
			controller.HandleHealthCheck,
			middleware.LoggingMiddleWare(),
			middleware.MethodMiddleWare()))

	mux.HandleFunc("GET /api/v1/items",
		middleware.ChainMiddleWares(
			controller.HandleGetItems,
			middleware.LoggingMiddleWare(),
			middleware.MethodMiddleWare()))

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
