package command

import (
	"github.com/banggibima/be-itam/modules/positions/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type PositionCommandUsecase struct {
	Config          *config.Config
	PositionService *domain.PositionService
}

func NewPositionCommandUsecase(
	config *config.Config,
	positionService *domain.PositionService,
) *PositionCommandUsecase {
	return &PositionCommandUsecase{
		Config:          config,
		PositionService: positionService,
	}
}

func (u *PositionCommandUsecase) Create(dto *CreatePositionRequestDTO) (*CreatePositionResponseDTO, error) {
	position := &domain.Position{
		Name: &dto.Name,
	}

	if err := u.PositionService.Create(position); err != nil {
		return nil, err
	}

	response := &CreatePositionResponseDTO{
		ID:        position.ID,
		Name:      position.Name,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}

	return response, nil
}

func (u *PositionCommandUsecase) Update(dto *UpdatePositionRequestDTO) (*UpdatePositionResponseDTO, error) {
	position := &domain.Position{
		ID:   dto.ID,
		Name: &dto.Name,
	}

	if err := u.PositionService.Update(position); err != nil {
		return nil, err
	}

	response := &UpdatePositionResponseDTO{
		ID:        position.ID,
		Name:      position.Name,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}

	return response, nil
}

func (u *PositionCommandUsecase) UpdatePartial(dto *UpdatePartialPositionRequestDTO) (*UpdatePartialPositionResponseDTO, error) {
	position := &domain.Position{
		ID: dto.ID,
	}

	if dto.Name != nil {
		position.Name = dto.Name
	}

	if err := u.PositionService.UpdatePartial(position); err != nil {
		return nil, err
	}

	response := &UpdatePartialPositionResponseDTO{
		ID:        position.ID,
		Name:      position.Name,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}

	return response, nil
}

func (u *PositionCommandUsecase) Delete(dto *DeletePositionRequestDTO) (*DeletePositionResponseDTO, error) {
	position := &domain.Position{
		ID: dto.ID,
	}

	if err := u.PositionService.Delete(position); err != nil {
		return nil, err
	}

	response := &DeletePositionResponseDTO{
		ID: position.ID,
	}

	return response, nil
}
