package service

import (
    "errors"
    "github.com/itp-backend/backend-b-antar-jemput/models/database"
    "github.com/itp-backend/backend-b-antar-jemput/models/web"
    "github.com/itp-backend/backend-b-antar-jemput/repositories"
    "github.com/itp-backend/backend-b-antar-jemput/utils"
    "gorm.io/gorm"
    "strconv"
)

type TransactionService interface {
	CreateTransaction(request web.CreateTransactionRequest, customerId string) error
	GetAllTransactionByUserId(userId string) ([]database.Transactions, error)
    IsUserAgent(userId string) (*bool, error)
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

func (t *transactionService) CreateTransaction(request web.CreateTransactionRequest, customerId string) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
    err = t.validateUserForTransaction(request, customerId)
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
    customerIdUint,_ := strconv.ParseUint(customerId, 10, 32)
	transaction := database.Transactions{
		TransactionTypeId: request.TransactionTypeId,
		CustomerId:        uint(customerIdUint),
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

func (t *transactionService) IsUserAgent(userId string) (*bool, error) {
    customerId,err := strconv.ParseUint(userId, 10, 32)
    if err != nil {
        return nil, errors.New("Error handling for converstion customerId")
    }
    isAgent, err := t.userRepository.IsUserAgent(uint(customerId))
    return isAgent, err
}


func (t *transactionService) validateUserForTransaction(request web.CreateTransactionRequest, customerId string) error {
    //validate userId must not same with agentId
    customerIdUint,_ := strconv.ParseUint(customerId, 10, 32)
    if strconv.FormatUint(uint64(request.AgentId), 32) == customerId {
        return errors.New("Agent Id and Customer Id Must not the same")
    }
    //validate exist districtId
    _, err := t.districtRepository.FindById(request.DistrictId)
    if err != nil {
        return err
    }
    //validate exist transactionType
    err = t.transactionTypeRepository.FindById(request.TransactionTypeId)
    if err != nil {
        return err
    }
    //validate if the request customerId is customer
    isAgent, err := t.IsUserAgent(customerId)
    //validate the exist is customerId
    isExistUserCustomer, err := t.userRepository.IsExist(uint(customerIdUint))
    //validate if the request agentId is agent
    isNotAgent, err := t.IsUserAgent(strconv.FormatUint(uint64(request.AgentId), 32))
    //validate the exist is AgentId
    isExistUserAgent, err := t.userRepository.IsExist(request.AgentId)
    if !*isNotAgent {
        return errors.New("Your Input Agent Id is The User Customer")
    }
    if *isAgent {
        return errors.New("Your Input CustomerId is The User Agent")
    }
    if !isExistUserCustomer {
        return errors.New("Your CustomerId is not exist")
    }
    if !isExistUserAgent {
        return errors.New("Your AgentId is not exist")
    }
    return err
}
