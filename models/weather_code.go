package models

type WeatherCityCode struct {
	ID         int    `json:"id" db:"id"`                   // 主键
	CityID     int    `json:"city_id" db:"city_id"`         // 城市名
	Code       int    `json:"code" db:"code"`               // 气象代码
	CityPinyin string `json:"city_pinyin" db:"city_pinyin"` // 气象城市拼音
	BaseEntity
}

func (c WeatherCityCode) TableName() string {
	return "weather_city_code"
}
