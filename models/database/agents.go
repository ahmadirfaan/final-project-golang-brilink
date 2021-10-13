package database

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Id          uint           `gorm:"autoIncrement;primary key" json:"-"`
	AgentName   string         `gorm:"type:varchar(250);not null"`
	DistrictId  string         `gorm:"type:char(7);not null"`
	Address     string         `gorm:"type:text;not null"`
	NoHandphone string         `gorm:"type:varchar(12);not null"`
	Rating      *uint8          `gorm:"type:tinyint"`
}
