package entities

import (
	"encoding/json"
	"errors"
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
func (c *Captor) AddValuesFromJson(jsonValues []byte) (*Captor,error) {
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
	var val *CaptorValue
	var err error
	for i := 0; i < end ; i++ {
		// add the "}" lost during the split
		if i != end - 1{
			ps[i] += "}"
		}
		//fmt.Println("DEBUG : ps[",i,"] =", ps[i])
		val, err = CreateACaptorValue([]byte(ps[i]))
		if err != nil {
			break
		}
		c.Values = append(c.Values, val)
	}
	//fmt.Println("DEBUG : c.CaptorToString ",c.CaptorToString())
	return c, err
}
func (c *Captor) IsEmpty() bool {
	res := true
	for i := 0 ; i < len(c.Values) ; i++ {
		res = res && c.Values[i].IsEmpty()
	}
	return res
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
	return  string(c.CaptorToSliceByte())
}
func (c *Captor) CaptorToJson() string {
	res,err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}
//func (c *Captor) CaptorToString() string {
//	return  `{` +
//		`"idCaptor":` 	+ c.GetIdCaptorToString()	+ `,` 	+
//		`"idAirport":"` + c.GetIdAirportToString()	+ `",` 	+
//		`"measure":"` 	+ c.GetMeasureToString() 	+ `",` 	+
//		`"values":`		+ c.GetValuesToString() 	+
//		`}`
//}
func (c *Captor) CaptorToSliceByte() []byte{
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
func (c *Captor) GetDayDate(idVal int) string {
	return c.Values[idVal].GetDayDate()
}
func (c *Captor) GetDayDateAsInt(idVal int) int{
	return c.Values[idVal].GetDayDateAsInt()
}
func (c *Captor) GetMean() (float64, error) {
	len := len(c.Values)
	if len == 0 {
		return 0, errors.New("WARNING : " +
			c.GetIdAirportToString()+ ":" +
			c.GetMeasureToString() 	+ ":" +
			c.GetIdCaptorToString() +  " has no values for the mean.")
	}

	var res, sum float64
	for _, i := range c.Values {
		sum += i.Value
	}
	res = sum / float64(len)
	return res, nil
}
func GetSliceMean(captors []Captor) (float64, error) {
	len := len(captors)
	if len == 0 {
		return 0, errors.New("There is no captor in this slice.")
	}

	var res, sum float64
	var err error
	for _, i := range captors {
		val, err2 := i.GetMean()
		if err2 != nil {
			err = err2
		}
		sum += val
	}
	res = sum / float64(len)
	return res, err
}
func MergeEqualsCaptors(captors []Captor) []Captor {
	var res []Captor
	for _, i := range captors {
		if len(res) == 0 {
			res = append(res, i)
			continue
		}
		for _, j := range res {
			if j.IdCaptor == i.IdCaptor && j.IdAirport == i.IdAirport && j.Measure == i.Measure {
				j.Values = append(j.Values, i.Values...)
				break
			}
		}
	}
	return res
}
