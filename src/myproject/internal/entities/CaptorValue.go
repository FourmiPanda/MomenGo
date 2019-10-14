package entities

import (
	"strconv"
	"time"
)

type CaptorValue struct {
	Value		float64
	Timestamp  	time.Time
}

func (r *CaptorValue) GetCaptorValueToString () string {
	return  `{` +
		`"value":` 		+ r.GetValueToString() 		+ `,` 	+
		`"timestamp":"` + r.GetTimestampToString() 	+ `"` 	+
		`}`
}

func (r *CaptorValue) GetValueToString () string {
	return  strconv.FormatFloat(r.Value, 'E', -1, 64)
}

func (r *CaptorValue) GetTimestampToString () string {
	return r.Timestamp.String()
}