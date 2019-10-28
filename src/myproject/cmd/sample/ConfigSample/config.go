package main

import (
	"encoding/json"
	"fmt"
	"myproject/internal/entities"
	"os"
	"path/filepath"
)

func main() {
	configPath, _ := filepath.Abs("src/config/config.json")
	file, err := os.Open(configPath)
	if err != nil {
		handleError(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := entities.Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		handleError(err)
	}
	fmt.Println(configuration)
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}
