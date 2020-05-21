package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

var (
	argparser *flags.Parser

	Logger    *DaemonLogger
	Verbose   bool

	// Git version information
	gitCommit = "<unknown>"
	gitTag    = "<unknown>"
)

var opts struct {
	// general settings
	Verbose    []bool `long:"verbose" short:"v"  env:"VERBOSE"  description:"Verbose mode"`
	ConfigPath string `long:"config" short:"c"  env:"CONFIG"   description:"Config path" required:"true"`

	// exporter settings
	ExporterDaysToFetchNewYear int64  `env:"DAYS_TO_NEXT_YEAR"   description:"days to next year to fetch also next years data" default:"30"`

	// server settings
	ServerBind string `long:"bind"  env:"SERVER_BIND"  description:"Server address"  default:":8000"`
}

func main() {
	initArgparser()

	// set verbosity
	Verbose = len(opts.Verbose) >= 1

	Logger = NewLogger(log.Lshortfile, Verbose)
	defer Logger.Close()

	Logger.Infof("Starting public holiday exporter v%s (%s)", gitTag, gitCommit)

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
			fmt.Println()
			argparser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
	}
}
