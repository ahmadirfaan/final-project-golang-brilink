package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    jwtware "github.com/gofiber/jwt/v3"
    application "github.com/itp-backend/backend-b-antar-jemput/app"
    "os"
)

func LoggerRoute(app *fiber.App) {
    app.Use(logger.New(logger.Config{
        Format:"${pid} ${status} - ${method} ${path}\n",
        Output:     os.Stdout,
    }))
}

// Config for authorization with JWT
func MiddlewareAuth(app *fiber.App) {
    envApp := application.Init()
    app.Use(jwtware.New(jwtware.Config{
        SigningKey: []byte(envApp.Config.JWTSecret),
        ErrorHandler: func(ctx *fiber.Ctx, err error) error {
            ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "code": fiber.StatusUnauthorized,
                "message": "Missing or malformed JWT",
                "data": nil,
            })
            return nil
        },
    }))
}
