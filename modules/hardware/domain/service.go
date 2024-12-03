package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type HardwareService struct {
	Config                     *config.Config
	PostgresHardwareRepository PostgresHardwareRepository
}

func NewHardwareService(
	config *config.Config,
	postgresHardwareRepository PostgresHardwareRepository,
) *HardwareService {
	return &HardwareService{
		Config:                     config,
		PostgresHardwareRepository: postgresHardwareRepository,
	}
}

func (s *HardwareService) CountAll() (int, error) {
	return s.PostgresHardwareRepository.CountAll()
}

func (s *HardwareService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Hardware, error) {
	return s.PostgresHardwareRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *HardwareService) FindByID(id int) (*Hardware, error) {
	return s.PostgresHardwareRepository.FindByID(id)
}

func (s *HardwareService) Create(hardware *Hardware) error {
	return s.PostgresHardwareRepository.Create(hardware)
}

func (s *HardwareService) Update(hardware *Hardware) error {
	return s.PostgresHardwareRepository.Update(hardware)
}

func (s *HardwareService) UpdatePartial(hardware *Hardware) error {
	return s.PostgresHardwareRepository.UpdatePartial(hardware)
}

func (s *HardwareService) Delete(hardware *Hardware) error {
	return s.PostgresHardwareRepository.Delete(hardware)
}
