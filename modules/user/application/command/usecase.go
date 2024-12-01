package command

import (
	"errors"

	"github.com/banggibima/be-itam/modules/user/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/utils"
)

type UserCommandUsecase struct {
	Config      *config.Config
	UserService *domain.UserService
}

func NewUserCommandUsecase(
	config *config.Config,
	userService *domain.UserService,
) *UserCommandUsecase {
	return &UserCommandUsecase{
		Config:      config,
		UserService: userService,
	}
}

func (u *UserCommandUsecase) Create(dto *UserCreateRequestDTO) (*UserCreateResponseDTO, error) {
	user := &domain.User{
		NIP:        &dto.NIP,
		Name:       &dto.Name,
		Email:      &dto.Email,
		Password:   &dto.Password,
		Role:       &dto.Role,
		JoinDate:   &dto.JoinDate,
		DivisionID: dto.DivisionID,
		PositionID: dto.PositionID,
	}

	hashed, err := utils.BcryptHashPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	if err := utils.BcryptComparePassword(hashed, *user.Password); err != nil {
		return nil, errors.New("kata sandi tidak cocok")
	}

	user.Password = &hashed

	if err := u.UserService.Create(user); err != nil {
		return nil, err
	}

	response := &UserCreateResponseDTO{
		ID:         user.ID,
		NIP:        *user.NIP,
		Name:       *user.Name,
		Email:      *user.Email,
		Role:       *user.Role,
		JoinDate:   *user.JoinDate,
		DivisionID: user.DivisionID,
		PositionID: user.PositionID,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}

	return response, nil
}

func (u *UserCommandUsecase) Update(dto *UserUpdateRequestDTO) (*UserUpdateResponseDTO, error) {
	user := &domain.User{
		ID:         dto.ID,
		NIP:        &dto.NIP,
		Name:       &dto.Name,
		Email:      &dto.Email,
		Password:   &dto.Password,
		Role:       &dto.Role,
		JoinDate:   &dto.JoinDate,
		DivisionID: dto.DivisionID,
		PositionID: dto.PositionID,
	}

	hashed, err := utils.BcryptHashPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = &hashed

	if err := u.UserService.Update(user); err != nil {
		return nil, err
	}

	response := &UserUpdateResponseDTO{
		ID:         user.ID,
		NIP:        *user.NIP,
		Name:       *user.Name,
		Email:      *user.Email,
		Role:       *user.Role,
		JoinDate:   *user.JoinDate,
		DivisionID: user.DivisionID,
		PositionID: user.PositionID,
	}

	return response, nil
}

func (u *UserCommandUsecase) UpdatePartial(dto *UserUpdatePartialRequestDTO) (*UserUpdateResponseDTO, error) {
	user := &domain.User{
		ID: dto.ID,
	}

	if dto.NIP != nil {
		user.NIP = dto.NIP
	}

	if dto.Name != nil {
		user.Name = dto.Name
	}

	if dto.Email != nil {
		user.Email = dto.Email
	}

	if dto.Password != nil {
		hashed, err := utils.BcryptHashPassword(*dto.Password)
		if err != nil {
			return nil, err
		}

		user.Password = &hashed
	}

	if dto.Role != nil {
		user.Role = dto.Role
	}

	if dto.JoinDate != nil {
		user.JoinDate = dto.JoinDate
	}

	if dto.DivisionID != nil {
		user.DivisionID = dto.DivisionID
	}

	if dto.PositionID != nil {
		user.PositionID = dto.PositionID
	}

	if err := u.UserService.Update(user); err != nil {
		return nil, err
	}

	response := &UserUpdateResponseDTO{
		ID:         user.ID,
		NIP:        *user.NIP,
		Name:       *user.Name,
		Email:      *user.Email,
		Role:       *user.Role,
		JoinDate:   *user.JoinDate,
		DivisionID: user.DivisionID,
		PositionID: user.PositionID,
	}

	return response, nil
}

func (u *UserCommandUsecase) Delete(dto *UserDeleteRequestDTO) (*UserDeleteResponseDTO, error) {
	user := &domain.User{
		ID: dto.ID,
	}

	if err := u.UserService.Delete(user); err != nil {
		return nil, err
	}

	response := &UserDeleteResponseDTO{
		ID: user.ID,
	}

	return response, nil
}
