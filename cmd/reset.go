package main

import (
	"weather-colly/global"
	"weather-colly/initial"
	"weather-colly/services"
	"weather-colly/utils"
)

func main() {
	initial.Init()
	global.Logger.Infoln("reset is running!")

	global.Logger.Infoln("reset country!")
	utils.ExecuteWithErrorHandling(services.DeleteAllCountry)
	global.Logger.Infoln("reset city")
	utils.ExecuteWithErrorHandling(services.DeleteAllCity)
	global.Logger.Infoln("reset weather city code")
	utils.ExecuteWithErrorHandling(services.DeleteAllWeatherCode)

}
