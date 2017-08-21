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

// Server is a generic struct for a server architecture.  hopefully useful for both server and database, with the server type being helpful in defining its use.
type Server struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
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

// LoadServerConfig takes the path to a json cofig file for loading into the program.
func LoadServerConfig(path string) Server {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Config File Missing. ", err)
	}

	var config Server
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal("Config Parse Error: ", err)
	}
	return config
}
