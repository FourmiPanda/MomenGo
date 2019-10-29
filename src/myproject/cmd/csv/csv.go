package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"myproject/internal/entities"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {

	file, _ := os.Create("donneesMeteo.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"IATA", "Type_Mesure", "Date_Mesure", "Valeur"})

	conf := getConfig()
	client := connect(conf.Broker.Url+":"+conf.Broker.Port, "clientcsv")
	var wg sync.WaitGroup
	wg.Add(1)
	client.Subscribe("/capteurs/#", 0, func(client mqtt.Client, msg mqtt.Message) {
		var c entities.CaptorValue
		json.Unmarshal([]byte(msg.Payload()), &c)
		var tab = strings.Split(msg.Topic(), "/")
		//csvLine := "\"" + tab[2] + "\",\"" + tab[3] + "\",\"" + c.Timestamp.String() + "\"," + fmt.Sprintf("%f", c.Value) + ";"
		//fmt.Println("Data received : ", csvLine)
		writer.Write([]string{tab[2], tab[3], c.Timestamp.String(), fmt.Sprintf("%f", c.Value)})
		writer.Flush()
	})

	wg.Wait()
}
func handleError(err error) {
	log.Fatal(err)
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
