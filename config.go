package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Server serverConfig
	// DBConn databaseConnection
}

type serverConfig struct {
	Port       string `json:"port"`
	Name       string `json:"name"`
	DBUserName string `json:"dbUsername"`
	DBPassword string `json:"dbPassword"`
	DBHostname string `json:"dbHostname"`
	DBPort     string `json:"dbPort"`
	DBDatabase string `json:"dbDatabase"`
}

// type databaseConnection struct {
// 	DBUserName string `json:"DBUsername"`
// 	DBPassword string `json:"DBPassword"`
// 	DBHostname string `json:"DBHostname"`
// 	DBPort     string `json:"DBPort"`
// 	DBDatabase string `json:"DBDatabase"`
// }

func LoadConfig(path string) Configuration {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}
	return config
}
