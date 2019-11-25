Public-holiday Exporter
==================

[![license](https://img.shields.io/github/license/webdevops/public-holiday-exporter.svg)](https://github.com/webdevops/public-holiday-exporter/blob/master/LICENSE)
[![Docker](https://img.shields.io/badge/docker-webdevops%2Fpublic--holiday--exporter-blue.svg?longCache=true&style=flat&logo=docker)](https://hub.docker.com/r/webdevops/public-holiday-exporter/)
[![Docker Build Status](https://img.shields.io/docker/build/webdevops/public-holiday-exporter.svg)](https://hub.docker.com/r/webdevops/public-holiday-exporter/)

Prometheus exporter for public-holiday informations (countries with timezones...)

Example metrics
---------------

```
publicholiday_active{countryCode="DE",county="",date="2019-01-01",fixed="true",global="true",launchYear="1967",localName="Neujahr",name="New Year's Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-04-19",fixed="false",global="true",launchYear="0",localName="Karfreitag",name="Good Friday",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-04-22",fixed="false",global="true",launchYear="1642",localName="Ostermontag",name="Easter Monday",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-05-01",fixed="true",global="true",launchYear="0",localName="Tag der Arbeit",name="Labour Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-05-30",fixed="false",global="true",launchYear="0",localName="Christi Himmelfahrt",name="Ascension Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-06-10",fixed="false",global="true",launchYear="0",localName="Pfingstmontag",name="Whit Monday",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-10-03",fixed="true",global="true",launchYear="0",localName="Tag der Deutschen Einheit",name="German Unity Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-12-25",fixed="true",global="true",launchYear="0",localName="Erster Weihnachtstag",name="Christmas Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="",date="2019-12-26",fixed="true",global="true",launchYear="0",localName="Zweiter Weihnachtstag",name="St. Stephen's Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BB",date="2019-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BE",date="2019-03-08",fixed="true",global="false",launchYear="2019",localName="Internationaler Frauentag",name="International Women's Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BW",date="2019-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BW",date="2019-06-20",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BW",date="2019-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BY",date="2019-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BY",date="2019-06-20",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",type="Public"} 0
publicholiday_active{countryCode="DE",county="DE-BY",date="2019-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",type="Public"} 0
```
