package repositories

import (
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type AgentRepository interface {
	Save(agent database.Agent) (database.Agent, error)
	WithTrx(trxHandle *gorm.DB) agentRepo
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

func (a agentRepo) WithTrx(trxHandle *gorm.DB) agentRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return a
	}
	a.DB = trxHandle
	return a
}
