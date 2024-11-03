package main

import (
	"weather-colly/collect"
	"weather-colly/initial"
)

func main() {
	initial.Init()
	collect.CityCollect()
}
