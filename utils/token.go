package utils

import (
    "github.com/golang-jwt/jwt/v4"
    "github.com/itp-backend/backend-b-antar-jemput/app"
    "github.com/itp-backend/backend-b-antar-jemput/models/database"
    "time"
)



// This function to generates Access Token
func GenerateToken(user database.User) (*string, error) {
    app := app.Init()
    jwtSecretKey := app.Config.JWTSecret

    // Generates Access Token Token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["sub"] = 1
    claims["roleId"] = user.Role.Id
    claims["roleName"] = user.Role.Role
    claims["userId"] = user.Id
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    t, err := token.SignedString([]byte(jwtSecretKey))
    if err != nil {
        return nil, err
    }
    return &t, err
}
