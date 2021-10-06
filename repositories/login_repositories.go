package repositories

import (
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type LoginRepository interface {
	Save(login database.Login, db *gorm.DB) (database.Login, error)
}

type LoginRepoImpl struct {
}

func NewLoginRepository() LoginRepository {
	return &LoginRepoImpl{}
}

func (l LoginRepoImpl) Save(login database.Login, db *gorm.DB) (database.Login, error) {
	err := db.Debug().Create(&login).Error
	log.Printf("Login:%+v\n", login)
	return login, err
}
