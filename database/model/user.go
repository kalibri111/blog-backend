package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Articles  []Article
	Photos    []Photo
	Login     string
	Password  string
	Email     string
	Telegram  string
	Instagram string
}
