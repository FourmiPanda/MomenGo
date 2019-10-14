package entities

import (
	"encoding/json"
	"log"
	"strconv"
)

type Captor struct {
	IdCaptor 	int
	IdAirport 	string
	Measure		string
	Values 		[]CaptorValue
}


func (r *Captor) GetIdCaptorToString () string {
	return strconv.Itoa(r.IdCaptor)
}
func (r *Captor) GetIdAirportToString () string {
	return r.IdAirport
}
func (r *Captor) GetMeasureToString () string {
	return r.Measure
}
func (r *Captor) GetValuesToString () string {
	res := "["
	for i := 0 ; i < len(r.Values) ; i++ {
		res += r.Values[i].GetCaptorValueToString() + ","
	}
	res += "]"
	return res
}
func (r *Captor) CaptorToString() string {
	return  `{` +
		`"idCaptor":` 	+ r.GetIdCaptorToString()	+ `,` 	+
		`"idAirport":"` + r.GetIdAirportToString()	+ `",` 	+
		`"measure":"` 	+ r.GetMeasureToString() 	+ `",` 	+
		`"value":`		+ r.GetValuesToString() 		+
		`}`

}
func CreateACaptor(jsonEntry []byte) *Captor{
	var e Captor
	err := json.Unmarshal(jsonEntry, &e)
	if err != nil {
		log.Fatal(err)
	}
	return &e
}

func (c *Captor) CaptorToJson () []byte{
	res,err := json.Marshal(c.CaptorToString())
	if err != nil {
		log.Fatal(err)
	}
	return res
}