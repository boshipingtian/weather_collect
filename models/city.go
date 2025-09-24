package models

type City struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`      // 主键
	CityName  string `json:"city_name" gorm:"column:city_name"`   // 城市名
	CityType  int    `json:"city_type" gorm:"column:city_type"`   // 城市类型
	ParentID  int    `json:"parent_id" gorm:"column:parent_id"`   // 父节点ID
	CountryID int    `json:"country_id" gorm:"column:country_id"` // 国家ID
	Sorts     int    `json:"sorts" gorm:"column:sorts"`           // 排序
	BaseEntity
}

func (city City) TableName() string {
	return "city"
}

type CityType struct {
	ID        int    `json:"id" gorm:"primaryKey;column:id"`      // 主键
	Name      string `json:"Name" gorm:"column:name"`             // 类型
	CountryID int    `json:"country_id" gorm:"column:country_id"` // 国家ID
	BaseEntity
}

func (cityType CityType) TableName() string {
	return "city_type"
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
