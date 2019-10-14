package entities

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

/*
## Structure des messages MQTT
```json
{
 "idCaptor":  1,
 "idAirport": "BIA",
 "measure":   "Temperature",
 "values":     [
       {
        "value": 27.8,
        "time":  "2007-03-01T13:00:00Z"
       },
       {
        "value": 21.9,
        "time":  "2008-03-01T13:00:00Z"
       }
     ]
}
```
 */

type RedisEntry struct {
	IdCaptor 	int
	IdAirport 	string
	Measure		string
	Values		[]RedisCaptorValue
}

func (r *RedisEntry) getIdCaptorToString () string {
	return strconv.Itoa(r.IdCaptor)
}
func (r *RedisEntry) getIdAirportToString () string {
	return r.IdAirport
}
func (r *RedisEntry) getMeasureToString () string {
	return r.Measure
}
func (r *RedisEntry) getValuesToJson () string {
	res := "["
	for i := 0 ; i < len(r.Values) ; i++ {
		res += r.Values[i].getRedisCaptorToJson() + ","
	}
	res += "]"
	return res
}

type RedisCaptorValue struct {
	Value		float64
	Timestamp  	time.Time
}

func (r *RedisCaptorValue) getRedisCaptorToJson () string {
	return  `{` +
			`"value":` 		+ r.getValueToString() 		+ `,` 	+
			`"timestamp":"` + r.getTimestampToString() 	+ `"` 	+
		`}`
}

func (r *RedisCaptorValue) getValueToString () string {
	return  strconv.FormatFloat(r.Value, 'E', -1, 64)
}

func (r *RedisCaptorValue) getTimestampToString () string {
	return r.Timestamp.String()
}



func CreateARedisEntry(jsonEntry []byte) *RedisEntry{
	var e RedisEntry
	err := json.Unmarshal(jsonEntry, &e)
	if err != nil {
		log.Fatal(err)
	}
	return &e
}

func (r *RedisEntry) RedisEntryTOString() string {
	return  `{` +
				`"idCaptor":` 	+ r.getIdCaptorToString()	+ `,` 	+
				`"idAirport":"` + r.getIdAirportToString()	+ `",` 	+
				`"measure":"` 	+ r.getMeasureToString() 	+ `",` 	+
				`"value":`		+ r.getValuesToJson() 		+
			`}`

}
func (r *RedisEntry) RedisEntryToByte () []byte{
	return []byte(r.RedisEntryTOString())
}

func (r *RedisEntry) PrintAll() {
	println(r.RedisEntryTOString())
}

//func main()  {
//	e := []byte(
//		`{
//		"idCaptor":1,
//		"idAirport":"AAI",
//		"measure":"Temperature",
//		"value":27.2,
//		"timestamp":"2007-03-01T13:00:00Z"
//	}`)
//}