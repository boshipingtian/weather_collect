package collect

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
	"weather-colly/global"
	"weather-colly/models"
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
		Where("C.CITY_TYPE = ?", models.CityTypeEnumCity.Id).
		Scan(&results).Error
	if err != nil {
		global.Logger.Errorln("查询失败:", err)
		return
	}
	i := 1
	for _, item := range results {
		weather, _ := collectWeather(item.CityPinyin + "-" + strconv.Itoa(item.Code))
		global.Logger.Infoln(weather)
		modelsWeather := convertToModelsWeather(weather, item)
		global.Logger.Infoln(modelsWeather)
		insertToDatabase(&modelsWeather)
		global.Logger.Infoln(fmt.Sprintf("当前 %d，剩余 %d", i, len(results)-i))
		i++
	}
}

// collectWeather 采集一个页面上的数据
func collectWeather(cityUrl string) (CollyWeather, string) {
	defer func() {
		if r := recover(); r != nil {
			global.Logger.Errorln(fmt.Sprintf("收集%s天气数据时发生错误:", cityUrl), r)
		}
	}()

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
		global.Logger.Errorln(fmt.Sprintf("采集出现异常，定位标记：%s", cityUrl), err.Error())
		return CollyWeather{}, ""
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

func convertToModelsWeather(weather CollyWeather, weatherPinyin WeatherPinyin) []models.Weather {
	datetime := weather.Datetime
	if datetime == nil {
		return nil
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	year, month, day := datetime.Date()
	hour, _, _ := datetime.Clock()
	date := time.Date(year, month, day, 0, 0, 0, 0, loc)
	timeHour := time.Date(year, month, day, hour, 0, 0, 0, loc)

	var results []models.Weather
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeTemperature, fmt.Sprint(weather.Temperature)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeBodyTemperature, fmt.Sprint(weather.BodyTemperature)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeHumidity, fmt.Sprint(weather.Humidity)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeVisibility, fmt.Sprint(weather.Visibility)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypePrecipitation, fmt.Sprint(weather.Precipitation)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypePressure, fmt.Sprint(weather.Pressure)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeWind, fmt.Sprint(weather.Wind)))
	results = append(results, buildModelWeather(&date, &timeHour, weatherPinyin.Code,
		models.WeatherTypeWindDirection, fmt.Sprint(weather.WindDirection)))
	return results
}

func buildModelWeather(date *time.Time, timeHour *time.Time, cityCode int,
	typing models.WeatherTypeEnum, value string) models.Weather {
	mWeather := models.Weather{
		Date:     date,
		Time:     timeHour,
		CityCode: cityCode,
		Type:     typing.Id,
		Value:    value,
	}
	return mWeather
}

func insertToDatabase(weather *[]models.Weather) {
	db := global.DB
	if weather == nil || len(*weather) == 0 {
		return
	}
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚", r)
		}
	}()
	for _, item := range *weather {
		result := db.Model(&models.Weather{}).
			Where("CITY_CODE =? AND TYPE =? AND DATE =? AND TIME =?", item.CityCode, item.Type, item.Date, item.Time).
			Updates(item)
		global.Logger.Debugln(fmt.Sprintf("更新影响了 %d", result.RowsAffected))
		if result.RowsAffected == 0 {
			// 更新影响行数为 0，进行插入
			if err := tx.Create(&item).Error; err != nil {
				tx.Rollback()
				global.Logger.Errorln("事务回滚", err)
				return
			}
		}
	}
	tx.Commit()
	global.Logger.Infoln(fmt.Sprintf("入库 %d 条", len(*weather)))
}
