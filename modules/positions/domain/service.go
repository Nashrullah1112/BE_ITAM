package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type PositionService struct {
	Config                     *config.Config
	PostgresPositionRepository PostgresPositionRepository
}

func NewPositionService(
	config *config.Config,
	postgresPositionRepository PostgresPositionRepository,
) *PositionService {
	return &PositionService{
		Config:                     config,
		PostgresPositionRepository: postgresPositionRepository,
	}
}

func (s *PositionService) CountAll() (int, error) {
	return s.PostgresPositionRepository.CountAll()
}

func (s *PositionService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Position, error) {
	return s.PostgresPositionRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *PositionService) FindByID(id int) (*Position, error) {
	return s.PostgresPositionRepository.FindByID(id)
}

func (s *PositionService) Create(position *Position) error {
	return s.PostgresPositionRepository.Create(position)
}

func (s *PositionService) Update(position *Position) error {
	return s.PostgresPositionRepository.Update(position)
}

func (s *PositionService) UpdatePartial(position *Position) error {
	return s.PostgresPositionRepository.UpdatePartial(position)
}

func (s *PositionService) Delete(position *Position) error {
	return s.PostgresPositionRepository.Delete(position)
}
