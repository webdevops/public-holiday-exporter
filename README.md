Public-holiday Exporter
==================

[![license](https://img.shields.io/github/license/webdevops/public-holiday-exporter.svg)](https://github.com/webdevops/public-holiday-exporter/blob/master/LICENSE)
[![DockerHub](https://img.shields.io/badge/DockerHub-webdevops%2Fpublic--holiday--exporter-blue)](https://hub.docker.com/r/webdevops/public-holiday-exporter/)
[![Quay.io](https://img.shields.io/badge/Quay.io-webdevops%2Fpublic--holiday--exporter-blue)](https://quay.io/repository/webdevops/public-holiday-exporter)

Prometheus exporter for public-holiday information (countries with timezones...)

Configuration
-------------

| Environment variable              | DefaultValue                | Description                |
|-----------------------------------|-----------------------------|----------------------------|
| `CONFIG`                          | `empty`                     | Config path                |
| `SERVER_BIND`                     | `:8000`                     | Server address             |
| `DAYS_TO_NEXT_YEAR`               | `30`                        | days to next year to fetch also next years data  |



Configuration file
------------------

see [example.yaml](example.yaml)

Example metrics
---------------

```
# HELP publicholiday_date Date (unix timestamp) of public holidays
# TYPE publicholiday_date gauge
publicholiday_date{countryCode="DE",county="",date="2020-01-01",fixed="true",global="true",launchYear="1967",localName="Neujahr",name="New Year's Day",timezone="Europe/Berlin",type="Public"} 1.5778332e+09
publicholiday_date{countryCode="DE",county="",date="2020-04-10",fixed="false",global="true",launchYear="0",localName="Karfreitag",name="Good Friday",timezone="Europe/Berlin",type="Public"} 1.5864696e+09
publicholiday_date{countryCode="DE",county="",date="2020-04-13",fixed="false",global="true",launchYear="1642",localName="Ostermontag",name="Easter Monday",timezone="Europe/Berlin",type="Public"} 1.5867288e+09
publicholiday_date{countryCode="DE",county="",date="2020-05-01",fixed="true",global="true",launchYear="0",localName="Tag der Arbeit",name="Labour Day",timezone="Europe/Berlin",type="Public"} 1.588284e+09
publicholiday_date{countryCode="DE",county="",date="2020-05-21",fixed="false",global="true",launchYear="0",localName="Christi Himmelfahrt",name="Ascension Day",timezone="Europe/Berlin",type="Public"} 1.590012e+09
publicholiday_date{countryCode="DE",county="",date="2020-06-01",fixed="false",global="true",launchYear="0",localName="Pfingstmontag",name="Whit Monday",timezone="Europe/Berlin",type="Public"} 1.5909624e+09
publicholiday_date{countryCode="DE",county="",date="2020-10-03",fixed="true",global="true",launchYear="0",localName="Tag der Deutschen Einheit",name="German Unity Day",timezone="Europe/Berlin",type="Public"} 1.601676e+09
publicholiday_date{countryCode="DE",county="",date="2020-12-25",fixed="true",global="true",launchYear="0",localName="Erster Weihnachtstag",name="Christmas Day",timezone="Europe/Berlin",type="Public"} 1.6088508e+09
publicholiday_date{countryCode="DE",county="",date="2020-12-26",fixed="true",global="true",launchYear="0",localName="Zweiter Weihnachtstag",name="St. Stephen's Day",timezone="Europe/Berlin",type="Public"} 1.6089372e+09
publicholiday_date{countryCode="DE",county="",date="2021-01-01",fixed="true",global="true",launchYear="1967",localName="Neujahr",name="New Year's Day",timezone="Europe/Berlin",type="Public"} 1.6094556e+09
publicholiday_date{countryCode="DE",county="",date="2021-04-02",fixed="false",global="true",launchYear="0",localName="Karfreitag",name="Good Friday",timezone="Europe/Berlin",type="Public"} 1.6173144e+09
publicholiday_date{countryCode="DE",county="",date="2021-04-05",fixed="false",global="true",launchYear="1642",localName="Ostermontag",name="Easter Monday",timezone="Europe/Berlin",type="Public"} 1.6175736e+09
publicholiday_date{countryCode="DE",county="",date="2021-05-01",fixed="true",global="true",launchYear="0",localName="Tag der Arbeit",name="Labour Day",timezone="Europe/Berlin",type="Public"} 1.61982e+09
publicholiday_date{countryCode="DE",county="",date="2021-05-13",fixed="false",global="true",launchYear="0",localName="Christi Himmelfahrt",name="Ascension Day",timezone="Europe/Berlin",type="Public"} 1.6208568e+09
publicholiday_date{countryCode="DE",county="",date="2021-05-24",fixed="false",global="true",launchYear="0",localName="Pfingstmontag",name="Whit Monday",timezone="Europe/Berlin",type="Public"} 1.6218072e+09
publicholiday_date{countryCode="DE",county="",date="2021-10-03",fixed="true",global="true",launchYear="0",localName="Tag der Deutschen Einheit",name="German Unity Day",timezone="Europe/Berlin",type="Public"} 1.633212e+09
publicholiday_date{countryCode="DE",county="",date="2021-12-25",fixed="true",global="true",launchYear="0",localName="Erster Weihnachtstag",name="Christmas Day",timezone="Europe/Berlin",type="Public"} 1.6403868e+09
publicholiday_date{countryCode="DE",county="",date="2021-12-26",fixed="true",global="true",launchYear="0",localName="Zweiter Weihnachtstag",name="St. Stephen's Day",timezone="Europe/Berlin",type="Public"} 1.6404732e+09
publicholiday_date{countryCode="DE",county="DE-BB",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-BB",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-BE",date="2020-03-08",fixed="true",global="false",launchYear="2019",localName="Internationaler Frauentag",name="International Women's Day",timezone="Europe/Berlin",type="Public"} 1.583622e+09
publicholiday_date{countryCode="DE",county="DE-BE",date="2020-05-08",fixed="false",global="false",launchYear="0",localName="Tag der Befreiung",name="Liberation Day",timezone="Europe/Berlin",type="Public"} 1.5888888e+09
publicholiday_date{countryCode="DE",county="DE-BE",date="2021-03-08",fixed="true",global="false",launchYear="2019",localName="Internationaler Frauentag",name="International Women's Day",timezone="Europe/Berlin",type="Public"} 1.615158e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2020-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.5782652e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2020-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2021-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.6098876e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-BW",date="2021-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2020-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.5782652e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2020-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2021-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.6098876e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-BY",date="2021-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="DE",county="DE-HB",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-HB",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-HE",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-HE",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-HH",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-HH",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-MV",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-MV",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-NI",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-NI",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-NW",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-NW",date="2020-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="DE",county="DE-NW",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-NW",date="2021-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="DE",county="DE-RP",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-RP",date="2020-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="DE",county="DE-RP",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-RP",date="2021-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="DE",county="DE-SH",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-SH",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2020-06-11",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.5918264e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2020-08-15",fixed="true",global="false",launchYear="0",localName="Mariä Himmelfahrt",name="Assumption Day",timezone="Europe/Berlin",type="Public"} 1.5974424e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2020-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2021-06-03",fixed="false",global="false",launchYear="0",localName="Fronleichnam",name="Corpus Christi",timezone="Europe/Berlin",type="Public"} 1.6226712e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2021-08-15",fixed="true",global="false",launchYear="0",localName="Mariä Himmelfahrt",name="Assumption Day",timezone="Europe/Berlin",type="Public"} 1.6289784e+09
publicholiday_date{countryCode="DE",county="DE-SL",date="2021-11-01",fixed="true",global="false",launchYear="0",localName="Allerheiligen",name="All Saints' Day",timezone="Europe/Berlin",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="DE",county="DE-SN",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-SN",date="2020-11-18",fixed="false",global="false",launchYear="0",localName="Buß- und Bettag",name="Repentance and Prayer Day",timezone="Europe/Berlin",type="Public"} 1.605654e+09
publicholiday_date{countryCode="DE",county="DE-SN",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-SN",date="2021-11-17",fixed="false",global="false",launchYear="0",localName="Buß- und Bettag",name="Repentance and Prayer Day",timezone="Europe/Berlin",type="Public"} 1.6371036e+09
publicholiday_date{countryCode="DE",county="DE-ST",date="2020-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.5782652e+09
publicholiday_date{countryCode="DE",county="DE-ST",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-ST",date="2021-01-06",fixed="true",global="false",launchYear="1967",localName="Heilige Drei Könige",name="Epiphany",timezone="Europe/Berlin",type="Public"} 1.6098876e+09
publicholiday_date{countryCode="DE",county="DE-ST",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="DE",county="DE-TH",date="2020-09-20",fixed="true",global="false",launchYear="2019",localName="Weltkindertag",name="World Children's Day",timezone="Europe/Berlin",type="Public"} 1.6005528e+09
publicholiday_date{countryCode="DE",county="DE-TH",date="2020-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6040988e+09
publicholiday_date{countryCode="DE",county="DE-TH",date="2021-09-20",fixed="true",global="false",launchYear="2019",localName="Weltkindertag",name="World Children's Day",timezone="Europe/Berlin",type="Public"} 1.6320888e+09
publicholiday_date{countryCode="DE",county="DE-TH",date="2021-10-31",fixed="true",global="false",launchYear="0",localName="Reformationstag",name="Reformation Day",timezone="Europe/Berlin",type="Public"} 1.6356312e+09
publicholiday_date{countryCode="FR",county="",date="2020-01-01",fixed="true",global="true",launchYear="1967",localName="Jour de l'an",name="New Year's Day",timezone="Europe/Paris",type="Public"} 1.5778332e+09
publicholiday_date{countryCode="FR",county="",date="2020-04-13",fixed="false",global="true",launchYear="1642",localName="Lundi de Pâques",name="Easter Monday",timezone="Europe/Paris",type="Public"} 1.5867288e+09
publicholiday_date{countryCode="FR",county="",date="2020-05-01",fixed="true",global="true",launchYear="0",localName="Fête du premier mai",name="Labour Day",timezone="Europe/Paris",type="Public"} 1.588284e+09
publicholiday_date{countryCode="FR",county="",date="2020-05-08",fixed="true",global="true",launchYear="0",localName="Fête de la Victoire",name="Victory in Europe Day",timezone="Europe/Paris",type="Public"} 1.5888888e+09
publicholiday_date{countryCode="FR",county="",date="2020-05-21",fixed="false",global="true",launchYear="0",localName="Jour de l'Ascension",name="Ascension Day",timezone="Europe/Paris",type="Public"} 1.590012e+09
publicholiday_date{countryCode="FR",county="",date="2020-06-01",fixed="false",global="true",launchYear="0",localName="Lundi de Pentecôte",name="Whit Monday",timezone="Europe/Paris",type="Public"} 1.5909624e+09
publicholiday_date{countryCode="FR",county="",date="2020-07-14",fixed="true",global="true",launchYear="0",localName="Fête nationale",name="Bastille Day",timezone="Europe/Paris",type="Public"} 1.5946776e+09
publicholiday_date{countryCode="FR",county="",date="2020-08-15",fixed="true",global="true",launchYear="0",localName="L'Assomption de Marie",name="Assumption Day",timezone="Europe/Paris",type="Public"} 1.5974424e+09
publicholiday_date{countryCode="FR",county="",date="2020-11-01",fixed="true",global="true",launchYear="0",localName="La Toussaint",name="All Saints' Day",timezone="Europe/Paris",type="Public"} 1.6041852e+09
publicholiday_date{countryCode="FR",county="",date="2020-11-11",fixed="true",global="true",launchYear="0",localName="Armistice de 1918",name="Armistice Day",timezone="Europe/Paris",type="Public"} 1.6050492e+09
publicholiday_date{countryCode="FR",county="",date="2020-12-25",fixed="true",global="true",launchYear="0",localName="Noël",name="Christmas Day",timezone="Europe/Paris",type="Public"} 1.6088508e+09
publicholiday_date{countryCode="FR",county="",date="2021-01-01",fixed="true",global="true",launchYear="1967",localName="Jour de l'an",name="New Year's Day",timezone="Europe/Paris",type="Public"} 1.6094556e+09
publicholiday_date{countryCode="FR",county="",date="2021-04-05",fixed="false",global="true",launchYear="1642",localName="Lundi de Pâques",name="Easter Monday",timezone="Europe/Paris",type="Public"} 1.6175736e+09
publicholiday_date{countryCode="FR",county="",date="2021-05-01",fixed="true",global="true",launchYear="0",localName="Fête du premier mai",name="Labour Day",timezone="Europe/Paris",type="Public"} 1.61982e+09
publicholiday_date{countryCode="FR",county="",date="2021-05-08",fixed="true",global="true",launchYear="0",localName="Fête de la Victoire",name="Victory in Europe Day",timezone="Europe/Paris",type="Public"} 1.6204248e+09
publicholiday_date{countryCode="FR",county="",date="2021-05-13",fixed="false",global="true",launchYear="0",localName="Jour de l'Ascension",name="Ascension Day",timezone="Europe/Paris",type="Public"} 1.6208568e+09
publicholiday_date{countryCode="FR",county="",date="2021-05-24",fixed="false",global="true",launchYear="0",localName="Lundi de Pentecôte",name="Whit Monday",timezone="Europe/Paris",type="Public"} 1.6218072e+09
publicholiday_date{countryCode="FR",county="",date="2021-07-14",fixed="true",global="true",launchYear="0",localName="Fête nationale",name="Bastille Day",timezone="Europe/Paris",type="Public"} 1.6262136e+09
publicholiday_date{countryCode="FR",county="",date="2021-08-15",fixed="true",global="true",launchYear="0",localName="L'Assomption de Marie",name="Assumption Day",timezone="Europe/Paris",type="Public"} 1.6289784e+09
publicholiday_date{countryCode="FR",county="",date="2021-11-01",fixed="true",global="true",launchYear="0",localName="La Toussaint",name="All Saints' Day",timezone="Europe/Paris",type="Public"} 1.6357212e+09
publicholiday_date{countryCode="FR",county="",date="2021-11-11",fixed="true",global="true",launchYear="0",localName="Armistice de 1918",name="Armistice Day",timezone="Europe/Paris",type="Public"} 1.6365852e+09
publicholiday_date{countryCode="FR",county="",date="2021-12-25",fixed="true",global="true",launchYear="0",localName="Noël",name="Christmas Day",timezone="Europe/Paris",type="Public"} 1.6403868e+09
publicholiday_date{countryCode="FR",county="FR-57",date="2020-04-10",fixed="false",global="false",launchYear="0",localName="Vendredi saint",name="Good Friday",timezone="Europe/Paris",type="Public"} 1.5864696e+09
publicholiday_date{countryCode="FR",county="FR-57",date="2020-12-26",fixed="true",global="false",launchYear="0",localName="Saint-Étienne",name="St. Stephen's Day",timezone="Europe/Paris",type="Public"} 1.6089372e+09
publicholiday_date{countryCode="FR",county="FR-57",date="2021-04-02",fixed="false",global="false",launchYear="0",localName="Vendredi saint",name="Good Friday",timezone="Europe/Paris",type="Public"} 1.6173144e+09
publicholiday_date{countryCode="FR",county="FR-57",date="2021-12-26",fixed="true",global="false",launchYear="0",localName="Saint-Étienne",name="St. Stephen's Day",timezone="Europe/Paris",type="Public"} 1.6404732e+09
publicholiday_date{countryCode="FR",county="FR-A",date="2020-04-10",fixed="false",global="false",launchYear="0",localName="Vendredi saint",name="Good Friday",timezone="Europe/Paris",type="Public"} 1.5864696e+09
publicholiday_date{countryCode="FR",county="FR-A",date="2020-12-26",fixed="true",global="false",launchYear="0",localName="Saint-Étienne",name="St. Stephen's Day",timezone="Europe/Paris",type="Public"} 1.6089372e+09
publicholiday_date{countryCode="FR",county="FR-A",date="2021-04-02",fixed="false",global="false",launchYear="0",localName="Vendredi saint",name="Good Friday",timezone="Europe/Paris",type="Public"} 1.6173144e+09
publicholiday_date{countryCode="FR",county="FR-A",date="2021-12-26",fixed="true",global="false",launchYear="0",localName="Saint-Étienne",name="St. Stephen's Day",timezone="Europe/Paris",type="Public"} 1.6404732e+09
publicholiday_date{countryCode="FR",county="FR-BL",date="2020-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.5905304e+09
publicholiday_date{countryCode="FR",county="FR-BL",date="2021-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.6220664e+09
publicholiday_date{countryCode="FR",county="FR-GP",date="2020-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.5905304e+09
publicholiday_date{countryCode="FR",county="FR-GP",date="2021-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.6220664e+09
publicholiday_date{countryCode="FR",county="FR-MF",date="2020-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.5905304e+09
publicholiday_date{countryCode="FR",county="FR-MF",date="2021-05-27",fixed="true",global="false",launchYear="0",localName="Abolition of Slavery",name="Abolition de l'esclavage",timezone="Europe/Paris",type="Public"} 1.6220664e+09
publicholiday_date{countryCode="FR",county="FR-MQ",date="2020-05-22",fixed="true",global="false",launchYear="0",localName="Abolition de l'esclavage",name="Abolition of Slavery",timezone="Europe/Paris",type="Public"} 1.5900984e+09
publicholiday_date{countryCode="FR",county="FR-MQ",date="2021-05-22",fixed="true",global="false",launchYear="0",localName="Abolition de l'esclavage",name="Abolition of Slavery",timezone="Europe/Paris",type="Public"} 1.6216344e+09
publicholiday_date{countryCode="GB",county="",date="2020-01-01",fixed="false",global="true",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Europe/London",type="Public"} 1.5778368e+09
publicholiday_date{countryCode="GB",county="",date="2020-04-10",fixed="false",global="true",launchYear="0",localName="Good Friday",name="Good Friday",timezone="Europe/London",type="Public"} 1.5864732e+09
publicholiday_date{countryCode="GB",county="",date="2020-04-13",fixed="false",global="true",launchYear="0",localName="Easter Monday",name="Easter Monday",timezone="Europe/London",type="Public"} 1.5867324e+09
publicholiday_date{countryCode="GB",county="",date="2020-05-08",fixed="false",global="true",launchYear="1978",localName="Early May Bank Holiday",name="Early May Bank Holiday",timezone="Europe/London",type="Public"} 1.5888924e+09
publicholiday_date{countryCode="GB",county="",date="2020-05-25",fixed="false",global="true",launchYear="1971",localName="Spring Bank Holiday",name="Spring Bank Holiday",timezone="Europe/London",type="Public"} 1.5903612e+09
publicholiday_date{countryCode="GB",county="",date="2020-12-25",fixed="false",global="true",launchYear="0",localName="Christmas Day",name="Christmas Day",timezone="Europe/London",type="Public"} 1.6088544e+09
publicholiday_date{countryCode="GB",county="",date="2020-12-28",fixed="false",global="true",launchYear="0",localName="Boxing Day",name="St. Stephen's Day",timezone="Europe/London",type="Public"} 1.6091136e+09
publicholiday_date{countryCode="GB",county="",date="2021-01-01",fixed="false",global="true",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Europe/London",type="Public"} 1.6094592e+09
publicholiday_date{countryCode="GB",county="",date="2021-04-02",fixed="false",global="true",launchYear="0",localName="Good Friday",name="Good Friday",timezone="Europe/London",type="Public"} 1.617318e+09
publicholiday_date{countryCode="GB",county="",date="2021-04-05",fixed="false",global="true",launchYear="0",localName="Easter Monday",name="Easter Monday",timezone="Europe/London",type="Public"} 1.6175772e+09
publicholiday_date{countryCode="GB",county="",date="2021-05-03",fixed="false",global="true",launchYear="1978",localName="Early May Bank Holiday",name="Early May Bank Holiday",timezone="Europe/London",type="Public"} 1.6199964e+09
publicholiday_date{countryCode="GB",county="",date="2021-05-31",fixed="false",global="true",launchYear="1971",localName="Spring Bank Holiday",name="Spring Bank Holiday",timezone="Europe/London",type="Public"} 1.6224156e+09
publicholiday_date{countryCode="GB",county="",date="2021-12-27",fixed="false",global="true",launchYear="0",localName="Boxing Day",name="St. Stephen's Day",timezone="Europe/London",type="Public"} 1.6405632e+09
publicholiday_date{countryCode="GB",county="",date="2021-12-28",fixed="false",global="true",launchYear="0",localName="Christmas Day",name="Christmas Day",timezone="Europe/London",type="Public"} 1.6406496e+09
publicholiday_date{countryCode="GB",county="GB-ENG",date="2020-08-31",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.5988284e+09
publicholiday_date{countryCode="GB",county="GB-ENG",date="2021-08-30",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.630278e+09
publicholiday_date{countryCode="GB",county="GB-NIR",date="2020-03-17",fixed="true",global="false",launchYear="0",localName="Saint Patrick's Day",name="Saint Patrick's Day",timezone="Europe/London",type="Public"} 1.5844032e+09
publicholiday_date{countryCode="GB",county="GB-NIR",date="2020-07-12",fixed="true",global="false",launchYear="0",localName="Battle of the Boyne",name="Battle of the Boyne",timezone="Europe/London",type="Public"} 1.5945084e+09
publicholiday_date{countryCode="GB",county="GB-NIR",date="2021-03-17",fixed="true",global="false",launchYear="0",localName="Saint Patrick's Day",name="Saint Patrick's Day",timezone="Europe/London",type="Public"} 1.6159392e+09
publicholiday_date{countryCode="GB",county="GB-NIR",date="2021-07-12",fixed="true",global="false",launchYear="0",localName="Battle of the Boyne",name="Battle of the Boyne",timezone="Europe/London",type="Public"} 1.6260444e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2020-01-02",fixed="false",global="false",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Europe/London",type="Public"} 1.5779232e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2020-08-03",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.5964092e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2020-11-30",fixed="true",global="false",launchYear="0",localName="Saint Andrew's Day",name="Saint Andrew's Day",timezone="Europe/London",type="Public"} 1.6066944e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2021-01-04",fixed="false",global="false",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Europe/London",type="Public"} 1.6097184e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2021-08-02",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.6278588e+09
publicholiday_date{countryCode="GB",county="GB-SCT",date="2021-11-30",fixed="true",global="false",launchYear="0",localName="Saint Andrew's Day",name="Saint Andrew's Day",timezone="Europe/London",type="Public"} 1.6382304e+09
publicholiday_date{countryCode="GB",county="GB-WLS",date="2020-08-31",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.5988284e+09
publicholiday_date{countryCode="GB",county="GB-WLS",date="2021-08-30",fixed="false",global="false",launchYear="1971",localName="Summer Bank Holiday",name="Summer Bank Holiday",timezone="Europe/London",type="Public"} 1.630278e+09
publicholiday_date{countryCode="NZ",county="",date="2020-01-01",fixed="false",global="true",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Pacific/Auckland",type="Public"} 1.57779e+09
publicholiday_date{countryCode="NZ",county="",date="2020-01-02",fixed="false",global="true",launchYear="0",localName="Day after New Year's Day",name="Day after New Year's Day",timezone="Pacific/Auckland",type="Public"} 1.5778764e+09
publicholiday_date{countryCode="NZ",county="",date="2020-02-06",fixed="false",global="true",launchYear="0",localName="Waitangi Day",name="Waitangi Day",timezone="Pacific/Auckland",type="Public"} 1.5809004e+09
publicholiday_date{countryCode="NZ",county="",date="2020-04-10",fixed="false",global="true",launchYear="0",localName="Good Friday",name="Good Friday",timezone="Pacific/Auckland",type="Public"} 1.5864336e+09
publicholiday_date{countryCode="NZ",county="",date="2020-04-13",fixed="false",global="true",launchYear="0",localName="Easter Monday",name="Easter Monday",timezone="Pacific/Auckland",type="Public"} 1.5866928e+09
publicholiday_date{countryCode="NZ",county="",date="2020-04-27",fixed="false",global="true",launchYear="0",localName="Anzac Day",name="Anzac Day",timezone="Pacific/Auckland",type="Public"} 1.5879024e+09
publicholiday_date{countryCode="NZ",county="",date="2020-06-01",fixed="false",global="true",launchYear="0",localName="Queen's Birthday",name="Queen's Birthday",timezone="Pacific/Auckland",type="Public"} 1.5909264e+09
publicholiday_date{countryCode="NZ",county="",date="2020-10-26",fixed="false",global="true",launchYear="0",localName="Labour Day",name="Labour Day",timezone="Pacific/Auckland",type="Public"} 1.6036236e+09
publicholiday_date{countryCode="NZ",county="",date="2020-12-25",fixed="false",global="true",launchYear="0",localName="Christmas Day",name="Christmas Day",timezone="Pacific/Auckland",type="Public"} 1.6088076e+09
publicholiday_date{countryCode="NZ",county="",date="2020-12-28",fixed="false",global="true",launchYear="0",localName="Boxing Day",name="Boxing Day",timezone="Pacific/Auckland",type="Public"} 1.6090668e+09
publicholiday_date{countryCode="NZ",county="",date="2021-01-01",fixed="false",global="true",launchYear="0",localName="New Year's Day",name="New Year's Day",timezone="Pacific/Auckland",type="Public"} 1.6094124e+09
publicholiday_date{countryCode="NZ",county="",date="2021-01-04",fixed="false",global="true",launchYear="0",localName="Day after New Year's Day",name="Day after New Year's Day",timezone="Pacific/Auckland",type="Public"} 1.6096716e+09
publicholiday_date{countryCode="NZ",county="",date="2021-02-08",fixed="false",global="true",launchYear="0",localName="Waitangi Day",name="Waitangi Day",timezone="Pacific/Auckland",type="Public"} 1.6126956e+09
publicholiday_date{countryCode="NZ",county="",date="2021-04-02",fixed="false",global="true",launchYear="0",localName="Good Friday",name="Good Friday",timezone="Pacific/Auckland",type="Public"} 1.6172748e+09
publicholiday_date{countryCode="NZ",county="",date="2021-04-05",fixed="false",global="true",launchYear="0",localName="Easter Monday",name="Easter Monday",timezone="Pacific/Auckland",type="Public"} 1.6175376e+09
publicholiday_date{countryCode="NZ",county="",date="2021-04-26",fixed="false",global="true",launchYear="0",localName="Anzac Day",name="Anzac Day",timezone="Pacific/Auckland",type="Public"} 1.619352e+09
publicholiday_date{countryCode="NZ",county="",date="2021-06-07",fixed="false",global="true",launchYear="0",localName="Queen's Birthday",name="Queen's Birthday",timezone="Pacific/Auckland",type="Public"} 1.6229808e+09
publicholiday_date{countryCode="NZ",county="",date="2021-10-25",fixed="false",global="true",launchYear="0",localName="Labour Day",name="Labour Day",timezone="Pacific/Auckland",type="Public"} 1.6350732e+09
publicholiday_date{countryCode="NZ",county="",date="2021-12-27",fixed="false",global="true",launchYear="0",localName="Christmas Day",name="Christmas Day",timezone="Pacific/Auckland",type="Public"} 1.6405164e+09
publicholiday_date{countryCode="NZ",county="",date="2021-12-28",fixed="false",global="true",launchYear="0",localName="Boxing Day",name="Boxing Day",timezone="Pacific/Auckland",type="Public"} 1.6406028e+09
```


The Public holidays are fetched from [date.nager](https://date.nager.at/)
