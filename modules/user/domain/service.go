package domain

import "github.com/banggibima/be-itam/pkg/config"

type UserService struct {
	Config                 *config.Config
	PostgresUserRepository PostgresUserRepository
}

func NewUserService(
	config *config.Config,
	postgresUserRepository PostgresUserRepository,
) *UserService {
	return &UserService{
		Config:                 config,
		PostgresUserRepository: postgresUserRepository,
	}
}

func (s *UserService) CountAll() (int, error) {
	return s.PostgresUserRepository.CountAll()
}

func (s *UserService) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*User, error) {
	return s.PostgresUserRepository.FindAll(offset, limit, sort, order, filters)
}

func (s *UserService) FindByID(id int) (*User, error) {
	return s.PostgresUserRepository.FindByID(id)
}

func (s *UserService) FindByEmail(email string) (*User, error) {
	return s.PostgresUserRepository.FindByEmail(email)
}

func (s *UserService) Create(user *User) error {
	return s.PostgresUserRepository.Create(user)
}

func (s *UserService) Update(user *User) error {
	return s.PostgresUserRepository.Update(user)
}

func (s *UserService) UpdatePartial(user *User) error {
	return s.PostgresUserRepository.UpdatePartial(user)
}

func (s *UserService) Delete(user *User) error {
	return s.PostgresUserRepository.Delete(user)
}
