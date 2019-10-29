/**
 * API
 *
 * @description :: Init the REST API.
 */
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"myproject/internal/entities"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	// TODO: Start listening for incoming HTTP requests

	http.HandleFunc("/mean", GetMean)
	err := http.ListenAndServe(":2019", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetMean(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query() // Parse query & return Values
	date := strings.Split(queryValues.Get("date"), "-")
	switch {
	case queryValues.Get("date") == "":
		fmt.Fprint(w, `{"error":"`+errors.New("date is not entered").Error()+`"}`)
		return
	case len(date) < 3:
		fmt.Fprint(w, `{"error":"`+errors.New("date is not valid").Error()+`"}`)
		return
	}
	// Assign year to y, month to m and day to d
	y, errY := strconv.Atoi(date[0])
	m, errM := strconv.Atoi(date[1])
	d, errD := strconv.Atoi(date[2])

	switch {
	case errY != nil:
		log.Println(errY)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errY.Error(),`"`,"", -1)+`"}`)
		return
	case errM != nil:
		log.Println(errM)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errY.Error(),`"`,"", -1)+`"}`)
		return
	case errD != nil:
		log.Println(errD)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errY.Error(),`"`,"", -1)+`"}`)
		return
	}

	rc := entities.CreateARedisClientFromConfig(entities.GetConfig())
	temp, errT := rc.GetAllCaptorValuesOfTempForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))
	pres, errP := rc.GetAllCaptorValuesOfPresForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))
	wind, errW := rc.GetAllCaptorValuesOfWindForADay(time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC))

	switch {
	case errT != nil:
		log.Println(errT)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errT.Error(),`"`,"", -1)+`"}`)
		return
	case errP != nil:
		log.Println(errP)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errP.Error(),`"`,"", -1)+`"}`)
		return
	case errW != nil:
		log.Println(errW)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errW.Error(),`"`,"", -1)+`"}`)
		return
	}

	meanT, errMT := entities.GetSliceMean(temp)
	meanP, errMP := entities.GetSliceMean(pres)
	meanW, errMW := entities.GetSliceMean(wind)

	switch {
	case errMT != nil:
		log.Println(errMT)
		meanT = 0
	case errMP != nil:
		log.Println(errMP)
		meanP = 0
	case errMW != nil:
		log.Println(errMW)
		meanW = 0
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
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+err.Error()+`"}`)
		return
	}

	// For Debug purpose
	fmt.Println("DEBUG : mean was called")
	fmt.Println("DEBUG : date reveived ", y, m, d)
	fmt.Fprint(w, string(res))

}
