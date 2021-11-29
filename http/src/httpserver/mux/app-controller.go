package httpserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nethttp/src/article"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func TestGet(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	qValue, ok := req.URL.Query()["value"]
	value := qValue[0]

	if !ok || len(value) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	fmt.Fprintf(w, "{ \"value\": "+value+" }")
}

func TestGet2(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	value := mux.Vars(req)["value"]

	fmt.Fprintf(w, "{ \"value\": "+value+" }")
}

func TestPost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var article article.Article

	reqBody, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(reqBody, &article)
	// fmt.Fprintf(w, "%+v", article.Title)
	fmt.Fprintf(w, "%+v", article)
}
