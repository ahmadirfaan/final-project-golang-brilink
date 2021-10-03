package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {

			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"code": 404,
				"msg":  "Not Found Route",
			})
		},
	)
}
