package service

import (
	"github.com/itp-backend/backend-b-antar-jemput/models/database"
	"github.com/itp-backend/backend-b-antar-jemput/repositories"
)

type LocationService interface {
	GetAllLocationProvince() ([]database.Provinces, error)
    GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error)
}

type locationService struct {
	ProvinceRepository repositories.ProvinceRepository
    RegencyRepository repositories.RegencyRepository

}

func NewLocationService(pr repositories.ProvinceRepository, rr repositories.RegencyRepository) LocationService {
	return &locationService{
		ProvinceRepository: pr,
        RegencyRepository: rr,
	}
}

func (l *locationService) GetAllLocationProvince() ([]database.Provinces, error) {
	provinces, err := l.ProvinceRepository.GetAll()
	return provinces, err
}

func (l *locationService) GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error) {
    regencies, err := l.RegencyRepository.FindByProvinceId(provinceId)
    return regencies, err
}

