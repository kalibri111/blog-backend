package model

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	UserID uint
	Camera string
	Lens   string
	Photo  string
	Date   time.Time
}
