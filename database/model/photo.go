package model

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	Camera string
	Lens   string
	Date   time.Time
}
