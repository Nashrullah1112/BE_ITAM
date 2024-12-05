package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type LicenseService struct {
	Config                    *config.Config
	PostgresLicenseRepository PostgresLicenseRepository
}

func NewLicenseService(
	config *config.Config,
	postgresLicenseRepository PostgresLicenseRepository,
) *LicenseService {
	return &LicenseService{
		Config:                    config,
		PostgresLicenseRepository: postgresLicenseRepository,
	}
}

func (s *LicenseService) CountAll() (int, error) {
	return s.PostgresLicenseRepository.CountAll()
}

func (s *LicenseService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*License, error) {
	return s.PostgresLicenseRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *LicenseService) FindByID(id int) (*License, error) {
	return s.PostgresLicenseRepository.FindByID(id)
}

func (s *LicenseService) Create(license *License) error {
	return s.PostgresLicenseRepository.Create(license)
}

func (s *LicenseService) Update(license *License) error {
	return s.PostgresLicenseRepository.Update(license)
}

func (s *LicenseService) UpdatePartial(license *License) error {
	return s.PostgresLicenseRepository.UpdatePartial(license)
}

func (s *LicenseService) Delete(license *License) error {
	return s.PostgresLicenseRepository.Delete(license)
}
