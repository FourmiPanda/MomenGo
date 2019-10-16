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

func (c *CaptorValue) GetCaptorValueToString () string {
	return  `{` +
			`"value":` 		+ c.GetValueToString() 		+ `,` 	+
			`"timestamp":"` + c.GetTimestampToString() 	+ `"` 	+
		`}`
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