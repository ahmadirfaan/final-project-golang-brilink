package database

import (
	"gorm.io/gorm"
	"time"
)

type ServiceTypeTransaction struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Id          uint           `gorm:"autoIncrement;primary key" json:"-"`
	NameService string         `gorm:"varchar(255);not null"`
}
