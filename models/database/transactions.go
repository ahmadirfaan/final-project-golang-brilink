package database

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt  `gorm:"index"`
	Id                uint            `gorm:"autoIncrement;primary key" json:"-"`
	TransactionTypeId uint            `gorm:"not null"`
	TransactionType   TransactionType `gorm:"foreignkey:TransactionTypeId;references:Id";`
	CustomerId        uint            `gorm:"not null"`
	AgentId           uint            `gorm:"not null"`
	Address           string          `gorm:"type:text;not null"`
	DistrictId        uint            `gorm:"type:char(7);not null"`
	Amount            uint64          `gorm:"not null"`
	StatusTransaction uint8           `gorm:"not null;type:tinyint"`
}
