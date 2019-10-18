package entities

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

type CaptorValue struct {
	Value		float64
	Timestamp  	time.Time
}

func CreateACaptorValue(jsonEntry []byte) *CaptorValue{
	// jsonEntry is supposed to have this format :
	/*
		{
			"value": 10,
			"timestamp": "2007-03-01T13:00:00Z"
		},
		{
			"value":32.1,
			"timestamp":"2008-03-01T13:00:00Z"
		}
	*/
	var e CaptorValue
	err := json.Unmarshal(jsonEntry, &e)
	if err != nil {
		log.Fatal(err)
	}
	return &e
}

//func (c *CaptorValue) GetCaptorValueToString () string {
//	return  `{` +
//		`"value":` 		+ c.GetValueToString() 		+ `,` 	+
//		`"timestamp":"` + c.GetTimestampToString() 	+ `"` 	+
//		`}`
//}
func (c *CaptorValue) GetCaptorValueToString () string {
	return  string(c.GetCaptorValueToJson())
}

func (c *CaptorValue) GetValueToString () string {
	return  strconv.FormatFloat(c.Value, 'E', -1, 64)
}

func (c *CaptorValue) GetTimestampToString () string {
	return c.Timestamp.String()
}

func (c *CaptorValue) GetCaptorValueToJson () []byte {
	res,err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return res
}