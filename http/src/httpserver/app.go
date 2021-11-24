package httpserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/hello", Hello)
	router.HandleFunc("/headers", Headers)
	router.HandleFunc("/testget", TestGet)
	router.HandleFunc("/testget/{value}", TestGet2)
	router.HandleFunc("/testpost", TestPost).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", router))
}
