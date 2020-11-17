package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type (
	Opts struct {
		// logger
		Logger struct {
			Debug   bool `           long:"debug"        env:"DEBUG"    description:"debug mode"`
			Verbose bool `short:"v"  long:"verbose"      env:"VERBOSE"  description:"verbose mode"`
			LogJson bool `           long:"log.json"     env:"LOG_JSON" description:"Switch log output to json format"`
		}

		// general options
		ServerBind string `long:"bind"                env:"SERVER_BIND"   description:"Server address"                default:":8080"`

		ConfigPath string `long:"config" short:"c"  env:"CONFIG"   description:"Config path" required:"true"`

		// exporter settings
		ExporterDaysToFetchNewYear int64  `env:"DAYS_TO_NEXT_YEAR"   description:"days to next year to fetch also next years data" default:"30"`
	}
)

func (o *Opts) GetJson() []byte {
	jsonBytes, err := json.Marshal(o)
	if err != nil {
		log.Panic(err)
	}
	return jsonBytes
}
