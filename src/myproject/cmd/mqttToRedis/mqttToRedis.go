/**
 * BROKER
 *
 * @description :: Init the MQTT BROKER.
 */

package main

import (
	"fmt"
	"myproject/internal/entities"
)

func main() {
	c := entities.GetConfig()
	m := entities.CreateAMqttClientFromConfig(c)

	// Le client MQTT s'abonne à tous les topics contenu dnas capteurs
	m.SubscribeAToATopic("/capteurs/#")
	fmt.Println("Listening on topic : /capteurs/# ...")

	for {
		continue
	}
}
