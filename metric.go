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
			publicHolidayActive *prometheus.GaugeVec
			publicHolidayDate *prometheus.GaugeVec
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
	c.restClient.SetHeader("User-Agent", "public-holiday-exporter/"+gitTag)
	c.restClient.SetHostURL("https://date.nager.at/api/v2/publicholidays/")
	c.restClient.SetHeader("Accept", "application/json")

	c.prometheus.publicHolidayActive = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "publicholiday_active",
			Help: "Active public holidays",
		},
		[]string{
			"countryCode",
			"county",
			"date",
			"timezone",
			"fixed",
			"global",
			"launchYear",
			"localName",
			"name",
			"type",
		},
	)

	c.prometheus.publicHolidayDate = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "publicholiday_date",
			Help: "Date (unix timestamp) of public holidays",
		},
		[]string{
			"countryCode",
			"county",
			"date",
			"timezone",
			"fixed",
			"global",
			"launchYear",
			"localName",
			"name",
			"type",
		},
	)

	prometheus.MustRegister(c.prometheus.publicHolidayActive)
	prometheus.MustRegister(c.prometheus.publicHolidayDate)
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
			c.prometheus.publicHolidayActive.Reset()
			c.prometheus.publicHolidayDate.Reset()

			currentTime := time.Now()
			nextYear := time.Date(currentTime.Year() + 1, 1, 1, 0, 0, 0, 0, currentTime.Location())
			daysToNextYear := int64(nextYear.Sub(currentTime).Hours() / 24)

			for _, country := range c.Countries {
				location, err := time.LoadLocation(country.Timezone)
				if err != nil {
					panic(err)
				}

				locationTime := currentTime.In(location)

				year := locationTime.Year()
				todayDate := locationTime.Format("2006-01-02")

				c.ensurePublicHolidays(country.CountryCode, year)
				c.setCountryMetrics(country.CountryCode, year, todayDate, location)

				if daysToNextYear <= opts.ExporterDaysToFetchNewYear {
					c.ensurePublicHolidays(country.CountryCode, year + 1 )
					c.setCountryMetrics(country.CountryCode, year + 1 , todayDate, location)
				}
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

func (c *MetricCollector) setCountryMetrics(country string, year int, todayDate string, location *time.Location) {
	if _, ok := c.PublicHolidays[country]; !ok {
		return
	}

	if _, ok := c.PublicHolidays[country][year]; !ok {
		return
	}

	// get metrics and create them
	for _, holiday := range c.PublicHolidays[country][year] {

		// check if holiday is today (in location)
		holidayIsActive := holiday.Date == todayDate

		// parse timestamp of start of holiday
		holidayTimestamp := float64(0)
		holidayTime, err := time.ParseInLocation("2006-01-02", holiday.Date, location)
		if err == nil {
			holidayTimestamp = float64(holidayTime.Unix())
		}

		if holiday.Counties != nil && len(*holiday.Counties) > 0 {
			for _, county := range *holiday.Counties {
				if holidayIsActive {
					c.prometheus.publicHolidayActive.With(
						prometheus.Labels{
							"countryCode": holiday.CountryCode,
							"county":      county,
							"date":        holiday.Date,
							"timezone":    location.String(),
							"fixed":       boolToString(holiday.Fixed),
							"global":      boolToString(holiday.Global),
							"launchYear":  int64ToString(holiday.LaunchYear),
							"localName":   holiday.LocalName,
							"name":        holiday.Name,
							"type":        holiday.Type,
						}).Set(1)
				}

				c.prometheus.publicHolidayDate.With(
					prometheus.Labels{
						"countryCode": holiday.CountryCode,
						"county":      county,
						"date":        holiday.Date,
						"timezone":    location.String(),
						"fixed":       boolToString(holiday.Fixed),
						"global":      boolToString(holiday.Global),
						"launchYear":  int64ToString(holiday.LaunchYear),
						"localName":   holiday.LocalName,
						"name":        holiday.Name,
						"type":        holiday.Type,
					}).Set(holidayTimestamp)
			}
		} else {
			if holidayIsActive {
				c.prometheus.publicHolidayActive.With(
					prometheus.Labels{
						"countryCode": holiday.CountryCode,
						"county":      "",
						"date":        holiday.Date,
						"timezone":    location.String(),
						"fixed":       boolToString(holiday.Fixed),
						"global":      boolToString(holiday.Global),
						"launchYear":  int64ToString(holiday.LaunchYear),
						"localName":   holiday.LocalName,
						"name":        holiday.Name,
						"type":        holiday.Type,
					}).Set(1)
			}

			c.prometheus.publicHolidayDate.With(
				prometheus.Labels{
					"countryCode": holiday.CountryCode,
					"county":      "",
					"date":        holiday.Date,
					"timezone":    location.String(),
					"fixed":       boolToString(holiday.Fixed),
					"global":      boolToString(holiday.Global),
					"launchYear":  int64ToString(holiday.LaunchYear),
					"localName":   holiday.LocalName,
					"name":        holiday.Name,
					"type":        holiday.Type,
				}).Set(holidayTimestamp)
		}
	}
}
