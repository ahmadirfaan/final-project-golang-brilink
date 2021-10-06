package repositories

import (
	"fmt"
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user database.User, db *gorm.DB) (database.User, error)
}

type UserRepoImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepoImpl{}
}

func (u UserRepoImpl) Save(user database.User, db *gorm.DB) (database.User, error) {
	log.Println("Ini Error sebelum di save")
	err := db.Debug().Save(&user).Error
	fmt.Println(err)
	log.Printf("Users Repositories:%+v\n", user)
	return user, err
}
