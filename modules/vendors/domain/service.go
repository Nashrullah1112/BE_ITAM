package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type VendorService struct {
	Config                   *config.Config
	PostgresVendorRepository PostgresVendorRepository
}

func NewVendorService(
	config *config.Config,
	postgresVendorRepository PostgresVendorRepository,
) *VendorService {
	return &VendorService{
		Config:                   config,
		PostgresVendorRepository: postgresVendorRepository,
	}
}

func (s *VendorService) CountAll() (int, error) {
	return s.PostgresVendorRepository.CountAll()
}

func (s *VendorService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Vendor, error) {
	return s.PostgresVendorRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *VendorService) FindByID(id int) (*Vendor, error) {
	return s.PostgresVendorRepository.FindByID(id)
}

func (s *VendorService) Create(vendor *Vendor) error {
	return s.PostgresVendorRepository.Create(vendor)
}

func (s *VendorService) Update(vendor *Vendor) error {
	return s.PostgresVendorRepository.Update(vendor)
}

func (s *VendorService) UpdatePartial(vendor *Vendor) error {
	return s.PostgresVendorRepository.UpdatePartial(vendor)
}

func (s *VendorService) Delete(vendor *Vendor) error {
	return s.PostgresVendorRepository.Delete(vendor)
}
