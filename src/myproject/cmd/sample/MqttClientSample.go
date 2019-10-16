package main

import (
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
)

func main ()  {
	b := entities.Broker{Url:"tcp://localhost",Port:"8080"}
	m := redisMqtt.CreateAMqttClientFromABroker(&b)
	m.PublishAMessage("lol")
}