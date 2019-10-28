package main

import (
	"myproject/internal/entities"
	"strings"
)

func main() {
	// How to create a MqttMessage
	// You need a JSON entry
	// To have a JSON just create a byte array with JSON Formatted string
	j := []byte(
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
	m := entities.CreateAMqttMessageFromByte(j)
	println(m.MqttMessageToString())
	println(strings.Join(strings.Fields(string(j)), ""))

	// Or you can use a Captor
	m = entities.CreateAMqttMessage(entities.CreateACaptor(j))
	println(m.MqttMessageToString())

	c := entities.CreateACaptor(m.MqttMessageToJson())
	println(c.CaptorToString())

	// To convert a MqttMessage into a json byte array
	println("MqttMessage in JSON: ", "\t", string(m.MqttMessageToJson()))
	println("Original JSON", "\t\t", strings.Join(strings.Fields(string(j)), ""))
}
