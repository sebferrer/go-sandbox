package main

import (
	"fmt"
	httpserver "nethttp/src/httpserver/gin"
)

func main() {
	fmt.Println("Start")

	httpserver.Serve()
}
