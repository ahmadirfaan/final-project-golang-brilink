package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	web "github.com/itp-backend/backend-b-antar-jemput/models/web/customer"
	"github.com/itp-backend/backend-b-antar-jemput/service"
)

type CustomerController interface {
	RegisterCustomer(c *fiber.Ctx) error
}

type customerController struct {
	customerService service.CustomerService
}

//NewCustomerController -> returns new customer controller
func NewCustomerController(s service.CustomerService) CustomerController {
	return customerController{
		customerService: s,
	}
}

func (cs customerController) RegisterCustomer(c *fiber.Ctx) error {
	log.Print("[CustomerController]...add Customer")
	var customer web.RegisterCustomerRequest
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := cs.customerService.RegisterCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Sukses Membuat Akun",
	})
}
