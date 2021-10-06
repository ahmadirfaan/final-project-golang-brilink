package repositories

import (
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type LoginRepository interface {
	Save(login database.Login) (database.Login, error)
    WithTrx(trxHandle *gorm.DB) loginRepository
}

type loginRepository struct {
	DB *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{
		DB: db,
	}
}

func (l loginRepository) Save(login database.Login) (database.Login, error) {
	err := l.DB.Debug().Create(&login).Error
	log.Printf("Login:%+v\n", login)
	return login, err
}

func (l loginRepository) WithTrx(trxHandle *gorm.DB) loginRepository {
    if trxHandle == nil {
        log.Print("Transaction Database not found")
        return l
    }
    l.DB = trxHandle
    return l
}
