package controller

import (
    "github.com/itp-backend/backend-b-antar-jemput/models/web"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/itp-backend/backend-b-antar-jemput/service"
)

type CustomerController interface {
	RegisterCustomer(c *fiber.Ctx) error
}

type customerController struct {
	CustomerService service.CustomerService
}

//NewCustomerController -> returns new customer controller
func NewCustomerController(s service.CustomerService) CustomerController {
	return customerController{
		CustomerService: s,
	}
}

func (cs customerController) RegisterCustomer(c *fiber.Ctx) error {
	log.Print("[CustomerController]...add Customer")
	var customer web.RegisterCustomerRequest
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	err := cs.CustomerService.RegisterCustomer(customer)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Sukses Membuat Akun",
		"data": nil,
	})
}
