package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type DivisionService struct {
	Config                     *config.Config
	PostgresDivisionRepository PostgresDivisionRepository
}

func NewDivisionService(
	config *config.Config,
	postgresDivisionRepository PostgresDivisionRepository,
) *DivisionService {
	return &DivisionService{
		Config:                     config,
		PostgresDivisionRepository: postgresDivisionRepository,
	}
}

func (s *DivisionService) CountAll() (int, error) {
	return s.PostgresDivisionRepository.CountAll()
}

func (s *DivisionService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Division, error) {
	return s.PostgresDivisionRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *DivisionService) FindByID(id int) (*Division, error) {
	return s.PostgresDivisionRepository.FindByID(id)
}

func (s *DivisionService) Create(division *Division) error {
	return s.PostgresDivisionRepository.Create(division)
}

func (s *DivisionService) Update(division *Division) error {
	return s.PostgresDivisionRepository.Update(division)
}

func (s *DivisionService) UpdatePartial(division *Division) error {
	return s.PostgresDivisionRepository.UpdatePartial(division)
}

func (s *DivisionService) Delete(division *Division) error {
	return s.PostgresDivisionRepository.Delete(division)
}
