package models

type WeatherCityCode struct {
	ID          int    `json:"id" db:"Id"`                     // 主键
	CityID      int    `json:"city_id" db:"CITY_ID"`           // 城市名
	Code        int    `json:"code" db:"CODE"`                 // 气象代码
	CityPingyin string `json:"city_pingyin" db:"CITY_PINGYIN"` // 气象城市拼音
	BaseEntity
}

func (c WeatherCityCode) TableName() string {
	return "WEATHER_CITY_CODE"
}
