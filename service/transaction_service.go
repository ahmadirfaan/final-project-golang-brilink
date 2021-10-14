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
	GetAllTransactionByUserId(userId string) ([]database.Transactions, error)
}

type transactionService struct {
	transactionsRepository    repositories.TransactionRepository
	transactionTypeRepository repositories.TransactionTypeRepository
	districtRepository        repositories.DistrictRepository
	userRepository            repositories.UserRepository
	DB                        *gorm.DB
}

func NewTransactionService(tr repositories.TransactionRepository,
	ttr repositories.TransactionTypeRepository,
	dr repositories.DistrictRepository,
	ur repositories.UserRepository,
	db *gorm.DB) TransactionService {
	return &transactionService{
		transactionsRepository:    tr,
		DB:                        db,
		transactionTypeRepository: ttr,
		districtRepository:        dr,
		userRepository:            ur,
	}
}

func (t *transactionService) GetAllTransactionByUserId(userId string) ([]database.Transactions, error) {
	transactions, err := t.transactionsRepository.FindTransactionWithUserId(userId)
	return transactions, err
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
	_, err = t.districtRepository.FindById(request.DistrictId)
	if err != nil {
		return err
	}
	//validate exist transactionType
	err = t.transactionTypeRepository.FindById(request.TransactionTypeId)
	if err != nil {
		return err
	}
	//validate if the request customerId is customer
	isAgent, err := t.userRepository.IsUserAgent(request.CustomerId)
	//validate the exist is customerId
	isExistUserCustomer, err := t.userRepository.IsExist(request.CustomerId)
	//validate if the request agentId is agent
	isNotAgent, err := t.userRepository.IsUserAgent(request.AgentId)
	//validate the exist is AgentId
	isExistUserAgent, err := t.userRepository.IsExist(request.AgentId)
	if !isNotAgent {
		return errors.New("Your Input Agent Id is The User Customer")
	}
	if isAgent {
		return errors.New("Your Input CustomerId is The User Agent")
	}
	if !isExistUserCustomer {
		return errors.New("Your CustomerId is not exist")
	}
	if !isExistUserAgent {
		return errors.New("Your AgentId is not exist")
	}
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
		DistrictId:        request.DistrictId,
		StatusTransaction: 0,
	}
	transaction, err = t.transactionsRepository.WithTrx(tx).Save(transaction)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	return tx.Commit().Error
}
