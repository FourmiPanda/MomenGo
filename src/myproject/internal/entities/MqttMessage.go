package entities

import (
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
func CreateAMqttMessageFromPublish(topic string, payload []byte) (*MqttMessage, error){
	m := MqttMessage{}
	m.createAMqttMessageFromTopic(topic)
	_, err := m.addValuesFromPayload(payload)
	return &m, err
}
func (m *MqttMessage) addValuesFromPayload(payload []byte) (*MqttMessage, error) {
	//fmt.Println("DEBUG :", "addValuesFromPayload")
	_, err := m.Captor.AddValuesFromJson(payload)
	//fmt.Println("DEBUG : MqttMessage ",m.MqttMessageToString())
	return  m, err
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
	return m.Captor.CaptorToSliceByte()
}
func (m* MqttMessage) createAMqttMessageFromTopic(topic string) *MqttMessage {
	return &MqttMessage{m.createACaptorFromATopic(topic)}
}
func (m* MqttMessage) createACaptorFromATopic(topic string) *Captor {
	t := strings.Split(topic,"/")
	//fmt.Println("DEBUG : topic ",t)
	idAirport := t[2]
	IdCaptor, _ := strconv.ParseInt(t[4], 10, 64)
	measure := t[3]
	emptyValue := []*CaptorValue{}
	m.Captor = &Captor{IdAirport:idAirport,IdCaptor:int(IdCaptor),Measure:measure, Values:emptyValue}
	//fmt.Println("DEBUG : MqttMessage ",m.MqttMessageToString())
	return m.Captor
}
func (m *MqttMessage) MqttMessageToSliceString() [][]string  {
	/*
		[][]string{
			[]string{
				"AAI",
				"TEMP",
				"1",
				"23.8",
				"2007-03-01T13:00:00Z"}}
	 */
	var res [][]string
	for i := 0; i < len(m.Captor.Values); i++{
		res[i] = []string{
			m.Captor.GetIdAirportToString(),
			m.Captor.GetMeasureToString(),
			m.Captor.GetIdCaptorToString(),
			m.Captor.Values[i].GetValueToString(),
			m.Captor.Values[i].GetTimestampToString()}
	}
	return res
}