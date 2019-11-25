package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

const (
	Version     = "0.1.0"
)

var (
	argparser *flags.Parser
	args      []string

	Logger    *DaemonLogger
	Verbose   bool
)

var opts struct {
	// general settings
	Verbose    []bool `long:"verbose" short:"v"  env:"VERBOSE"  description:"Verbose mode"`
	ConfigPath string `long:"config" short:"c"  env:"CONFIG"   description:"Config path" required:"true"`

	// server settings
	ServerBind string `long:"bind"  env:"SERVER_BIND"  description:"Server address"  default:":8000"`
}

func main() {
	initArgparser()

	// set verbosity
	Verbose = len(opts.Verbose) >= 1

	Logger = NewLogger(log.Lshortfile, Verbose)
	defer Logger.Close()

	Logger.Infof("Starting public-holiday exporter version %v", Version)

	Logger.Infof("parsing configuration file %v", opts.ConfigPath)

	config := NewAppConfig(opts.ConfigPath)

	collector := NewMetricCollector()
	for _, line := range config.Countries {
		Logger.Infof("adding country %v with timezone %v", line.Country, line.Timezone)
		collector.AddCountry(line.Country,line.Timezone)
	}
	collector.Run()

	Logger.Infof("starting service at %v", opts.ServerBind)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(opts.ServerBind, nil))
}

// init argparser and parse/validate arguments
func initArgparser() {
	argparser = flags.NewParser(&opts, flags.Default)
	_, err := argparser.Parse()

	// check if there is an parse error
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			fmt.Println(err)
			fmt.Println()
			argparser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}
}
