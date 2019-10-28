package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	conf := getConfig()
	client := connect(conf.Broker.Url+":"+conf.Broker.Port, "capteurs")
	var wg sync.WaitGroup

	for _, v := range conf.Capteurs {
		wg.add(1)
		fmt.Println("[CAPTEUR-" + strconv.Itoa(v.Id) + "] : Publishing data for airport " + v.IATA)
		client.Subscribe("/capteurs/"+v.IATA+"/"+v.Type, byte(v.QoS), func(client mqtt.Client, msg mqtt.Message) {
			json.Unmarshal([]byte(msg.Payload()), &data)
			fmt.Println("Data received : ", data)
			csvwritter.AddDataToCsv(data, boot.CsvDataPath)
		})
		wg.wait()
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
