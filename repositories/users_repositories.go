package repositories

import (
	"fmt"
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user database.User) (database.User, error)
    WithTrx(trxHandle *gorm.DB) userRepository
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u userRepository) Save(user database.User) (database.User, error) {
	log.Println("Ini Error sebelum di save")
	err := u.DB.Debug().Save(&user).Error
	fmt.Println(err)
	log.Printf("Users Repositories:%+v\n", user)
	return user, err
}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
    if trxHandle == nil {
        log.Print("Transaction Database not found")
        return u
    }
    u.DB = trxHandle
    return u
}
