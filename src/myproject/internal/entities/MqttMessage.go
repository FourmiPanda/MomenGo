package entities

import (
	"fmt"
	"strconv"
	"strings"
)

/*
## Structure des messages MQTT
```json
{
 "idCaptor":  1,
 "idAirport": "BIA",
 "measure":   "Temperature",
 "values":     [
       {
        "value": 27.8,
        "time":  "2007-03-01T13:00:00Z"
       }
     ]
}
Une seule value dans values pour les messages MQTT
```
*/
type MqttMessage struct {
	Captor *Captor
}

func CreateAMqttMessage(captor *Captor) *MqttMessage{
	c := MqttMessage{
		Captor: captor,
	}
	return &c
}
func CreateAMqttMessageFromByte(json []byte) *MqttMessage{
	c := MqttMessage{
		Captor: CreateACaptor(json),
	}
	return &c
}

func (r *MqttMessage) MqttMessageToString() string {
	return  r.Captor.CaptorToString()
}

func (m *MqttMessage) MqttMessageToJson() []byte {
	return m.Captor.CaptorToJson()
}
func (m* MqttMessage) createAMqttMessageFromTopic(topic string) *MqttMessage {
	return &MqttMessage{m.createACaptorFromATopic(topic)}
}
func (m* MqttMessage) createACaptorFromATopic(topic string) *Captor {
	t := strings.Split(topic,"/")
	fmt.Println("DEBUG :",t)
	idAirport := t[2]
	IdCaptor, _ := strconv.ParseInt(t[4], 10, 64)
	measure := t[3]
	emptyValue := []*CaptorValue{}
	m.Captor = &Captor{IdAirport:idAirport,IdCaptor:int(IdCaptor),Measure:measure, Values:emptyValue}
	fmt.Println("DEBUG :",m.MqttMessageToString())
	return m.Captor
}