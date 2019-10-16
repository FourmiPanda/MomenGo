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
	Values 		[]*CaptorValue
}


func CreateACaptor(jsonEntry []byte) *Captor{
	var e Captor
	err := json.Unmarshal(jsonEntry, &e)
	if err != nil {
		log.Fatal(err)
	}
	return &e
}
func (c *Captor) GetIdCaptorToString () string {
	return strconv.Itoa(c.IdCaptor)
}
func (c *Captor) GetIdAirportToString () string {
	return c.IdAirport
}
func (c *Captor) GetMeasureToString () string {
	return c.Measure
}
func (c *Captor) GetValuesToString () string {
	res := "["
	for i := 0 ; i < len(c.Values) ; i++ {
		res += c.Values[i].GetCaptorValueToString()
		if i != (len(c.Values) - 1) {
			res += ","
		}
	}
	res += "]"
	return res
}
func (c *Captor) CaptorToString() string {
	return  `{` +
		`"idCaptor":` 	+ c.GetIdCaptorToString()	+ `,` 	+
		`"idAirport":"` + c.GetIdAirportToString()	+ `",` 	+
		`"measure":"` 	+ c.GetMeasureToString() 	+ `",` 	+
		`"values":`		+ c.GetValuesToString() 	+
		`}`

}
func (c *Captor) CaptorToJson () []byte{
	res,err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func (c *Captor) GetCaptorValues() []string {
	var values []string
	for i := 0 ; i < len(c.Values) ; i++ {
		values = append(values, string(c.Values[i].GetCaptorValueToJson()))
	}
	return values
}