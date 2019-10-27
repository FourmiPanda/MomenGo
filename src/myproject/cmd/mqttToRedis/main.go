package main

import (
	"myproject/cmd/redisMqtt"
	"myproject/internal/entities"
)

func main() {
	c := entities.GetConfig()
	m := redisMqtt.CreateAMqttClientFromConfig(c)

	// Le client MQTT s'abonne à tous les topics contenu dnas capteurs
	m.SubscribeAToATopic("/capteurs/#")

	for {
		continue
	}
}
