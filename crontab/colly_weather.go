package crontab

import "weather-colly/global"

type CollyWeather struct {
}

func (c CollyWeather) getName() string {
	return "colly_weather"
}

func (c CollyWeather) Run() {
	global.Logger.Infoln("cronjob")
}
