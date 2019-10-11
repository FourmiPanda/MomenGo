package entities

import (
	"encoding/json"
	"log"
	"strconv"
)

/*
## Structure des messages MQTT
```json
{
 "idCaptor":1,
 "idAirport":"BIA",
 "measure":"Temperature",
 "value":27,
 "timestamp":"2007-03-01T13:00:00Z"
}
```
 */

type RedisEntry struct {
	IdCaptor 	int
	IdAirport 	string
	Measure		string
	Value		float64
	Timestamp  	string
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
				`"idCaptor":` 	+ strconv.Itoa(r.IdCaptor) 					+ `,` 	+
				`"idAirport":"` + r.IdAirport 								+ `",` 	+
				`"measure":"` 	+ r.Measure 								+ `",` 	+
				`"value":` 		+ strconv.FormatFloat(r.Value, 'E', -1, 64) + `,` 	+
				`"timestamp":"` + r.Timestamp 								+ `"` 	+
			`}`

}
func (r *RedisEntry) RedisEntryToByte () []byte{
	return []byte(r.RedisEntryTOString())
}

func (r *RedisEntry) PrintAll() {
	println(r.IdCaptor)
	println(r.IdAirport)
	println(r.Measure)
	println(r.Value)
	println(r.Timestamp)
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