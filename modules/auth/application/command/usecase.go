package command

import (
	"errors"
	"fmt"
	"time"

	userquery "github.com/banggibima/be-itam/modules/user/application/query"
	userdomain "github.com/banggibima/be-itam/modules/user/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/middleware"
	"github.com/banggibima/be-itam/pkg/utils"
)

type AuthCommandUsecase struct {
	Config      *config.Config
	UserService *userdomain.UserService
}

func NewAuthCommandUsecase(
	config *config.Config,
	userService *userdomain.UserService,
) *AuthCommandUsecase {
	return &AuthCommandUsecase{
		Config:      config,
		UserService: userService,
	}
}

func (u *AuthCommandUsecase) Register(dto *AuthRegisterRequestDTO) (*AuthRegisterResponseDTO, error) {
	role := "user"

	user := &userdomain.User{
		NIP:      &dto.NIP,
		Name:     &dto.Name,
		Email:    &dto.Email,
		Password: &dto.Password,
		Role:     &role,
	}

	fmt.Println(user)

	if dto.Password != dto.ConfirmPassword {
		return nil, errors.New("kata sandi tidak cocok")
	}

	hashed, err := utils.BcryptHashPassword(*user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = &hashed

	if err := u.UserService.Create(user); err != nil {
		return nil, err
	}

	token, err := middleware.EncodedAccess(&middleware.JWT{
		Secret:   u.Config.JWT.SecretAccess,
		Expire:   u.Config.JWT.ExpireAccess,
		Audience: u.Config.JWT.Audience,
		Issuer:   u.Config.JWT.Issuer,
	}, &userquery.UserResponseDTO{
		ID:   user.ID,
		Role: *user.Role,
	})

	expire := time.Now().Add(time.Duration(u.Config.JWT.ExpireAccess) * time.Second)

	response := &AuthRegisterResponseDTO{
		Token:  token.Raw,
		Expire: expire.UTC(),
		User: UserResponseDTO{
			ID:         user.ID,
			NIP:        *user.NIP,
			Name:       *user.Name,
			Email:      *user.Email,
			Role:       *user.Role,
			JoinDate:   user.JoinDate,
			DivisionID: user.DivisionID,
			PositionID: user.PositionID,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
	}

	return response, nil
}

func (u *AuthCommandUsecase) Login(dto *AuthLoginRequestDTO) (*AuthLoginResponseDTO, error) {
	user, err := u.UserService.FindByEmail(dto.Email)
	if err != nil {
		if err.Error() == "data tidak ditemukan" {
			return nil, errors.New("email atau kata sandi salah")
		}
		return nil, err
	}

	if err := utils.BcryptComparePassword(*user.Password, dto.Password); err != nil {
		if err.Error() == "kata sandi tidak cocok" {
			return nil, errors.New("email atau kata sandi salah")
		}
		return nil, err
	}

	token, err := middleware.EncodedAccess(&middleware.JWT{
		Secret:   u.Config.JWT.SecretAccess,
		Expire:   u.Config.JWT.ExpireAccess,
		Audience: u.Config.JWT.Audience,
		Issuer:   u.Config.JWT.Issuer,
	}, &userquery.UserResponseDTO{
		ID:   user.ID,
		Role: *user.Role,
	})

	expire := time.Now().Add(time.Duration(u.Config.JWT.ExpireAccess) * time.Second)

	response := &AuthLoginResponseDTO{
		Token:  token.Raw,
		Expire: expire.UTC(),
		User: UserResponseDTO{
			ID:         user.ID,
			NIP:        *user.NIP,
			Name:       *user.Name,
			Email:      *user.Email,
			Role:       *user.Role,
			JoinDate:   user.JoinDate,
			DivisionID: user.DivisionID,
			PositionID: user.PositionID,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
	}

	return response, nil
}
