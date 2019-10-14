package entities

import (
	"math/rand"
)

type Capteur struct {
	Id   int
	IATA string
	Type string
	QoS  int
}

func (c *Capteur) GetValue() int {
	return rand.Intn(100)
}
