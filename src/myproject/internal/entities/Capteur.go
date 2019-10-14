package entities

import (
	"math/rand"
)

type Capteur struct {
	Id   int
	IATA string
	Type string
}

func (c *Capteur) GetValue() int {
	return rand.Intn(100)
}
