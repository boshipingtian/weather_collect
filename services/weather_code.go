package services

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"strings"
	"weather-colly/collect"
	"weather-colly/global"
	"weather-colly/models"
)

func DeleteAllWeatherCode() error {
	// start transaction
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚:", r)
		}
	}()
	if err := db.Exec("TRUNCATE TABLE weather_colly.WEATHER_CITY_CODE").Error; err != nil {
		tx.Rollback()
		return tx.Error
	}
	// commit
	tx.Commit()
	global.Logger.Infoln("delete all weather_city_code success!")
	return nil
}

func InsertAllWeatherCode() error {
	db := global.DB
	var cityNameMap = make(map[string]int)
	var cities []models.City

	if err := db.Where(models.City{CityType: models.CityTypeEnumCity.Id}).Find(&cities).Error; err != nil {
		return err
	}
	for _, item := range cities {
		cityName := item.CityName
		if strings.HasSuffix(cityName, "市") {
			cityName = cityName[:len(cityName)-3] // 截取去掉最后的“市”
		}
		cityNameMap[cityName] = item.ID
	}
	weatherCode := collect.CollyWeatherCode()
	var weatherCityCode []models.WeatherCityCode
	var existCityCode = make(map[int]int)
	for _, item := range weatherCode {
		if value, exist := cityNameMap[item.Name]; exist {
			if _, ok := existCityCode[value]; !ok {
				pingyin := createPingyin(item.Name)
				cityCode := models.WeatherCityCode{CityID: value, Code: item.Code, CityPingyin: pingyin}
				weatherCityCode = append(weatherCityCode, cityCode)
				existCityCode[value] = value
			}
		}
	}
	global.Logger.Infoln(fmt.Sprintf("准备插入 %d 条天气代码!", len(weatherCityCode)))

	tx := db.Begin()
	defer func() {
		// recover()管理panic,如果出现,则不为空
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚:", r)
		}
	}()
	if err := tx.Create(&weatherCityCode).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	global.Logger.Infoln(fmt.Sprintf("成功插入 %d 条天气代码!", len(weatherCityCode)))
	return nil
}

func createPingyin(name string) string {
	args := pinyin.NewArgs()
	pinyinList := pinyin.Pinyin(name, args)
	var strs []string
	for _, arr := range pinyinList {
		strs = append(strs, arr[0])
	}
	return strings.Join(strs, "")
}
