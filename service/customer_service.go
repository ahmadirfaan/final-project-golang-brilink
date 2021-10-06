package service

import (
	"fmt"
	"log"

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
	loginRepository    repositories.LoginRepository
	DB                 *gorm.DB
}

func NewCustomerService(cr repositories.CustomerRepository, ur repositories.UserRepository, lr repositories.LoginRepository, db *gorm.DB) CustomerService {
	return &customerService{
		customerRepository: cr,
		userRepository:     ur,
		loginRepository:    lr,
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
	customer, err = c.customerRepository.Save(customer, tx)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	// log.Println("Ini Harusnya Customer setelah di save: ", customer)
	user := database.User{
		RoleId:     2,
		CustomerId: &customer.Id,
	}
	log.Printf("Ini CustomerId sebelum di save: %d, roleId: %d", user.CustomerId, user.RoleId)
	user, err = c.userRepository.Save(user, tx)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	log.Println("Ini User setelah di save: ", user)
	login := database.Login{
		Username: request.Username,
		Password: request.Password,
		UserId:   user.Id,
	}
	log.Println("Ini Login sebelum di save: ", login)
	login, err = c.loginRepository.Save(login, tx)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	log.Println("Ini Login setelah di save: ", login)
	log.Println("Ini Harusnya Commit: ", request)
	return tx.Commit().Error
}
