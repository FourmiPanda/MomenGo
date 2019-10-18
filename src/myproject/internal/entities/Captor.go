package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
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
func (c *Captor) AddValuesFromJson(jsonValues []byte) *Captor {
	//fmt.Println("DEBUG :", "AddValuesFromJson")
	/* jsonValues is supposed to have this format :
		{
			"value": 10,
			"timestamp": "2007-03-01T13:00:00Z"
		},
		{
			"value":32.1,
			"timestamp":"2008-03-01T13:00:00Z"
		}
	*/
	// Remove space frome the payload and convert it to string

	p := strings.Join(
		strings.Fields(string(jsonValues)),
		"")
	//// Create a slice of every values contained in the payload
	ps := strings.Split(p,"},")
	end := len(ps)
	//For Debug purpose
	//fmt.Println("DEBUG : p =", p)
	//fmt.Println("DEBUG : ps =", ps)
	for i := 0; i < end ; i++ {
		// add the "}" lost during the split
		if i != end - 1{
			ps[i] += "}"
		}
		//fmt.Println("DEBUG : ps[",i,"] =", ps[i])
		c.Values = append(c.Values, CreateACaptorValue([]byte(ps[i])))
	}
	fmt.Println(c.CaptorToString())
	return c
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
	return  string(c.CaptorToJson())
}
//func (c *Captor) CaptorToString() string {
//	return  `{` +
//		`"idCaptor":` 	+ c.GetIdCaptorToString()	+ `,` 	+
//		`"idAirport":"` + c.GetIdAirportToString()	+ `",` 	+
//		`"measure":"` 	+ c.GetMeasureToString() 	+ `",` 	+
//		`"values":`		+ c.GetValuesToString() 	+
//		`}`
//}
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