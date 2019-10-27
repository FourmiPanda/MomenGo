package api

import (
	"fmt"
	"log"
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Mean(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query() // Parse query & return Values
	date := strings.Split(queryValues.Get("date"),"-")

	y, errY := strconv.Atoi(date[0])
	m, errM := strconv.Atoi(date[1])
	d, errD := strconv.Atoi(date[2])

	switch {
	case errY != nil:
		log.Fatal(errY)
	case errM != nil:
		log.Fatal(errM)
	case errD != nil:
		log.Fatal(errD)
	}

	rc := redisMqtt.CreateARedisClientFromConfig(entities.Config.Redis)
	temp , _ := rc.GetAllCaptorValuesOfTempForADay(time.Date(y,time.Month(m),d,0,0,0,0,time.UTC))
	pres , _ := rc.GetAllCaptorValuesOfPresForADay(time.Date(y,time.Month(m),d,0,0,0,0,time.UTC))
	wind , _ := rc.GetAllCaptorValuesOfWindForADay(time.Date(y,time.Month(m),d,0,0,0,0,time.UTC))

	meanT ,_ := entities.GetSliceMean(temp)
	meanP ,_ := entities.GetSliceMean(pres)
	meanW ,_ := entities.GetSliceMean(wind)
	res := `{
				{
					"temperature mean":`+ strconv.FormatFloat(meanT, 'E', -1, 64) +`
				},
				{
					"pressure mean":`	+ strconv.FormatFloat(meanP, 'E', -1, 64) +`
				},
				{
					"wind mean":`		+ strconv.FormatFloat(meanW, 'E', -1, 64) +`
				},
			}`


		// For Debug purpose
	fmt.Println("DEBUG : mean was called")
	fmt.Println("DEBUG : date reveived ",y,m,d)
	fmt.Fprint(w, res)

}
