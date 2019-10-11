package entities

type Capteur struct {
	Id   int
	IATA string
	Type int
}

func (c *Capteur) getValue() int {
	return 99
}
