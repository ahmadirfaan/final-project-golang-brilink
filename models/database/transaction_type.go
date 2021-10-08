package database

import (
	"gorm.io/gorm"
	"time"
)

type TransactionType struct {
	CreatedAt                time.Time
	UpdatedAt                time.Time
	DeletedAt                gorm.DeletedAt `gorm:"index"`
	Id                       uint           `gorm:"autoIncrement;primary key"`
	ServiceTypeTransactionId uint
	ServiceTypeTransaction   ServiceTypeTransaction `gorm:"foreignkey:ServiceTypeTransactionId;references:Id;not null""`
	NameTypeTransaction      string                 `gorm:"varchar(255);not null"`
}
