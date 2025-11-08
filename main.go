package main

import (
	"fmt"
	"resttest/webserver"
)

func main() {
	fmt.Println("Hello from resttest")

	error := webserver.NewServer().ListenAndServe()
	_ = fmt.Errorf("error: %v", error)
}
