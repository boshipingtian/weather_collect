package models

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"column:CREATED_TIME;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:UPDATED_TIME;autoUpdateTime:milli"`
}
