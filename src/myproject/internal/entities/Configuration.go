package entities

type Configuration struct {
	Broker   Broker
	Capteurs []Capteur
	Redis    RedisDB
}

type Broker struct {
	Url  string
	Port string
}

type RedisDB struct {
	Network string
	Address string
}
