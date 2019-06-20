package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	ServerPort                int
	ServerTls                 bool
	ServerTimeout             int
	ServerDialTimeout         int
	ServerTlsHandshakeTimeout int
	ServerFileCrt             string
	ServerFileKey             string
	AuthApiEndpoint           string
	AuthApiType               string
	FilesRootDirectory        string
}

func Init() *config {
	var location = "config.yaml"
	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
		panic(fmt.Sprintf("Load config ERROR  #%v ", err))
	}
	cfg := &config{}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal ERROR: %v", err))
	}
	return cfg
}
