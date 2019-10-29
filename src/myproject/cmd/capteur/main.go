/**
 * CAPTORS
 *
 * @description :: Init the CAPTORS. Captors publish values to the MQTT BROKER.
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"myproject/internal/entities"
	"os"
	"path/filepath"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	conf := getConfig()
	client := connect(conf.Broker.Url+":"+conf.Broker.Port, "capteurs")

	for {
		for _, v := range conf.Capteurs {
			fmt.Println(time.Now().Format("2006-01-02T15:04:05Z") + " : [CAPTEUR-" + strconv.Itoa(v.Id) + "] : Publishing data for airport " + v.IATA)
			client.Publish("/capteurs/"+v.IATA+"/"+v.Type+"/"+strconv.Itoa(v.Id), byte(v.QoS), false,
				"{\"timestamp\" : \""+time.Now().Format("2006-01-02T15:04:05Z")+"\", \"value\": "+strconv.Itoa(v.GetValue())+" }")
		}
		duration := time.Duration(10) * time.Second
		time.Sleep(duration)
	}

}

func getConfig() entities.Configuration {
	configPath, _ := filepath.Abs("src/config/config.json")
	file, err := os.Open(configPath)
	if err != nil {
		handleError(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := entities.Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		handleError(err)
	}
	return configuration
}

func connect(brokerURI string, clientId string) mqtt.Client {
	fmt.Println("Trying to connect (" + brokerURI + ", " + clientId + ")...")
	opts := createClientOptions(brokerURI, clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		handleError(err)
	}
	return client
}

func createClientOptions(brokerURI string, clientId string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	// AddBroker adds a broker URI to the list of brokers to be used.
	// The format should be "scheme://host:port"
	opts.AddBroker(brokerURI)
	// opts.SetUsername(user)
	// opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func handleError(err error) {
	log.Fatal(err)
}
