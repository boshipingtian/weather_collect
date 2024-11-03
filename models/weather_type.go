package models

type WeatherType struct {
	ID   int    `gorm:"primaryKey" json:"id" comment:"主键"`
	Name string `gorm:"type:varchar(255);not null" json:"name" comment:"类型"`
	Unit string `gorm:"type:varchar(255);not null" json:"unit" comment:"单位"`
	BaseEntity
}

func (WeatherType) TableName() string {
	return "WEATHER_TYPE"
}

type WeatherTypeEnum struct {
	BaseEnum
	Unit string // 单位
}

var (
	WeatherTypeTemperature = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 1, Name: "temperature"},
		Unit:     "°C",
	}
	WeatherTypeBodyTemperature = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 2, Name: "body temperature"},
		Unit:     "°C",
	}
	WeatherTypeHumidity = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 3, Name: "humidity"},
		Unit:     "%",
	}
	WeatherTypeVisibility = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 4, Name: "visibility"},
		Unit:     "km",
	}
	WeatherTypePrecipitation = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 5, Name: "precipitation"},
		Unit:     "mm",
	}
	WeatherTypePressure = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 6, Name: "pressure"},
		Unit:     "hPa",
	}
	WeatherTypeWind = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 7, Name: "wind"},
		Unit:     "级",
	}
	WeatherTypeWindDirection = WeatherTypeEnum{
		BaseEnum: BaseEnum{Id: 8, Name: "wind direction"},
	}
)

func (e WeatherTypeEnum) List() []WeatherTypeEnum {
	return []WeatherTypeEnum{WeatherTypeTemperature, WeatherTypeBodyTemperature, WeatherTypeHumidity,
		WeatherTypeVisibility, WeatherTypePrecipitation, WeatherTypePressure, WeatherTypeWind, WeatherTypeWindDirection}
}
