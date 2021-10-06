package database

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Id          uint           `gorm:"autoIncrement;primary key"`
	AgentName   string         `gorm:"varchar(250);not null"`
	Address     string         `gorm:"type:text;not null"`
	Province    string         `gorm:"varchar(250);not null"`
	City        string         `gorm:"varchar(250);not null"`
	District    string         `gorm:"varchar(250);not null"`
	NoHandphone string         `gorm:"varchar(12);not null"`
}
