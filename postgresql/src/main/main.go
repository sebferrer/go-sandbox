package main

import (
	"fmt"
	psql "psqlsample/src/dbconnector"
	// httpserver "psqlsample/src/httpserver/gin"
)

func main() {
	fmt.Println("Start")

	psql.TestDB()

	// articles, _ := psql.GetArticles()
	// fmt.Printf("%s\n", articles)

	// article, _ := psql.GetArticle(1)
	// fmt.Printf("%s\n", article)

	// subArticle, _ := psql.GetSubArticle("1")
	// fmt.Printf("%s\n", subArticle)

	// id, _ := psql.AddSubArticle([]byte("{\"id\":\"5\",\"published\":true,\"authors\":[\"charlie\"],\"categories\":[\"category3\"],\"tags\":[\"cats\"]}"))
	// fmt.Printf("%d\n", id)

	// id, _ := psql.DeleteSubArticle("5")
	// fmt.Printf("%d\n", id)

	id, _ := psql.UpdateSubArticle("5", []byte("{\"id\":\"5\",\"published\":true,\"authors\":[\"charlie\"],\"categories\":[\"category3\"],\"tags\":[\"dogs\"]}"))
	fmt.Printf("%d\n", id)

	// httpserver.Serve()
}
