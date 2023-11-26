package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Header  string
	Content string
}
