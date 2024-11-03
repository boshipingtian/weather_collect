package services

import (
	"weather-colly/entity"
	"weather-colly/global"
	"weather-colly/models"
)

func DeleteAllCountry() error {
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚！", r)
		}
	}()
	if err := tx.Exec("TRUNCATE TABLE weather_colly.COUNTRY").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	global.Logger.Infoln("delete all country success!")
	return nil
}

func InsertAllCountry() error {
	db := global.DB
	countries := entity.ReadCountry()
	var modelsCounties []models.Country
	for _, item := range countries {
		modelsCounties = append(modelsCounties, models.Country{
			Id:              item.Id,
			CountryCnName:   item.CnName,
			CountryName:     item.Name,
			CountryFullName: item.FullName,
		})
	}
	tx := db.Begin()
	if err := tx.Create(&modelsCounties).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	global.Logger.Infoln("insert all country success!")
	return nil
}

func FindCountry(name string) models.Country {
	db := global.DB
	country := models.Country{}
	if err := db.Where(&models.Country{CountryCnName: name}).First(&country).Error; err != nil {
		global.Logger.Errorln(err)
	}
	return country
}
