package httpserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func TestGet(c *gin.Context) {
	qValue, ok := c.Request.URL.Query()["value"]
	value := qValue[0]
	if !ok || len(value) < 1 {
		log.Println("Url Param 'value' is missing")
		return
	}

	message := "{ \"value\": " + value + " }"
	c.String(http.StatusOK, message)
}

func TestGet2(c *gin.Context) {
	value := c.Param("value")
	message := "{ \"value\": " + value + " }"
	c.String(http.StatusOK, message)
}

func TestGet3(c *gin.Context) {
	value := c.Param("value")
	action := c.Param("action")
	message := "{ \"value\": " + value + ", \"action\": " + action + " }"
	c.String(http.StatusOK, message)
}
