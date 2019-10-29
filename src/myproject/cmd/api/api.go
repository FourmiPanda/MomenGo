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

	http.HandleFunc("/mean", 		 GetMean)
	http.HandleFunc("/valuesByType", GetValuesByType)
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
func GetValuesByType(w http.ResponseWriter, r *http.Request) {
	// Supposed to receive :
	//measure 		: TEMP | PRES | WIND
	//start_date 	: YYYY-MM-DD-HH-MM-SS
	//end_date 		: YYYY-MM-DD-HH-MM-SS
	queryValues := r.URL.Query() // Parse query & return Values
	measure 	:= queryValues.Get("measure")
	startDate 	:= strings.Split(queryValues.Get("start_date"),"-")
	endDate 	:= strings.Split(queryValues.Get("end_date"),"-")
	switch {
	case measure == "":
		fmt.Fprint(w, `{"error":"`+errors.New("measure is not entered").Error()+`"}`)
		return
	case measure != "TEMP" && measure != "PRES" && measure != "WIND":
		fmt.Fprint(w, `{"error":"`+errors.New("measure is not correct").Error()+`"}`)
		return
	case queryValues.Get("start_date") == "":
		fmt.Fprint(w, `{"error":"`+errors.New("start_date is not entered").Error()+`"}`)
		return
	case len(startDate) < 6:
		fmt.Fprint(w, `{"error":"`+errors.New("start_date is not valid").Error()+`"}`)
		return
	case queryValues.Get("end_date") == "":
		fmt.Fprint(w, `{"error":"`+errors.New("end_date is not entered").Error()+`"}`)
		return
	case len(startDate) < 6:
		fmt.Fprint(w, `{"error":"`+errors.New("end_date is not valid").Error()+`"}`)
		return
	}
	sY, errSY := strconv.Atoi(startDate[0])
	sM, errSM := strconv.Atoi(startDate[1])
	sD, errSD := strconv.Atoi(startDate[2])
	sH, errSH := strconv.Atoi(startDate[3])
	sMi, errSMi := strconv.Atoi(startDate[4])
	sS, errSS := strconv.Atoi(startDate[5])

	switch {
	case errSY != nil:
		log.Println(errSY)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSY.Error(),`"`,"", -1)+`"}`)
		return
	case errSM != nil:
		log.Println(errSM)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSM.Error(),`"`,"", -1)+`"}`)
		return
	case errSD != nil:
		log.Println(errSD)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSD.Error(),`"`,"", -1)+`"}`)
		return
	case errSH != nil:
		log.Println(errSH)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSH.Error(),`"`,"", -1)+`"}`)
		return
	case errSMi != nil:
		log.Println(errSMi)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSMi.Error(),`"`,"", -1)+`"}`)
		return
	case errSS != nil:
		log.Println(errSS)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSS.Error(),`"`,"", -1)+`"}`)
		return
	}

	startTime := time.Date(sY,time.Month(sM),sD,sH,sMi,sS,0,time.UTC)


	switch {
	case queryValues.Get("end_date") == "":
		fmt.Fprint(w, `{"error":"`+errors.New("date is not entered").Error()+`"}`)
		return
	case len(startDate) < 3:
		fmt.Fprint(w, `{"error":"`+errors.New("date is not valid").Error()+`"}`)
		return
	}
	eY, errEY := strconv.Atoi(endDate[0])
	eM, errEM := strconv.Atoi(endDate[1])
	eD, errED := strconv.Atoi(endDate[2])
	eH, errEH := strconv.Atoi(endDate[3])
	eMi, errEMi := strconv.Atoi(endDate[4])
	eS, errES := strconv.Atoi(endDate[5])

	switch {
	case errEY != nil:
		log.Println(errEY)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errSD.Error(),`"`,"", -1)+`"}`)
		return
	case errSM != nil:
		log.Println(errSM)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errEM.Error(),`"`,"", -1)+`"}`)
		return
	case errSD != nil:
		log.Println(errSD)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errED.Error(),`"`,"", -1)+`"}`)
		return
	case errEH != nil:
		log.Println(errEH)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errEH.Error(),`"`,"", -1)+`"}`)
		return
	case errEMi != nil:
		log.Println(errEMi)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errEMi.Error(),`"`,"", -1)+`"}`)
		return
	case errES != nil:
		log.Println(errES)
		w.WriteHeader(500)
		fmt.Fprint(w, `{"error":"`+strings.Replace(errES.Error(),`"`,"", -1)+`"}`)
		return
	}
	endTime := time.Date(eY,time.Month(eM),eD,eH,eMi,eS,0,time.UTC)


	rc := entities.CreateARedisClientFromConfig(entities.GetConfig())
	captors, _ := rc.GetAllCaptorValuesOfATypeInInterval(
		measure,
		startTime,
		endTime)

	res, err := json.MarshalIndent(captors,"","    ")
	if err != nil {
		log.Println(err)
		fmt.Fprint(w, err)
		return
	}

	// For Debug purpose
	fmt.Println("DEBUG : mean was called")
	fmt.Println("DEBUG : start_date reveived ", sY, sM, sD)
	fmt.Fprint(w, string(res))

}
