package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port                int
		Timeout             int
		DialTimeout         int
		TLSEnabled          bool
		TLSHandshakeTimeout int
		TLSFileCrt          string
		TLSFileKey          string
	}
	Handler struct {
		RootDirectory    string
		AuthType         string
		JwtLoginEndpoint string
	}
}

func loadConfig(location string) (*Config, error) {

	cfg := &Config{}
	_, err := os.Stat(location)

	if os.IsNotExist(err) {
		return cfg, fmt.Errorf("Config file not found: %s", location)
	}

	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		return cfg, fmt.Errorf("Load config error %v ", err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return cfg, fmt.Errorf("Unmarshal: %v", err)
	}

	return cfg, err
}
