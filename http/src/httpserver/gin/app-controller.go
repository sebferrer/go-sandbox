package httpserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"nethttp/src/article"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func TestGet(c *gin.Context) {
	qValue, ok := c.Request.URL.Query()["value"]
	value := qValue[0]
	if !ok || len(value) < 1 {
		log.Println("Url Param 'key' is missing")
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

func TestPost(c *gin.Context) {
	// b := c.FullPath() == "/testpost/:name/*action" // true
	var article article.Article

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(reqBody, &article)

	// c.String(http.StatusOK, article.Title)
	c.String(http.StatusOK, article.String())
}
