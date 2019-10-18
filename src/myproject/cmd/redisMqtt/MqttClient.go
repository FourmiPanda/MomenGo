package redisMqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"myproject/internal/entities"
	"time"
)

type MqttClient struct {
	Configuration	*entities.Configuration
	MqttClientOpts 	mqtt.ClientOptions
	MqttClient 		mqtt.Client
	Token 			mqtt.Token
}
func (m *MqttClient) SetOptions (broker *entities.Broker) *mqtt.ClientOptions {
	m.MqttClientOpts = *mqtt.NewClientOptions()
	m.MqttClientOpts.AddBroker(m.Configuration.Broker.Url+":"+ m.Configuration.Broker.Port)
	m.MqttClientOpts.SetKeepAlive(2 * time.Second)
	m.MqttClientOpts.SetPingTimeout(1 * time.Second)
	return &m.MqttClientOpts
}
func CreateAMqttClientFromABroker(config *entities.Configuration) *MqttClient {
	m := MqttClient{Configuration: config}
	m.SetOptions(&config.Broker)
	m.MqttClient = mqtt.NewClient(&m.MqttClientOpts)
	return &m
}

func (m *MqttClient) ConnectClient() *MqttClient{
	m.Token = m.MqttClient.Connect()
	m.Token.Wait()
	return m
}

func (m *MqttClient) PublishAMessage(topic string, message string) *MqttClient{
	// Publishing a message //
	if !m.MqttClient.IsConnected() {
		m.ConnectClient()
	}
	m.MqttClient.Publish(topic,0,false, message).Wait()
	return m
}
func (m *MqttClient) SubscribeAToATopic(topic string) *MqttClient{
	// Subcribe to a topic //
	if !m.MqttClient.IsConnected() {
		m.ConnectClient()
	}
	m.Token = m.MqttClient.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		println("DEBUG :",msg.Topic())
		println("DEBUG :",string(msg.Payload()))
		// Create a MqttMessage from the topic and payload received
		mMqtt := entities.CreateAMqttMessageFromPublish(msg.Topic(),msg.Payload())
		println("DEBUG :",mMqtt.MqttMessageToString())
		println("DEBUG :",string(mMqtt.MqttMessageToJson()))
		// Create a RedisEntry from the MqttMessage
		re := entities.CreateARedisEntryFromMqtt(mMqtt)
		println("DEBUG :",re.RedisEntryToString())
		// Create a RedisClient from the configuration
		rc := CreateARedisClientFromConfig(m.Configuration.Redis)
		// Add the RedisEntru to the database
		rc.AddCaptorEntryToDB(re)
	})
	m.Token.Wait()
	if m.Token.Error() != nil {
		log.Fatal(m.Token.Error())
	}
	return m
}