package database

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Id        uint           `gorm:"autoIncrement"`
	Username  string         `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password  string         `gorm:"uniqueIndex;type:varchar(255);not null"`
}
