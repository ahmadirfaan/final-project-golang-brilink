package service

import (
	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"github.com/itp-backend/backend-b-antar-jemput/repositories"
)

type LocationService interface {
	GetAllLocationProvince() ([]database.Provinces, error)
}

type locationService struct {
	ProvinceRepository repositories.ProvinceRepository
}

func NewLocationService(pr repositories.ProvinceRepository) LocationService {
	return &locationService{
		ProvinceRepository: pr,
	}
}

func (l *locationService) GetAllLocationProvince() ([]database.Provinces, error) {
	provinces, err := l.ProvinceRepository.GetAll()
	return provinces, err
}
