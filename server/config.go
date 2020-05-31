package server

import (
	"encoding/json"
	"io/ioutil"
)

var DefaultConfig = Config{
	Address: "127.0.0.1:1337",
	FlashDBConnection: DatabaseConfig{
		Host:     "localhost:5432",
		Username: "postgres",
		Password: "secret",
		Database: "postgres",
	},
}

type Config struct {
	Address           string         `json:"address"`
	FlashDBConnection DatabaseConfig `json:"flash_db"`
}

type DatabaseConfig struct {
	Host     string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func ReadConfig(path string) (Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
