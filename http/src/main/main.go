package main

import (
	"fmt"
	"nethttp/src/httpserver"
)

func main() {
	fmt.Println("Start")

	httpserver.Serve()
}
