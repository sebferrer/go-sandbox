package httpserver

import (
	"net/http"
)

func Routing() {
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/headers", Headers)
	http.HandleFunc("/testget", TestGet)
	// http.HandleFunc("/testget/{value}", TestGet2)
	http.HandleFunc("/testpost", TestPost) /*.Methods("POST")*/

	http.ListenAndServe(":8090", nil)
}
