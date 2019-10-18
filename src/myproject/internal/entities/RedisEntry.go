package entities

import (
	"strconv"
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
	Captor 		*Captor
}

func CreateARedisEntryFromByte(jsonEntry []byte) *RedisEntry{
	r := RedisEntry{
		Captor: CreateACaptor(jsonEntry),
	}
	return &r
}

func CreateARedisEntryFromMqtt(mqtt *MqttMessage) *RedisEntry{
	r := RedisEntry{
		Captor: mqtt.Captor,
	}
	return &r
}

func CreateARedisEntryFromCaptor(captor *Captor) *RedisEntry{
	r := RedisEntry{
		Captor: captor,
	}
	return &r
}


func (r *RedisEntry) RedisEntryToJson () []byte{
	return r.Captor.CaptorToJson()
}
func (r *RedisEntry) RedisEntryToString () string{
	return r.Captor.CaptorToString()
}

func (r *RedisEntry) PrintAll() {
	println(r.Captor.CaptorToString())
}
func (r *RedisEntry) GetCaptorValues () []string{
	return r.Captor.GetCaptorValues()
}
func (r *RedisEntry) GetCaptorHashes () string{
	return "idCaptor " + r.Captor.GetIdCaptorToString() + " idAirport " + r.Captor.GetIdAirportToString() + " measure " + r.Captor.GetMeasureToString() + " values " +r.CaptorValuesKey()
}
func (r *RedisEntry) CaptorKey() string {
	return "Captor:" + r.GetLeftKey()
}
func (r *RedisEntry) CaptorValuesKey() string {
	return "CaptorValues:" + r.GetLeftKey() + ":" + r.GetDayDate()
}
func (r *RedisEntry) GetLeftKey() string {
	return r.Captor.GetIdAirportToString() + ":" + r.Captor.GetMeasureToString() + ":" + r.Captor.GetIdCaptorToString()
}
func (r *RedisEntry) GetDayDate() string {
	y,m,d := r.Captor.Values[0].Timestamp.Date()
	return strconv.Itoa(y) + ":" + strconv.Itoa(int(m)) + ":" + strconv.Itoa(d)
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