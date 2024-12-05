package query

import (
	"github.com/banggibima/be-itam/modules/users/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type UserQueryUsecase struct {
	Config      *config.Config
	UserService *domain.UserService
}

func NewUserQueryUsecase(
	config *config.Config,
	userService *domain.UserService,
) *UserQueryUsecase {
	return &UserQueryUsecase{
		Config:      config,
		UserService: userService,
	}
}

func (u *UserQueryUsecase) CountAll() (int, error) {
	response, err := u.UserService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *UserQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*UserResponseDTO, error) {
	users, err := u.UserService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*UserResponseDTO, 0)
	for _, user := range users {
		response = append(response, &UserResponseDTO{
			ID:         user.ID,
			NIP:        user.NIP,
			Name:       user.Name,
			Email:      user.Email,
			Role:       user.Role,
			JoinDate:   user.JoinDate,
			DivisionID: user.DivisionID,
			PositionID: user.PositionID,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		})
	}

	return response, nil
}

func (u *UserQueryUsecase) FindByID(id int) (*UserResponseDTO, error) {
	user, err := u.UserService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &UserResponseDTO{
		ID:         user.ID,
		NIP:        user.NIP,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		JoinDate:   user.JoinDate,
		DivisionID: user.DivisionID,
		PositionID: user.PositionID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return response, nil
}

func (u *UserQueryUsecase) FindByEmail(email string) (*UserResponseDTO, error) {
	user, err := u.UserService.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	response := &UserResponseDTO{
		ID:         user.ID,
		NIP:        user.NIP,
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		JoinDate:   user.JoinDate,
		DivisionID: user.DivisionID,
		PositionID: user.PositionID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return response, nil
}
