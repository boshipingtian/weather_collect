package collect

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
	"weather-colly/global"
)

var (
	basePath = "https://www.qweather.com/weather/%s.html"
)

type CollyWeather struct {
	City            string     // 城市
	Datetime        *time.Time // 时间
	Temperature     float64    // 温度
	Humidity        float64    // 湿度
	BodyTemperature float64    // 体感温度
	Visibility      float64    // 能见度
	Precipitation   float64    // 降雨量
	Pressure        float64    // 大气压
	Wind            float64    // 风速
	WindDirection   string     // 风向
}

type WeatherPinyin struct {
	CityName   string `gorm:"column:CITY_NAME"`
	CityPinyin string `gorm:"column:CITY_PINYIN"`
	Code       int    `gorm:"column:CODE"`
}

func CityCollect() {
	db := global.DB
	var results []WeatherPinyin
	err := db.Table("CITY as C").
		Select("C.CITY_NAME, WCC.CITY_PINYIN, WCC.CODE").
		Joins("inner join weather_colly.WEATHER_CITY_CODE WCC on C.ID = WCC.CITY_ID").
		Where("C.CITY_TYPE = ?", 2).
		Scan(&results).Error
	if err != nil {
		global.Logger.Errorln("查询失败:", err)
		return
	}
	for _, item := range results {
		weather, _ := collectWeather(item.CityPinyin + "-" + strconv.Itoa(item.Code))
		global.Logger.Infoln(weather)
	}
}

// collectWeather 采集一个页面上的数据
func collectWeather(cityUrl string) (CollyWeather, string) {
	collector := colly.NewCollector()
	// 设置请求头，模拟浏览器
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", CommonUserAgent)
	})
	var weatherArray [7][2]string
	// 处理响应
	collector.OnHTML("body > div.body-content.body-content--subpage > div.l-page-city-weather > div > div.l-page-city-weather__current > div > div > div.current-basic.d-flex.justify-content-between.align-items-center", func(e *colly.HTMLElement) {
		e.ForEach("div.current-basic___item", func(i int, element *colly.HTMLElement) {
			element.ForEach("p", func(j int, element *colly.HTMLElement) {
				weatherArray[i][j] = strings.TrimSpace(element.Text)
			})
		})
	})
	var city string
	collector.OnHTML("body > div.body-content.body-content--subpage > div.c-submenu > div.c-submenu__bg.hidden-740.jsSubmenu > div > div.c-submenu__scroll-container > div > div.d-flex.align-items-center > h1", func(e *colly.HTMLElement) {
		city = strings.TrimSpace(e.Text)
	})
	var currentTime string
	collector.OnHTML("body > div.body-content.body-content--subpage > div.l-page-city-weather > div > div.l-page-city-weather__current > div > div > p", func(e *colly.HTMLElement) {
		currentTime = e.Text
	})
	var temperatureStr string
	collector.OnHTML("body > div.body-content.body-content--subpage > div.l-page-city-weather > div > div.l-page-city-weather__current > div > div > div.current-live > div:nth-child(2) > p:nth-child(1)", func(e *colly.HTMLElement) {
		temperatureStr = strings.TrimSpace(e.Text)
	})
	// 爬取页面
	err := collector.Visit(fmt.Sprintf(basePath, cityUrl))
	if err != nil {
		global.Logger.Errorln(err.Error())
	}
	weather := CollyWeather{}
	weather.City = city
	temperature := temperatureStr[:strings.LastIndex(temperatureStr, "°")]
	if val, err := strconv.ParseFloat(temperature, 64); err == nil {
		weather.Temperature = val
	}
	if val, err := time.Parse("2006-01-02 15:04", currentTime); err == nil {
		weather.Datetime = &val
	} else {
		global.Logger.Errorln(err.Error())
	}
	for index, item := range weatherArray {
		switch index {
		case 0:
			wind := item[0][:strings.LastIndex(item[0], "级")]
			if val, err := strconv.ParseFloat(wind, 64); err == nil {
				weather.Wind = val
			}
			windDirection := item[1]
			weather.WindDirection = windDirection
			break
		case 1:
			humidity := item[0][:strings.LastIndex(item[0], "%")]
			if val, err := strconv.ParseFloat(humidity, 64); err == nil {
				weather.Humidity = val
			}
			break
		case 3:
			bodyTemperature := item[0][:strings.LastIndex(item[0], "°")]
			if val, err := strconv.ParseFloat(bodyTemperature, 64); err == nil {
				weather.BodyTemperature = val
			}
			break
		case 4:
			visibility := item[0][:strings.LastIndex(item[0], "km")]
			if val, err := strconv.ParseFloat(visibility, 64); err == nil {
				weather.Visibility = val
			}
			break
		case 5:
			precipitation := item[0][:strings.LastIndex(item[0], "mm")]
			if val, err := strconv.ParseFloat(precipitation, 64); err == nil {
				weather.Precipitation = val
			}
			break
		case 6:
			pressure := item[0][:strings.LastIndex(item[0], "hPa")]
			if val, err := strconv.ParseFloat(pressure, 64); err == nil {
				weather.Pressure = val
			}
			break
		}
	}
	return weather, currentTime
}
