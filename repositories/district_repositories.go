package repositories

import (
	"fmt"
	"log"

	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	FindByRegencyId(regencyId string) ([]database.Districts, error)
}

type districtRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtRepository{
		DB: db,
	}
}

func (u districtRepository) FindByRegencyId(regencyId string) ([]database.Districts, error) {
	log.Println("Ini Error sebelum di save")
	var district []database.Districts
	err := u.DB.Debug().Where("regency_id = ?", regencyId).Find(&district).Error
	fmt.Println(err)
	log.Printf("District Repositories:%+v\n", district)
	return district, err
}
