package repositories

import (
	"fmt"
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer database.Customer) (database.Customer, error)
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (c customerRepository) Save(customer database.Customer) (database.Customer, error) {
	err := c.DB.Debug().Create(&customer).Error
	log.Printf("Customer Repositories:%+v\n ", customer)
	fmt.Println(err)
	return customer, err
}
