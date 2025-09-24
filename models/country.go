package models

type Country struct {
	Id              int    `json:"id" gorm:"primaryKey;column:id"`                    // 主键
	CountryName     string `json:"country_name" gorm:"column:country_name"`           // 国家名
	CountryCnName   string `json:"country_cn_name" gorm:"column:country_cn_name"`     // 国家中文名
	CountryFullName string `json:"country_full_name" gorm:"column:country_full_name"` // 国家全名
	BaseEntity
}

func (country Country) TableName() string {
	return "country"
}
