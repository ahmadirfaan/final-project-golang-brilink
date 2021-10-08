package controller

import (
	"log"

	"github.com/gofiber/fiber/v2"
	web "github.com/itp-backend/backend-b-antar-jemput/models/web/agent"
	"github.com/itp-backend/backend-b-antar-jemput/service"
)

type AgentController interface {
	RegisterAgent(c *fiber.Ctx) error
}

type agentController struct {
	AgentService service.AgentService
}

func NewAgentController(s service.AgentService) AgentController {
	return agentController{
		AgentService: s,
	}
}

func (as agentController) RegisterAgent(c *fiber.Ctx) error {
	log.Print("[AgentController]...add Agent")
	var agent web.RegisterAgentRequest
	if err := c.BodyParser(&agent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	err := as.AgentService.RegisterAgent(agent)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"code":    fiber.StatusCreated,
		"message": " Sukses Membuat Akun",
		"data": fiber.Map{
			"token": agent.Username,
		},
	})
}
