package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nethttp/src/article"
	"nethttp/src/httpconnector"
	"nethttp/src/httpserver"
)

func main() {
	fmt.Println("Start")

	/** GET **/

	body := httpconnector.Get("https://jsonplaceholder.typicode.com/posts")
	// log.Printf(string(body))
	var posts []article.Post
	json.Unmarshal(body, &posts)
	log.Println(posts[0].Title)

	/** POST **/

	article := article.Article{
		Id:      "3",
		Title:   "Newly Created Post",
		Desc:    "The description for my new post",
		Content: "my articles content"}

	body2 := httpconnector.Post("https://postman-echo.com/post", article)

	res := string(body2)
	log.Printf(res)

	/** SERVE **/

	httpserver.Serve()
}
