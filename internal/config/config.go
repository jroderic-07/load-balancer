package config

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Backends []string `json:'backends'`
}

func New(configFile string) *Config {
	return &Config{
		readConfig(configFile),
	}
}

func readConfig(configFile string) []string {
	jsonFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		print("ERROR")
	}

	var backend Config
	_ = json.Unmarshal(jsonFile, &backend)
	
	return backend.Backends

}

func (c *Config) GetBackends() []string {
	return c.Backends
}
