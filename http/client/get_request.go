package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

func get(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray
}
