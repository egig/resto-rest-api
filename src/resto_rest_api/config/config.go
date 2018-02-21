package config

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	DBDriver, DBDSN string
}

func GetValue() Config {
	file, _ := os.Open("../../config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	return config;
}