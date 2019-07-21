package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Server struct {
		Host string "json:'host'"
		Port string "json:'port'"
	} "json:'server'"
	Options struct {
		Prefix string "json:'prefix'"
	} "json:'options'"
}

func ReadConfig(path string) (*Config, error) {
	ba, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var cfg Config

	if err := json.Unmarshal(ba, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
