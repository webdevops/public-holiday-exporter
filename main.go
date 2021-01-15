package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/webdevops/public-holiday-exporter/config"
	"net/http"
	"os"
	"path"
	"runtime"
	"strings"
)

const (
	Author = "webdevops.io"
)

var (
	argparser *flags.Parser

	// Git version information
	gitCommit = "<unknown>"
	gitTag    = "<unknown>"
)

var opts config.Opts

func main() {
	initArgparser()

	log.Infof("starting public-holiday-exporter v%s (%s; %s; by %v)", gitTag, gitCommit, runtime.Version(), Author)
	log.Info(string(opts.GetJson()))

	config := NewAppConfig(opts.ConfigPath)

	collector := NewMetricCollector()
	for _, line := range config.Countries {
		log.Infof("adding country %v with timezone %v", line.Country, line.Timezone)
		collector.AddCountry(line.Country, line.Timezone)
	}
	
	if opts.Cache.Path != "" {
		if _, err := os.Stat(opts.Cache.Path); err == nil {
			collector.LoadFromCache(opts.Cache.Path)
		}
	}

	// check if preload only
	// load data, save to cache and exit
	if opts.Preload {
		collector.Preload()
		if opts.Cache.Path != "" {
			collector.SaveToCache(opts.Cache.Path)
		}

		log.Info("preloading finished, exit")
		os.Exit(0)
	}

	collector.Run()

	log.Infof("listening at %v", opts.ServerBind)
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

	// verbose level
	if opts.Logger.Verbose {
		log.SetLevel(log.DebugLevel)
	}

	// debug level
	if opts.Logger.Debug {
		log.SetReportCaller(true)
		log.SetLevel(log.TraceLevel)
		log.SetFormatter(&log.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			},
		})
	}

	// json log format
	if opts.Logger.LogJson {
		log.SetReportCaller(true)
		log.SetFormatter(&log.JSONFormatter{
			DisableTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				return funcName, fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
			},
		})
	}
}
