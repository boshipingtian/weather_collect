package services

import (
	"weather-colly/global"
	"weather-colly/models"
)

func DeleteAllWeatherType() error {
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚", r)
		}
	}()
	if err := tx.Exec("TRUNCATE TABLE weather_colly.WEATHER_TYPE").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func InsertAllWeatherType() error {
	db := global.DB
	list := models.WeatherTypeEnum{}.List()
	var weatherTypes []models.WeatherType
	for _, item := range list {
		weatherType := models.WeatherType{
			ID:   item.Id,
			Name: item.Name,
			Unit: item.Unit,
		}
		weatherTypes = append(weatherTypes, weatherType)
	}
	tx := db.Begin()
	if err := tx.Create(&weatherTypes).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	global.Logger.Infoln("insert all weather type success!")
	return nil
}
