package models

import (
	"time"
)

type Weather struct {
	ID       int        `json:"id" gorm:"primaryKey;autoIncrement;column:ID"`
	Date     *time.Time `json:"date" gorm:"column:DATE"`
	Time     *time.Time `json:"time" gorm:"column:TIME"`
	CityCode int        `json:"cityCode" gorm:"column:CITY_CODE"`
	Type     int        `json:"type" gorm:"column:TYPE"`
	Value    string     `json:"value" gorm:"column:VALUE"`
	BaseEntity
}

func (w Weather) TableName() string {
	return "WEATHER"
}
