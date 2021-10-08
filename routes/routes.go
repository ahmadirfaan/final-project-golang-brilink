package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "os"
)

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code":    fiber.StatusNotFound,
				"message": "Not Found Route",
				"data":    nil,
			})
		},
	)
}

func LoggerRoute(app *fiber.App) {
    app.Use(logger.New(logger.Config{
        Format:"${pid} ${status} - ${method} ${path}\n",
        Output:     os.Stdout,
    }))
}

