package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Countries []ConfigCountry `yaml:"countries"`
	}

	ConfigCountry struct {
		Country  string `yaml:"country"`
		Timezone string `yaml:"timezone"`
	}
)

func NewAppConfig(path string) (config Config) {
	var filecontent []byte

	config = Config{}

	log.Infof("reading configuration from file %v", path)
	/* #nosec G304 */
	if data, err := os.ReadFile(path); err == nil {
		filecontent = data
	} else {
		panic(err)
	}

	log.Info("parsing configuration")
	if err := yaml.Unmarshal(filecontent, &config); err != nil {
		panic(err)
	}

	return
}
