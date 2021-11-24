package main

import (
	"fmt"
	"net/http"
	"nethttp/src/httpserver"
)

func main() {
	fmt.Println("Start")

	httpserver.Routing()
	http.ListenAndServe(":8090", nil)
}
