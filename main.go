package main

import (
	"fmt"
	"resttest/server"
)

func main() {
	fmt.Println("Hello from resttest")

	error := server.NewServer().ListenAndServe()
	_ = fmt.Errorf("error: %v", error)
}
