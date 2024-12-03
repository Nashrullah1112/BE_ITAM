package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type DeviceService struct {
	Config                   *config.Config
	PostgresDeviceRepository PostgresDeviceRepository
}

func NewDeviceService(
	config *config.Config,
	postgresDeviceRepository PostgresDeviceRepository,
) *DeviceService {
	return &DeviceService{
		Config:                   config,
		PostgresDeviceRepository: postgresDeviceRepository,
	}
}

func (s *DeviceService) CountAll() (int, error) {
	return s.PostgresDeviceRepository.CountAll()
}

func (s *DeviceService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Device, error) {
	return s.PostgresDeviceRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *DeviceService) FindByID(id int) (*Device, error) {
	return s.PostgresDeviceRepository.FindByID(id)
}

func (s *DeviceService) Create(device *Device) error {
	return s.PostgresDeviceRepository.Create(device)
}

func (s *DeviceService) Update(device *Device) error {
	return s.PostgresDeviceRepository.Update(device)
}

func (s *DeviceService) UpdatePartial(device *Device) error {
	return s.PostgresDeviceRepository.UpdatePartial(device)
}

func (s *DeviceService) Delete(device *Device) error {
	return s.PostgresDeviceRepository.Delete(device)
}
