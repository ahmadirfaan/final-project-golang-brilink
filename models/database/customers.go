package database

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Id          uint           `gorm:"autoIncrement;primary key"`
	Name        string         `gorm:"varchar(255); not null"`
	NoHandphone string         `gorm:"varchar(12);not null"`
}
