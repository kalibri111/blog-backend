package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	UserID  uint
	Header  string
	Content string
}
