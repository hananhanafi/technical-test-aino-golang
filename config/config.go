package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DB_USERNAME string `json:"DB_USERNAME"`
	DB_PASSWORD string `json:"DB_PASSWORD"`
	DB_PORT     string `json:"DB_PORT"`
	DB_HOST     string `json:"DB_HOST"`
	DB_NAME     string `json:"DB_NAME"`
}

func LoadConfiguration() Config {
	var config Config
	configFile, err := os.Open("config/config.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}
