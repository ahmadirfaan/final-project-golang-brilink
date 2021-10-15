package repositories

import (
    "log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction database.Transactions) (database.Transactions, error)
	FindTransactionWithUserId(userId string) ([]database.Transactions, error)
	WithTrx(trxHandle *gorm.DB) transactionRepo
}

type transactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{
		DB: db,
	}
}

func (t transactionRepo) Save(transactions database.Transactions) (database.Transactions, error) {
	err := t.DB.Debug().Create(&transactions).Error
	log.Printf("Transaction:%+v\n", transactions)
	return transactions, err
}

func (t transactionRepo) FindTransactionWithUserId(userId string) ([]database.Transactions, error) {
	var transactionList []database.Transactions
	err := t.DB.Debug().Where("agent_id = ? or customer_id = ?", userId, userId).
        Preload("UserCustomer").Preload("UserCustomer.Customer").
        Preload("TransactionType").Preload("TransactionType.ServiceTypeTransaction").
        Preload("UserAgent").Preload("UserAgent.Agent").Find(&transactionList).Error
	return transactionList, err
}

func (t transactionRepo) WithTrx(trxHandle *gorm.DB) transactionRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not  found")
		return t
	}
	t.DB = trxHandle
	return t
}
