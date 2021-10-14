package database

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	CreatedAt         time.Time       `json:"-"`
	UpdatedAt         time.Time       `json:"-"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`
	Id                uint            `gorm:"autoIncrement;primary key" json:"id"`
	TransactionTypeId uint            `gorm:"not null" json:"-"`
	TransactionType   TransactionType `gorm:"foreignkey:TransactionTypeId;references:Id" json:"transactionType"`
	CustomerId        uint            `gorm:"not null" json:"customerId"`
	AgentId           uint            `gorm:"not null" json:"agentId"`
	Address           string          `gorm:"type:text;not null" json:"address"`
	DistrictId        string          `gorm:"type:char(7);not null" json:"districtId"`
	Amount            uint64          `gorm:"not null" json:"amount"`
	StatusTransaction uint8           `gorm:"not null;type:tinyint" json:"statusTransaction"`
}
