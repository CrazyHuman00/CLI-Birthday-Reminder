package model

import "time"

type Birthday struct {
	ID   uint       `gorm:"primary_key"`
	Name string     `gorm:"uniqueIndex"`
	Date time.Time
}