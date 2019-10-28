package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/mean", api.GetMean)

	err := http.ListenAndServe(":2019", nil)

	if err != nil {
		log.Fatal(err)
	}
}
