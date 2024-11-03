package main

import (
	"weather-colly/global"
	"weather-colly/initial"
	"weather-colly/services"
	"weather-colly/utils"
)

func main() {
	initial.Init()
	global.Logger.Infoln("initial is running!")
	global.Logger.Infoln("initial country!")
	utils.ExecuteWithErrorHandling(services.InsertAllCountry)
	global.Logger.Infoln("initial city!")
	utils.ExecuteWithErrorHandling(services.InsertAllCity)
	global.Logger.Infoln("initial weather city code!")
	utils.ExecuteWithErrorHandling(services.InsertAllWeatherCode)
}
