package models

type City struct {
	ID        int    `json:"id" gorm:"primaryKey;column:Id"`      // 主键
	CityName  string `json:"city_name" gorm:"column:CITY_NAME"`   // 城市名
	CityType  int    `json:"city_type" gorm:"column:CITY_TYPE"`   // 城市类型
	ParentID  int    `json:"parent_id" gorm:"column:PARENT_ID"`   // 父节点ID
	CountryID int    `json:"country_id" gorm:"column:COUNTRY_ID"` // 国家ID
	Sorts     int    `json:"sorts" gorm:"column:SORTS"`           // 排序
	BaseEntity
}

func (city City) TableName() string {
	return "CITY"
}

type CityType struct {
	ID        int    `json:"id" gorm:"primaryKey;column:Id"`      // 主键
	Name      string `json:"Name" gorm:"column:NAME"`             // 类型
	CountryID int    `json:"country_id" gorm:"column:COUNTRY_ID"` // 国家ID
	BaseEntity
}

func (cityType CityType) TableName() string {
	return "CITY_TYPE"
}

// CityTypeEnum define enum
type CityTypeEnum struct {
	BaseEnum
}

var (
	CityTypeEnumProvince = CityTypeEnum{BaseEnum{Id: 1, Name: "province"}}
	CityTypeEnumCity     = CityTypeEnum{BaseEnum{Id: 2, Name: "city"}}
	CityTypeEnumArea     = CityTypeEnum{BaseEnum{Id: 3, Name: "area"}}
	CityTypeEnumTown     = CityTypeEnum{BaseEnum{Id: 4, Name: "town"}}
)

func (receiver CityTypeEnum) List() []CityTypeEnum {
	return []CityTypeEnum{CityTypeEnumProvince, CityTypeEnumCity, CityTypeEnumArea, CityTypeEnumTown}
}
