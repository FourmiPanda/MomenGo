package entities

type Capteur struct {
	Id   int
	IATA string
	Type string
}

func (c *Capteur) getValue() int {
	return 99
}
