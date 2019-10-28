/**
 * BROKER
 *
 * @description :: Init the MQTT BROKER.
 */
package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//TODO: Check if you can start a mqtt broker with paho
	wg.Add(1)

	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	wg.Wait()

}
