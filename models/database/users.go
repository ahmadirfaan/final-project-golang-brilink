package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Id         *uint          `gorm:"autoIncrement;primary key" json:"-"`
	Username   string         `gorm:"uniqueIndex;type:varchar(255);not null"`
	Password   string         `gorm:"type:varchar(255);not null"`
	RoleId     uint
	Role       Role `gorm:"foreignkey:RoleId;references:Id"`
	CustomerId *uint
	Customer   Customer `gorm:"foreignkey:CustomerId;references:Id"`
	AgentId    *uint
	Agent      Agent `gorm:"foreignkey:AgentId;references:Id"`
}
