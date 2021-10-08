package service

import (
	"fmt"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	web "github.com/itp-backend/backend-b-antar-jemput/models/web/agent"
	"github.com/itp-backend/backend-b-antar-jemput/repositories"
	"github.com/itp-backend/backend-b-antar-jemput/utils"
	"gorm.io/gorm"
)

type AgentService interface {
	RegisterAgent(request web.RegisterAgentRequest) error
}

type agentService struct {
	agentRepository repositories.AgentRepository
	userRepository  repositories.UserRepository
	DB              *gorm.DB
}

func NewAgentService(ar repositories.AgentRepository, ur repositories.UserRepository, db *gorm.DB) AgentService {
	return &agentService{
		agentRepository: ar,
		userRepository:  ur,
		DB:              db,
	}
}

func (a *agentService) RegisterAgent(request web.RegisterAgentRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	agent := database.Agent{
		AgentName:   request.AgentName,
		NoHandphone: request.NoHandphone,
		Address:     request.Address,
		DistrictId:  request.DistrictId,
	}
	agent, err = a.agentRepository.WithTrx(tx).Save(agent)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	user := database.User{
		RoleId:   1,
		AgentId:  &agent.Id,
		Username: request.Username,
		Password: request.Password,
	}
	user, err = a.userRepository.WithTrx(tx).Save(user)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}