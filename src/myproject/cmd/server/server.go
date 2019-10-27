package main

import (
	"log"
	"myproject/cmd/api"
	"net/http"
)

func main()  {
	http.HandleFunc("/mean", api.Mean)

	err := http.ListenAndServe(":2019",nil)

	if err != nil {
		log.Fatal(err)
	}
}


