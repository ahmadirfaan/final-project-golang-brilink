package service

import (
	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"github.com/itp-backend/backend-b-antar-jemput/models/web"
	"github.com/itp-backend/backend-b-antar-jemput/repositories"
	"github.com/itp-backend/backend-b-antar-jemput/utils"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(request web.CreateTransactionRequest) error
}

type transactionService struct {
	transactionsRepository repositories.TransactionRepository
	DB                     *gorm.DB
}

func NewTransacrtionService(tr repositories.TransactionRepository, db *gorm.DB) TransactionService {
	return &transactionService{
		transactionsRepository: tr,
		DB:                     db,
	}
}

func (t *transactionService) CreateTransaction(request web.CreateTransactionRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
	tx := t.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	transaction := database.Transactions{
		TransactionTypeId: request.TransactionTypeId,
		CustomerId:        request.CustomerId,
		AgentId:           request.AgentId,
		Amount:            request.Amount,
		StatusTransaction: 0,
	}
	transaction, err = t.transactionsRepository.WithTrx(tx).Save(transaction)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	return tx.Commit().Error
}
