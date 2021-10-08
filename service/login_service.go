package service

import (
    "github.com/itp-backend/backend-b-antar-jemput/models/database"
    "github.com/itp-backend/backend-b-antar-jemput/models/web"
    "github.com/itp-backend/backend-b-antar-jemput/repositories"
    "github.com/itp-backend/backend-b-antar-jemput/utils"
)

type LoginService interface {
    Login(request web.LoginRequest) (database.User, error)
}

type loginService struct {
    UserRepository repositories.UserRepository
}

func NewLoginService(ur repositories.UserRepository) LoginService {
    return &loginService{
        UserRepository: ur,
    }
}

func (l *loginService) Login(request web.LoginRequest) (database.User, error) {
    err := utils.NewValidator().Struct(&request)
    if err != nil {
        return database.User{},err
    }
    user, err := l.UserRepository.CheckUsernameAndPassword(request.Username, request.Password, request.Role)
    return user, err
}
