/**
 * API
 *
 * @description :: Init the REST API.
 */
package main

import (
	"fmt"
	"log"
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("** REST API is listening on 0.0.0.0:2019 **")

	http.HandleFunc("/search", search)
	err := http.ListenAndServe(":2019", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[" + time.Now().String() + "] : Incoming request on ' " + r.URL.Path + "?" + r.URL.RawQuery + "'")

	data := r.URL.Query()

	startDate := strings.Split(data.Get("start_date"), "-")
	if len(startDate) != 3 {
		badRequest(w)
		return
	}
	endDate := strings.Split(data.Get("end_date"), "-")
	if len(endDate) != 3 {
		badRequest(w)
		return
	}
	iata := data.Get("iata")
	if iata == "" {
		iata = "*"
	}
	measureType := data.Get("type")
	if measureType == "" {
		measureType = "*"
	}
	moyenne := false
	if data.Get("moyenne") == "" || data.Get("moyenne") == "true" {
		moyenne = true
	}

	//TODO: Write the redis query

	rc := redisMqtt.CreateARedisClientFromConfig(entities.GetConfig())
	res, err := rc.Find("CaptorValues:" + iata + ":" + measureType)
	if err != nil {
		badRequest(w)
		return
	}

	fmt.Println(res)

	if moyenne {

	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("[]"))

}

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(400)
	_, _ = w.Write([]byte("Bad request"))
}
