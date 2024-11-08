package models

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"column:CREATED_TIME;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:UPDATED_TIME;autoUpdateTime:milli"`
}

type BaseEnum struct {
	Id   int    // 枚举ID
	Name string // 枚举名
}
