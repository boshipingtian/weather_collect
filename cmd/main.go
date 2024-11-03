package main

import (
	"weather-colly/global"
	"weather-colly/initial"
	"weather-colly/services"
)

func main() {
	initial.Init()

	country := services.FindCountry("中国")

	global.Logger.Infoln(country)
}
