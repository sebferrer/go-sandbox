package main

import (
	"fmt"
	psql "psqlsample/src/dbconnector"
	httpserver "psqlsample/src/httpserver/gin"
)

func main() {
	fmt.Println("Start")

	psql.TestDB()

	// articles := psql.GetArticles()
	// fmt.Printf("%s\n", articles)

	// article := psql.GetArticle(1)
	// fmt.Printf("%s\n", article)

	subArticle := psql.GetSubArticle("1")
	fmt.Printf("%s\n", subArticle)

	httpserver.Serve()
}
