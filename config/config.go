package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Seed        string `json:"seed"`
	NodeAddress string `json:"nodeAddress"`
}

func LoadConfig(path string) (*Config, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(b, &config)
	return &config, err
}
