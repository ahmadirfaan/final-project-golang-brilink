package service

import (
    "fmt"
    "github.com/itp-backend/backend-b-antar-jemput/models/database"
    web "github.com/itp-backend/backend-b-antar-jemput/models/web/customer"
    "github.com/itp-backend/backend-b-antar-jemput/repositories"
    "github.com/itp-backend/backend-b-antar-jemput/utils"
    "gorm.io/gorm"
)

type CustomerService interface {
	RegisterCustomer(request web.RegisterCustomerRequest) error
}

type customerService struct {
	customerRepository repositories.CustomerRepository
	userRepository     repositories.UserRepository
	DB                 *gorm.DB
}

func NewCustomerService(cr repositories.CustomerRepository, ur repositories.UserRepository, db *gorm.DB) CustomerService {
	return &customerService{
		customerRepository: cr,
		userRepository:     ur,
		DB:                 db,
	}
}

func (c *customerService) RegisterCustomer(request web.RegisterCustomerRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
	// log.Println("Ini Request dr Website: ", request)
	tx := c.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	// log.Println("Ini Harusnya Recover: ", request)

	if err := tx.Error; err != nil {
		return err
	}
	// log.Println("Ini Harusnya Tidak Error: ", request)
	customer := database.Customer{
		Name:        request.Nama,
		NoHandphone: request.NoHandphone,
	}
	// log.Println("Ini Customer sebelum di save: ", customer)
	customer, err = c.customerRepository.WithTrx(tx).Save(customer)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	// log.Println("Ini Harusnya Customer setelah di save: ", customer)
	user := database.User{
		RoleId:     2,
		CustomerId: &customer.Id,
        Username: request.Username,
        Password: request.Password,
	}
	//log.Printf("Ini CustomerId sebelum di save: %d, roleId: %d", user.CustomerId, user.RoleId)
	user, err = c.userRepository.WithTrx(tx).Save(user)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
    //log.Println("Ini Harusnya Commit: ", request)
	return tx.Commit().Error
}
