package models

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `gorm:"column:created_time;autoCreateTime"`
	UpdatedAt *time.Time `gorm:"column:updated_time;autoUpdateTime:milli"`
}

type BaseEnum struct {
	Id   int    // 枚举ID
	Name string // 枚举名
}
