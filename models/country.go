package models

type Country struct {
	Id              int    `json:"id" gorm:"primaryKey;column:Id"`                    // 主键
	CountryName     string `json:"country_name" gorm:"column:COUNTRY_NAME"`           // 国家名
	CountryCnName   string `json:"country_cn_name" gorm:"column:COUNTRY_CN_NAME"`     // 国家中文名
	CountryFullName string `json:"country_full_name" gorm:"column:COUNTRY_FULL_NAME"` // 国家全名
	BaseEntity
}

func (country Country) TableName() string {
	return "COUNTRY"
}
