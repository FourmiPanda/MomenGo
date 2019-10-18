package main

import (
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
)

func main ()  {
	c := entities.Configuration{
		Broker:   entities.Broker{Url:"tcp://localhost",Port:"1883"},
		Capteurs: nil,
		Redis:    entities.RedisDB{Network: "tcp", Address: "localhost:6379"},
	}
	m := redisMqtt.CreateAMqttClientFromABroker(&c)

	// Le client MQTT s'abonne à tous les topics contenu dnas capteurs
	m.SubscribeAToATopic("/capteurs/#")

	// le client MQTT publie sur un topic un json
	// le client qui sera abonné au topic récupérera le message et l'ajoutera à la base redis
	m.PublishAMessage("/capteurs/BTZ/TEMP/1",
		`{
			"value": 10,
			"timestamp":"2007-03-01T13:00:00Z"
		},
		{
			"value":32.1,
			"timestamp":"2007-03-01T13:00:00Z"
		}`)
	// This message will fail beacause a } is missing at the end
	m.PublishAMessage("/capteurs/BIA/TEMP/1",
		`{
			"value":32.1,
			"timestamp":"2008-03-01T13:00:00Z"
		`)
	// This message will fail beacause the timestamp is not correctly formatted
	m.PublishAMessage("/capteurs/BIA/TEMP/1",
		`{
			"value":32.1,
			"timestamp":"2008-03-01-13-00-00"
		}`)
	// This message will fail beacause the two entries are not of the same day
	m.PublishAMessage("/capteurs/BIA/TEMP/1",
		`{
			"value": 10,
			"timestamp":"2007-03-01T13:00:00Z"
		},
		{
			"value":32.1,
			"timestamp":"2008-03-01T13:00:00Z"
		}`)
	for true {
		continue
	}
}