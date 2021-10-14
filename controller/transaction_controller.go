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
    GetAllTransactionByUserId(c *fiber.Ctx) error
}

type transactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(s service.TransactionService) TransactionController {
	return transactionController{
		TransactionService: s,
	}
}

func (ts transactionController) GetAllTransactionByUserId(c *fiber.Ctx) error {
    userId, err := utils.ExtractToken(c)
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "code":    fiber.StatusInternalServerError,
            "message": "Internal Server Error",
            "data":    nil,
        })
    }
    transactions, err := ts.TransactionService.GetAllTransactionByUserId(userId)
    if err != nil {
      return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
          "code":    fiber.StatusBadRequest,
          "message": err.Error(),
          "data":    nil,
      })
    }
    if len(transactions) == 0 {
      return c.Status(fiber.StatusOK).JSON(fiber.Map{
          "code":    fiber.StatusOK,
          "message": "No Transaction Data",
          "data":    nil,
      })
    } else {
      return c.Status(fiber.StatusOK).JSON(fiber.Map{
          "code":    fiber.StatusOK,
          "message": transactions,
          "data":    nil,
      })
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
		} else {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "code": fiber.StatusBadRequest,
                "message": err.Error(),
                "data": nil,
            })
        }
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": "Transaksi Diterima ",
		"data": nil,
	})
}
