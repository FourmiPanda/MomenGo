package main

import (
	"MomenGo/src/myproject/cmd/redisMqtt"
	"MomenGo/src/myproject/internal/entities"
)

func main()  {
	// To create un RedisClient you can use a config
	config := entities.RedisDB{Network: "tcp", Address: "localhost:6379"}

	// To simulate a mqqt message we need a json
	j := []byte(
		`{
			"idCaptor":1,
			"idAirport":"AAI",
			"measure":"TEST",
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
	println("This is the Mqtt message received :\t", m.MqttMessageToString())

	// Create a RedisEntry with the mqtt message receive
	r := entities.CreateARedisEntryFromMqtt(m)
	println("This is the RedisEntry produced :\t", r.RedisEntryToString())

	// Create a RedisClient with the config
	rc := redisMqtt.CreateARedisClientFromConfig(config)

	// Add the RedisEntry to the redis DB
	rc.AddCaptorEntryToDB(r)

	// Get the entry of the captor created
	println(rc.GetACaptorAttributeEntry(r.CaptorKey(), "idAirport"))

	// Get the entry of the captor created
	v := rc.GetAllCaptorValuesEntry(r.CaptorValuesKey())
	for i := 0 ; i < len(v) ; i++ {
		println(v[i])
	}
}
