package model

import "time"

type UserBirthday struct {
	ID   uint       `gorm:"primary_key"`
	Name string     `gorm:"not null"`
	Birthday time.Time
}