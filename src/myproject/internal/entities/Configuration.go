package entities

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

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

func GetConfig() *Configuration {
	configPath, _ := filepath.Abs("src/config/config.json")
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatal(err)
	}
	return &configuration
}