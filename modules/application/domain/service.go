package domain

import (
	"github.com/banggibima/be-itam/pkg/config"
)

type ApplicationService struct {
	Config                        *config.Config
	PostgresApplicationRepository PostgresApplicationRepository
}

func NewApplicationService(
	config *config.Config,
	postgresApplicationRepository PostgresApplicationRepository,
) *ApplicationService {
	return &ApplicationService{
		Config:                        config,
		PostgresApplicationRepository: postgresApplicationRepository,
	}
}

func (s *ApplicationService) CountAll() (int, error) {
	return s.PostgresApplicationRepository.CountAll()
}

func (s *ApplicationService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Application, error) {
	return s.PostgresApplicationRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *ApplicationService) FindByID(id int) (*Application, error) {
	return s.PostgresApplicationRepository.FindByID(id)
}

func (s *ApplicationService) Create(application *Application) error {
	return s.PostgresApplicationRepository.Create(application)
}

func (s *ApplicationService) Update(application *Application) error {
	return s.PostgresApplicationRepository.Update(application)
}

func (s *ApplicationService) UpdatePartial(application *Application) error {
	return s.PostgresApplicationRepository.UpdatePartial(application)
}

func (s *ApplicationService) Delete(application *Application) error {
	return s.PostgresApplicationRepository.Delete(application)
}
