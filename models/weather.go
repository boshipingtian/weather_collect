package models

import (
	"time"
)

type Weather struct {
	ID       int        `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	Date     *time.Time `json:"date" gorm:"column:date"`
	Time     *time.Time `json:"time" gorm:"column:time"`
	CityCode int        `json:"cityCode" gorm:"column:city_code"`
	Type     int        `json:"type" gorm:"column:type"`
	Value    string     `json:"value" gorm:"column:value"`
	BaseEntity
}

func (w Weather) TableName() string {
	return "weather"
}
