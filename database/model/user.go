package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Articles  []Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Login     string
	Password  string
	Email     string
	Telegram  string
	Instagram string
}
