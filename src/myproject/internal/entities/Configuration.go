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

var Config Configuration = Configuration{
	Broker:   Broker{
		Url: "tcp://localhost",
		Port: "1883"},
	Capteurs: nil,
	Redis:    RedisDB{
		Network: "tcp",
		Address: "localhost:6379"},
}
