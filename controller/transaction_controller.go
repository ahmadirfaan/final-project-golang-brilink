package controller

import (
	"errors"

	"github.com/go-playground/validator"
	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/itp-backend/backend-b-antar-jemput/models/web"
	"github.com/itp-backend/backend-b-antar-jemput/service"
	"github.com/itp-backend/backend-b-antar-jemput/utils"
)

type TransactionController interface {
	CreateTransaction(c *fiber.Ctx) error
}

type transactionController struct {
	TransactionService service.TransactionService
}

func NewTransacrtionController(s service.TransactionService) TransactionController {
	return transactionController{
		TransactionService: s,
	}
}

func (ts transactionController) CreateTransaction(c *fiber.Ctx) error {
	var transaction web.CreateTransactionRequest
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Error for handling your request",
			"data":    nil,
		})
	}

	err := ts.TransactionService.CreateTransaction(transaction)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &validator.ValidationErrors{}) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"code":    fiber.StatusBadRequest,
				"message": utils.ValidatorErrors(err),
				"data":    nil,
			})
		} else if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"code":    fiber.StatusConflict,
				"message": "Username Already is exist",
				"data":    nil,
			})
		}
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"code":    fiber.StatusAccepted,
		"message": " Transaksi Diterima ",
		"data":    fiber.Map{
            "transactionId":
        },

	})
}
