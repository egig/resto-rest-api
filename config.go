package main

import (
	"encoding/json"
	"os"
	"fmt"
)

type Config struct {
	MySQLHost string
	MySQLPort string
	MySQLUser string
	MySQLPassword string
}

file, _ := os.Open("config.json")
decoder := json.NewDecoder(file)
config := Config{}
err := decoder.Decode(&config)
if err != nil {
	fmt.Println("error:", err)
}