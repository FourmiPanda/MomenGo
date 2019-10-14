package entities

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
	Captor 		Captor
}

func CreateARedisEntry(jsonEntry []byte) *RedisEntry{
	r := RedisEntry{
		Captor: *CreateACaptor(jsonEntry),
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