package model

import "time"

type Todo struct {
	Id          int `gorm:"primaryKey"`
	Description string
	Statu       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
