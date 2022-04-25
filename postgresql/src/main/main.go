package main

import (
	"fmt"
	psql "psqlsample/src/dbconnector"
	httpserver "psqlsample/src/httpserver/gin"
)

func main() {
	fmt.Println("Start")

	psql.TestDB()

	psql.GetArticle(1)

	httpserver.Serve()
}
