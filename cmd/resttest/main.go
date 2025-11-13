package main

import (
	"context"
	"fmt"
	"os"
	"resttest/otel"
	"resttest/runner"
	"resttest/webserver"
	"syscall"
)

func main() {
	fmt.Println("Hello from resttest")

	var (
		ctx      context.Context = context.Background()
		runGroup runner.Group    = runner.Group{}
	)

	// Set up OpenTelemetry.
	otelShutdown, errOtel := otel.SetupOTelSDK(ctx)
	if errOtel != nil {
		fmt.Println("Error while instantiating OTel. Err:", errOtel)
		os.Exit(1)
	}
	defer otelShutdown(context.Background())

	// Handle SIGINT (CTRL+C) gracefully.
	runGroup.Add(runner.SignalHandler(ctx, os.Interrupt, syscall.SIGINT, syscall.SIGTERM))

	// Handle SIGINT (CTRL+C) gracefully.
	server := webserver.NewServer(ctx)
	runGroup.Add(func() error {
		return server.ListenAndServe()
	}, func(err error) {
		server.Shutdown(context.Background())
	})

	var cause error = runGroup.Run()
	fmt.Println(cause)
}
