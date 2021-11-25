package main

import (
	"encoding/json"
	"fmt"
	"log"

	"nethttp/src/article"
	"nethttp/src/httpconnector"
)

func main() {
	fmt.Println("Start")

	// httpserver.Serve()

	postsBody := httpconnector.Get("https://jsonplaceholder.typicode.com/posts")
	// log.Printf(string(postsBody))
	var posts []article.Post
	json.Unmarshal(postsBody, &posts)
	log.Println(posts[0].Title)
}
