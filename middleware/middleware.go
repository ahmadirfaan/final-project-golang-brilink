package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "os"
)

func LoggerRoute(app *fiber.App) {
    app.Use(logger.New(logger.Config{
        Format:"${pid} ${status} - ${method} ${path}\n",
        Output:     os.Stdout,
    }))
}
