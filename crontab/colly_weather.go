package crontab

import (
	"fmt"
	"time"
	"weather-colly/collect"
	"weather-colly/global"
)

type CollyWeather struct {
}

func (c CollyWeather) getName() string {
	return "colly_weather"
}

func (c CollyWeather) Run() {
	now := time.Now()
	global.Logger.Infoln(fmt.Sprintf("%s 开始执行 气象采集程序",
		now.Format("2006-01-02 15:04:05")))
	collect.CityCollect()
}
