package entities
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
	Captor Captor
}

func CreateAMqttMessage(jsonEntry []byte) *MqttMessage{
	c := MqttMessage{
		Captor: *CreateACaptor(jsonEntry),
	}
	return &c
}

func (r *MqttMessage) MqttMessageToString() string {
	return  r.Captor.CaptorToString()
}

func (m *MqttMessage) MqttMessageToJson() []byte {
	return m.Captor.CaptorToJson()
}