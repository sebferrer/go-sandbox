package httpserver

import (
	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	router.GET("/hello", Hello)
	router.GET("/testget", TestGet)
	router.GET("/testget/:value", TestGet2)
	router.GET("/testget/:value/*action", TestGet3)

	router.Run(":8080")
}
