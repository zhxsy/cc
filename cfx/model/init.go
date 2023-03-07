package model

import (
	"gorm.io/gorm"
	"time"
)

type Basic struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
