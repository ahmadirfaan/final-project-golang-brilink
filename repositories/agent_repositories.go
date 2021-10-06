package repositories

import (
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type AgentRepository interface {
	Save(agent database.Agent, db *gorm.DB) (database.Agent, error)
}

type AgentRepoImpl struct {
}

func NewAgentRepository() AgentRepository {
	return &AgentRepoImpl{}
}

func (a AgentRepoImpl) Save(agent database.Agent, db *gorm.DB) (database.Agent, error) {
	err := db.Debug().Create(&agent).Error
	log.Printf("Agent:%+v\n", agent)
	return agent, err
}
