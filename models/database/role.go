package database

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Id        uint           `gorm:"autoIncrement;primary key"`
	Role      string         `gorm:"varchar(50)"`
}
