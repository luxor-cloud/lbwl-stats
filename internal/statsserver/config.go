package statsserver

import (
	"encoding/json"
	"io/ioutil"
)

var DefaultConfig = Config{
	Address: "127.0.0.1:1337",
	FlashDBConnection: DatabaseConfig{
		Host:     "localhost:3306",
		Username: "root",
		Password: "secret",
		Database: "test",
	},
	UseTLS:  false,
	TLSKey:  "",
	TLSCert: "",
}

type Config struct {
	Address           string         `json:"address"`
	FlashDBConnection DatabaseConfig `json:"flash_db_connection"`
	UseTLS            bool           `json:"use_tls,omitempty"`
	TLSKey            string         `json:"tls_key_path,omitempty"`
	TLSCert           string         `json:"tls_cert_path,omitempty"`
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
