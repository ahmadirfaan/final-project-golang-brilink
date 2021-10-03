package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Id         uint           `gorm:"autoIncrement;primary key"`
	RoleId     uint
	Role       Role `gorm:"foreignkey:RoleId;references:Id"`
	CustomerId uint
	// Belong to CustomerId
	Customer Customer `gorm:"foreignkey:CustomerId;references:Id"`
	AgentId  uint
	// Belong to AgentId
	Agent Agents `gorm:"foreignkey:AgentId;references:Id"`
}
