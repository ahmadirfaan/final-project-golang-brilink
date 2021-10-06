package repositories

import (
	"fmt"
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer database.Customer, db *gorm.DB) (database.Customer, error)
}

type CustomerRepoImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepoImpl{}
}

func (c CustomerRepoImpl) Save(customer database.Customer, db *gorm.DB) (database.Customer, error) {
	err := db.Debug().Create(&customer).Error
	log.Printf("Customer Repositories:%+v\n ", customer)
	fmt.Println(err)
	return customer, err
}
