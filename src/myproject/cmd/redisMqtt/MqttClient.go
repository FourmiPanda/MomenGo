package redisMqtt

import (
	"myproject/internal/entities"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"sync"
)

type MqttClient struct {
	Broker 			entities.Broker
	MqttClientOpts 	mqtt.ClientOptions
	MqttClient 		mqtt.Client
	Token 			mqtt.Token
}
func (m *MqttClient) SetOptions (broker *entities.Broker) *mqtt.ClientOptions {
	m.MqttClientOpts = *mqtt.NewClientOptions()
	m.MqttClientOpts.AddBroker(m.Broker.Url+":"+ m.Broker.Port)
	return &m.MqttClientOpts
}
func CreateAMqttClientFromABroker(broker *entities.Broker) *MqttClient {
	m := MqttClient{Broker: *broker}
	m.SetOptions(broker)
	m.MqttClient = mqtt.NewClient(&m.MqttClientOpts)
	return &m
}

func (m *MqttClient) ConnectClient() *MqttClient{
	m.Token = m.MqttClient.Connect()
	m.Token.Wait()
	return m
}

func (m *MqttClient) PublishAMessage(message string) *MqttClient{
	// Publishing a message //
	if !m.MqttClient.IsConnected() {
		m.ConnectClient()
	}
	m.MqttClient.Publish("test",0,false, message).Wait()
	return m
}
func (m *MqttClient) SubscribeAToATopic(message string) *MqttClient{
	// Subcribe to a topic //
	if !m.MqttClient.IsConnected() {
		m.ConnectClient()
	}
	var wg sync.WaitGroup
	if token := m.MqttClient.Subscribe("TEST", 0, func(client mqtt.Client, msg mqtt.Message) {
		if string(msg.Payload()) != "mymessage" {
			log.Fatalf("want mymessage, got %s", msg.Payload())
		}
		wg.Done()
	}); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	return m
}