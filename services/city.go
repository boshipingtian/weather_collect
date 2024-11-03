package services

import (
	"fmt"
	"strconv"
	"weather-colly/entity"
	"weather-colly/global"
	"weather-colly/models"
)

func DeleteAllCity() error {
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚！", r)
		}
	}()
	tx.Exec("TRUNCATE TABLE weather_colly.CITY")
	if tx.Error != nil {
		db.Rollback()
		return tx.Error
	}
	tx.Commit()
	global.Logger.Infoln("delete all city success!")
	return nil
}

func InsertAllCity() error {
	db := global.DB
	modelsCities, _ := createCityModels()
	global.Logger.Infoln(fmt.Sprintf("ready to insert %d city!", len(modelsCities)))
	tx := db.Begin()
	if err := tx.Create(&modelsCities).Error; err != nil {
		tx.Rollback()
		return tx.Error
	}
	tx.Commit()
	global.Logger.Infoln("insert all country success!")
	return nil
}

func createCityModels() ([]models.City, error) {

	var result []models.City
	// build province list
	provinceMap, err := createProvince(&result)
	if err != nil {
		global.Logger.Errorln(err)
		return nil, err
	}
	// build city
	_, err = createCity(&result, provinceMap)
	if err != nil {
		global.Logger.Errorln(err)
		return nil, err
	}
	return result, nil
}

func createProvince(result *[]models.City) (map[int]models.City, error) {
	province := entity.ReadProvince()
	provinceMap := make(map[int]models.City)
	for _, item := range province {
		provinceId, err := strconv.Atoi(item.Code)
		if err != nil {
			return nil, err
		}
		sortCode, err := strconv.Atoi(item.Province)
		if err != nil {
			return nil, err
		}
		provinceCity := models.City{
			ID:       provinceId,
			CityType: models.CityTypeEnumProvince.Id,
			CityName: item.Name,
			// china
			CountryID: 45,
			Sorts:     sortCode,
		}
		provinceMap[provinceCity.Sorts] = provinceCity
		*result = append(*result, provinceCity)
	}
	return provinceMap, nil
}

func createCity(result *[]models.City, provinceMap map[int]models.City) (map[int][]models.City, error) {
	cities := entity.ReadCity()
	var cityProvinceMap = make(map[int][]models.City)
	for _, item := range cities {
		provinceKey, err := strconv.Atoi(item.Province)
		if err != nil {
			return nil, err
		}
		cityId, err := strconv.Atoi(item.Code)
		if err != nil {
			return nil, err
		}
		citySortId, err := strconv.Atoi(item.City)
		if err != nil {
			return nil, err
		}

		city := models.City{
			ID:        cityId,
			CityName:  item.Name,
			CityType:  models.CityTypeEnumCity.Id,
			ParentID:  provinceMap[provinceKey].ID,
			CountryID: 45,
			Sorts:     citySortId,
		}

		cityProvinceMap[provinceKey] = append(cityProvinceMap[provinceKey], city)
		*result = append(*result, city)
	}
	return cityProvinceMap, nil
}

func DeleteAllCityType() error {
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			global.Logger.Errorln("事务回滚", r)
		}
	}()
	if err := tx.Exec("TRUNCATE TABLE weather_colly.CITY_TYPE").Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func InsertAllCityType() error {
	db := global.DB
	list := models.CityTypeEnum{}.List()
	var cityTypes []models.CityType
	for _, item := range list {
		cityType := models.CityType{
			ID:        item.Id,
			Name:      item.Name,
			CountryID: 45,
		}
		cityTypes = append(cityTypes, cityType)
	}
	tx := db.Begin()
	if err := tx.Create(&cityTypes).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	global.Logger.Infoln("insert all city type success!")
	return nil
}
