package main

import (
	"fmt"
	httpserver "psqlsample/src/httpserver/gin"
	psql "psqlsample/src/dbconnector"
)

func main() {
	fmt.Println("Start")

	psql.TestDB()

	httpserver.Serve()
}