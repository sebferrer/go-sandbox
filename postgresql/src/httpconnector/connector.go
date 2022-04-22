package httpconnector

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func Post(url string, obj interface{}) []byte {
	postBody, _ := json.Marshal(obj)
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
