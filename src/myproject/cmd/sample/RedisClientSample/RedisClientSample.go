package main

import (
	"fmt"
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
	"time"
)

func main() {
	// To create un RedisClient you can use a config
	config := entities.RedisDB{Network: "tcp", Address: "localhost:6379"}

	// To simulate a mqqt message we need a json
	j := []byte(
		`{
			"idCaptor":1,
			"idAirport":"AAI",
			"measure":"TEMP",
			"values":[
				{
					"value":27.8,
					"timestamp":"2007-03-01T13:05:05Z"
				},
				{
					"value":37.8,
					"timestamp":"2007-03-01T13:06:05Z"
				},
				{
					"value":207.8,
					"timestamp":"2007-03-01T13:07:05Z"
				},
				{
					"value":32.1,
					"timestamp":"2008-03-01T13:05:05Z"
				},
				{
					"value":32.1,
					"timestamp":"2008-03-01T13:05:05Z"
				}
			]
		}`)
	m := entities.CreateAMqttMessageFromByte(j)
	println("This is the Mqtt message received :\t", m.MqttMessageToString())

	// Create a RedisEntry with the mqtt message receive
	r := entities.CreateARedisEntryFromMqtt(m)
	println("This is the RedisEntry produced :\t", r.RedisEntryToString())
	println("This is the key of the Captor hash")
	println(r.CaptorKey())

	// Create a RedisClient with the config
	rc := redisMqtt.CreateARedisClientFromConfig(config)

	// Add the RedisEntry to the redis DB
	rc.AddCaptorEntryToDB(r)

	// Get the keys corresponding to the interval between 01-01-2007 and 01-03-2008
	keys, _ := rc.GetCaptorValuesKeysInInterval(
		[]string{r.CaptorKey()},
		time.Date(2007,01,01,0,0,0,0,time.UTC),
		time.Date(2008,03,01,0,0,0,0,time.UTC))
	fmt.Println("Keys between 2000 and 2019",keys)

	// Get the values correspond to the interval chosen before
	var res []entities.Captor
	for i := 0; i < len(keys); i++ {
		val, _ := rc.GetAllCaptorValuesOfADay(keys[i])
		res = append(res,val)
	}
	fmt.Println("This are the values stored in the db :")
	for i := 0; i < len(res) ; i++  {
		fmt.Println(res[i].CaptorToString())
	}
}
