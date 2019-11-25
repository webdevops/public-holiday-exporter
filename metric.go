package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type (
	MetricCollector struct {
		Countries []MetricCollectorCountry

		PublicHolidays map[string]map[int][]PublicHoliday

		restClient *resty.Client

		interval time.Duration

		prometheus struct {
			publicHoliday *prometheus.GaugeVec
		}
	}

	MetricCollectorCountry struct {
		CountryCode string
		Timezone string
	}
)

func NewMetricCollector() *MetricCollector {
	collector := &MetricCollector{}
	collector.Setup()
	return collector
}

func (c *MetricCollector) Setup() {
	c.PublicHolidays = map[string]map[int][]PublicHoliday{}

	c.interval, _ = time.ParseDuration("30m")

	c.restClient = resty.New()
	c.restClient.SetHeader("User-Agent", "public-holiday-exporter/"+Version)
	c.restClient.SetHostURL("https://date.nager.at/api/v2/publicholidays/")
	c.restClient.SetHeader("Accept", "application/json")


	c.prometheus.publicHoliday = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "publicholiday_active",
			Help: "Date and country of the Public Holiday",
		},
		[]string{
			"countryCode",
			"county",
			"date",
			"fixed",
			"global",
			"launchYear",
			"localName",
			"name",
			"type",
		},
	)

	prometheus.MustRegister(c.prometheus.publicHoliday)
}

func (c *MetricCollector) AddCountry(country, timezone string) {
	c.Countries = append(c.Countries, MetricCollectorCountry{
		CountryCode: country,
		Timezone: timezone,
	})

	if _, ok := c.PublicHolidays[country]; !ok {
		c.PublicHolidays[country] = map[int][]PublicHoliday{}
	}
}


func (c *MetricCollector) Run() {
	go func() {
		for {
			fmt.Println("Run...")
			c.prometheus.publicHoliday.Reset()

			currentTime := time.Now()

			for _, country := range c.Countries {
				location, err := time.LoadLocation(country.Timezone)
				if err != nil {
					panic(err)
				}

				locationTime := currentTime.In(location)

				year := locationTime.Year()
				todayDate := locationTime.Format("2006-01-02")

				c.ensurePublicHolidays(country.CountryCode, year)
				c.setCountryMetrics(country.CountryCode, year, todayDate)
			}

			time.Sleep(c.interval)
		}
	}()
}

func (c *MetricCollector) ensurePublicHolidays(country string, year int) {
	// only fetch if not already stored
	if _, ok := c.PublicHolidays[country][year]; !ok {
		url := fmt.Sprintf("%v/%v", year, country)
		resp, err := c.restClient.R().Get(url)
		if err != nil {
			panic(err)
		}

		publicHolidayList := []PublicHoliday{}
		if err := json.Unmarshal(resp.Body(), &publicHolidayList); err != nil {
			panic(err)
		}

		c.PublicHolidays[country][year] = publicHolidayList
	}
}

func (c *MetricCollector) setCountryMetrics(country string, year int, todayDate string) {
	if _, ok := c.PublicHolidays[country]; !ok {
		return
	}

	if _, ok := c.PublicHolidays[country][year]; !ok {
		return
	}

	// get metrics and create them
	for _, holiday := range c.PublicHolidays[country][year] {

		metricValue := float64(0)
		if holiday.Date == todayDate {
			metricValue = 1
		}

		if holiday.Counties != nil && len(*holiday.Counties) > 0 {
			for _, county := range *holiday.Counties {
				c.prometheus.publicHoliday.With(
					prometheus.Labels{
						"countryCode": holiday.CountryCode,
						"county":      county,
						"date":        holiday.Date,
						"fixed":       boolToString(holiday.Fixed),
						"global":      boolToString(holiday.Global),
						"launchYear":  int64ToString(holiday.LaunchYear),
						"localName":   holiday.LocalName,
						"name":        holiday.Name,
						"type":        holiday.Type,
					}).Set(metricValue)
			}
		} else {
			c.prometheus.publicHoliday.With(
				prometheus.Labels{
					"countryCode": holiday.CountryCode,
					"county":      "",
					"date":        holiday.Date,
					"fixed":       boolToString(holiday.Fixed),
					"global":      boolToString(holiday.Global),
					"launchYear":  int64ToString(holiday.LaunchYear),
					"localName":   holiday.LocalName,
					"name":        holiday.Name,
					"type":        holiday.Type,
				}).Set(metricValue)
		}
	}
}
