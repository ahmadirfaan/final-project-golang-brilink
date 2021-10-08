package controller

import (
    "github.com/gofiber/fiber/v2"
    "github.com/itp-backend/backend-b-antar-jemput/models/web"
    "github.com/itp-backend/backend-b-antar-jemput/service"
    "log"
)

type LoginController interface {
    Login(c *fiber.Ctx) error
}

type loginController struct {
    LoginService service.LoginService
}

func NewLoginController(ls service.LoginService) LoginController {
    return loginController{
        LoginService: ls,
    }
}

func (cs loginController) Login(c *fiber.Ctx) error {
    var login web.LoginRequest
    if err := c.BodyParser(&login); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
    }

    user, err := cs.LoginService.Login(login)
    if err != nil || user.Id == 0 {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "code":    fiber.StatusBadRequest,
            "message": err.Error(),
            "data":    nil,
        })
    }
    log.Println("User: ", user)
    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "code":    fiber.StatusCreated,
        "message": nil,
        "data": fiber.Map{
            "role": user.Role.Role,
            "userId": user.Id,
        },
    })
}
