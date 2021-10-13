package service

import (
    "errors"
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
	transactionsRepository    repositories.TransactionRepository
	transactionTypeRepository repositories.TransactionTypeRepository
    districtRepository repositories.DistrictRepository
	DB                        *gorm.DB
}

func NewTransactionService(tr repositories.TransactionRepository,
	ttr repositories.TransactionTypeRepository,
    dr repositories.DistrictRepository,
	db *gorm.DB) TransactionService {
	return &transactionService{
		transactionsRepository:    tr,
		DB:                        db,
		transactionTypeRepository: ttr,
        districtRepository: dr,
	}
}

func (t *transactionService) CreateTransaction(request web.CreateTransactionRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
    //validate userId must not same with agentId
    if request.AgentId == request.CustomerId {
        return errors.New("Agent Id and Customer Id Must not the same")
    }
    //validate exist districtId
	_,err = t.districtRepository.FindById(request.DistrictId)
	if err != nil {
		return err
	}
    //validate exist transactionType
    err = t.transactionTypeRepository.FindById(request.TransactionTypeId)
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
		Address:           request.Address,
		StatusTransaction: 0,
	}
	transaction, err = t.transactionsRepository.WithTrx(tx).Save(transaction)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	return tx.Commit().Error
}
