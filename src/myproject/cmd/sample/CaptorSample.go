package main

import "myproject/internal/entities"

func main()  {
	// How to create a Captor
	// You need a JSON entry
	// To have a JSON just create a byte array with JSON Formatted string
	e := []byte(
		`{
			"idCaptor":1,
			"idAirport":"AAI",
			"measure":"Temperature",
			"values":[
				{
					"value":27.8,
					"timestamp":"2007-03-01T13:00:00Z"
				},
				{
					"value":32.1,
					"timestamp":"2008-03-01T13:00:00Z"
				}
			]
		}`)
	c := entities.CreateACaptor(e)
	println(c.CaptorToString())

	// To convert a Captor into a json byte array
	c = entities.CreateACaptor(c.CaptorToJson())
	println(c.CaptorToString())

	// To get access to Captor's values
	println("idAirport: ","\t",c.IdAirport)
	println("idCaptor: ","\t",c.IdCaptor)

	// As values can contain multiple value
	// So you can call a single value or use Captor's method to show them all
	println("First value of the captor : ",c.Values[0].GetValueToString())
	println("All value of the captor : ",c.GetValuesToString())
}
