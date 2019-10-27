/**
 * API
 *
 * @description :: Init the REST API.
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// TODO: Start listening for incoming HTTP requests
	fmt.Println("** REST API is listening on 0.0.0.0:2019 **")

	http.HandleFunc("/mean", GetMean)
	http.HandleFunc("/search", search)
	err := http.ListenAndServe(":2019", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	//TODO: Handle search request
	fmt.Println("[" + time.Now().String() + "] : Incoming request on ' " + r.URL.Path + "'")

	data := r.URL.Query()

	startDate := data.Get("start_date")
	endDate := data.Get("end_date")
	iata := data.Get("iata")
	measureType := data.Get("type")
	moyenne := data.Get("moyenne")

	fmt.Println(startDate + " " + endDate + " " + iata + " " + measureType + " " + moyenne)

	rc := redisMqtt.CreateARedisClientFromConfig(entities.GetConfig())

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte("{\"value:\" : \"1\"}"))

}

func GetMean(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query() // Parse query & return Values
	date := strings.Split(queryValues.Get("date"), "-")

	y, errY := strconv.Atoi(date[0])
	m, errM := strconv.Atoi(date[1])
	d, errD := strconv.Atoi(date[2])

	switch {
	case errY != nil:
		log.Println(errY)
		fmt.Fprint(w, errY)
		return
	case errM != nil:
		log.Println(errM)
		fmt.Fprint(w, errM)
		return
	case errD != nil:
		log.Println(errD)
		fmt.Fprint(w, errD)
		return
	}

	rc := redisMqtt.CreateARedisClientFromConfig(entities.GetConfig())
	temp, errT := rc.GetAllCaptorValuesOfTempForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))
	pres, errP := rc.GetAllCaptorValuesOfPresForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))
	wind, errW := rc.GetAllCaptorValuesOfWindForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))

	switch {
	case errT != nil:
		log.Println(errT)
		fmt.Fprint(w, errT)
		return
	case errP != nil:
		log.Println(errP)
		fmt.Fprint(w, errP)
		return
	case errW != nil:
		log.Println(errW)
		fmt.Fprint(w, errW)
		return
	}

	meanT, errMT := entities.GetSliceMean(temp)
	meanP, errMP := entities.GetSliceMean(pres)
	meanW, errMW := entities.GetSliceMean(wind)

	switch {
	case errMT != nil:
		log.Println(errMT)
		fmt.Fprint(w, errMT)
		return
	case errMP != nil:
		log.Println(errMP)
		fmt.Fprint(w, errMP)
		return
	case errMW != nil:
		log.Println(errMW)
		fmt.Fprint(w, errMW)
		return
	}

	meanT = math.Round(meanT*100) / 100
	meanP = math.Round(meanP*100) / 100
	meanW = math.Round(meanW*100) / 100

	j := entities.CreateMean(
		meanT,
		meanP,
		meanW,
	)
	res, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	// For Debug purpose
	fmt.Println("DEBUG : mean was called")
	fmt.Println("DEBUG : date reveived ", y, m, d)
	fmt.Fprint(w, string(res))

}
