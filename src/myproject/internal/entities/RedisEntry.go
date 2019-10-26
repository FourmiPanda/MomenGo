package entities

import (
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


func (r *RedisEntry) RedisEntryToJson () string{
	return r.Captor.CaptorToJson()
}
func (r *RedisEntry) RedisEntryToSliceByte () []byte{
	return r.Captor.CaptorToSliceByte()
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
func (r *RedisEntry) GetCaptorValueAsJson (idVal int) string {
	return r.Captor.Values[idVal].GetCaptorValueToString()
}
func (r *RedisEntry) CaptorKey() string {
	return "Captor:" + r.GetCaptorLeftKey()
}
func (r *RedisEntry) CaptorValuesKey(idVal int) string {
	return "CaptorValues:" + r.GetCaptorLeftKey() + ":" + r.GetDayDate(idVal)
}
func (r *RedisEntry) GetCaptorLeftKey() string {
	return r.Captor.GetIdAirportToString() + ":" + r.Captor.GetMeasureToString() + ":" + r.Captor.GetIdCaptorToString()
}
func (r *RedisEntry) GetDayDate(idVal int) string {
	y,m,d := r.Captor.Values[idVal].Timestamp.Date()
	return strconv.Itoa(y) + ":" + strconv.Itoa(int(m)) + ":" + strconv.Itoa(d)
}
func (r *RedisEntry) GetDayDateAsInt(idVal int) int{
	y,m,d := r.Captor.Values[idVal].Timestamp.Date()
	return int (time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix())
}
func (r *RedisEntry) GetTimestampAsInt(idVal int) int{
	return int (r.Captor.Values[idVal].Timestamp.Unix())
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