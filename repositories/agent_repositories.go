package repositories

import (
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type AgentRepository interface {
	Save(agent database.Agent) (database.Agent, error)
}

type agentRepo struct {
	DB *gorm.DB
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &agentRepo{
		DB: db,
	}
}

func (a agentRepo) Save(agent database.Agent) (database.Agent, error) {
	err := a.DB.Debug().Create(&agent).Error
	log.Printf("Agent:%+v\n", agent)
	return agent, err
}
