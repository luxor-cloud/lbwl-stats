package statsserver

import (
	"encoding/json"
	"io/ioutil"
)

var DefaultConfig = Config{
	Address: "127.0.0.1",
	FlashDB: map[string]string{ "type": "SQL", "address": "root:secret@/test"},
	UseTLS:  false,
	TLSKey:  "",
	TLSCert: "",
}

type Config struct {
	Address string            `json:"address"`
	FlashDB map[string]string `json:"flash_db"`
	UseTLS  bool              `json:"use_tls,omitempty"`
	TLSKey  string            `json:"tls_key_path,omitempty"`
	TLSCert string            `json:"tls_cert_path,omitempty"`
}

type SQLConfig struct {
	Address string `json:"address"`
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
