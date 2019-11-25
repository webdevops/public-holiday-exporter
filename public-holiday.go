package main

type (
	PublicHoliday struct {
		Counties    *[]string `json:"counties"`
		CountryCode string    `json:"countryCode"`
		Date        string    `json:"date"`
		Fixed       bool      `json:"fixed"`
		Global      bool      `json:"global"`
		LaunchYear  int64     `json:"launchYear"`
		LocalName   string    `json:"localName"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
	}
)
