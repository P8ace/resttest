package webserver

import (
	"context"
	"net"
	"net/http"
	controller "resttest/webcontrollers"
	middleware "resttest/webmiddlewares"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewServer(ctx context.Context) *http.Server {
	server := http.Server{
		Addr:         ":5432",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      registerControllers(),
	}

	return &server
}

// func registerControllers() *http.ServeMux {
func registerControllers() http.Handler {
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

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}
	handleFunc("GET /rolldice", controller.Rolldice)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
	// mux.HandleFunc("POST /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("PUT /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("PATCH /api/v1/item", controllers.HandleHealthCheck)
	// mux.HandleFunc("DELETE /api/v1/items/", controllers.HandleHealthCheck)
	// mux.HandleFunc("HEAD /api/v1/items", controllers.HandleHealthCheck)
	// mux.HandleFunc("OPTIONS /api/v1/healthcheck", controllers.HandleHealthCheck)
	// mux.HandleFunc("TRACE /api/v1/healthcheck", controllers.HandleHealthCheck)
	// mux.HandleFunc("CONNECT /api/v1/healthcheck", controllers.HandleHealthCheck)

	// return mux
}
