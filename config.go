package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	Config struct {
		Countries []ConfigCountry `yaml:"countries"`
		Proxy     *string         `yaml:"proxy,omitempty"`
		Endpoint  *string         `yaml:"endpoint,omitempty"`
	}

	ConfigCountry struct {
		Country  string `yaml:"country"`
		Timezone string `yaml:"timezone"`
	}
)

func NewAppConfig(path string) (config Config) {
	var fileContent []byte

	config = Config{}

	log.Infof("reading configuration from file %v", path)
	if data, err := ioutil.ReadFile(path); err == nil {
		fileContent = data
	} else {
		panic(err)
	}

	log.Info("parsing configuration")
	if err := yaml.Unmarshal(fileContent, &config); err != nil {
		panic(err)
	}

	return
}
