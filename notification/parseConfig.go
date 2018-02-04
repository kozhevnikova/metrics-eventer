package main

import (
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Name     string
		Host     string
	}
	Metrics struct {
		Critical int
	}
	Redis struct {
		Address  string
		Password string
		Database int
	}
	Mail struct {
		AddressFrom string
		Password    string
		ServerName  string
		Port        int
	}
}

func parseConfig() (Config, error) {
	var config Config
	f, err := os.Open("config.toml")
	if err != nil {
		return config, err
	}
	defer f.Close()
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}
