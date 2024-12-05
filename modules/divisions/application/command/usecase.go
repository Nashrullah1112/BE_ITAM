package command

import (
	"github.com/banggibima/be-itam/modules/divisions/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type DivisionCommandUsecase struct {
	Config          *config.Config
	DivisionService *domain.DivisionService
}

func NewDivisionCommandUsecase(
	config *config.Config,
	divisionService *domain.DivisionService,
) *DivisionCommandUsecase {
	return &DivisionCommandUsecase{
		Config:          config,
		DivisionService: divisionService,
	}
}

func (u *DivisionCommandUsecase) Create(dto *CreateDivisionRequestDTO) (*CreateDivisionResponseDTO, error) {
	division := &domain.Division{
		Name: &dto.Name,
	}

	if err := u.DivisionService.Create(division); err != nil {
		return nil, err
	}

	response := &CreateDivisionResponseDTO{
		ID:        division.ID,
		Name:      division.Name,
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
	}

	return response, nil
}

func (u *DivisionCommandUsecase) Update(dto *UpdateDivisionRequestDTO) (*UpdateDivisionResponseDTO, error) {
	division := &domain.Division{
		ID:   dto.ID,
		Name: &dto.Name,
	}

	if err := u.DivisionService.Update(division); err != nil {
		return nil, err
	}

	response := &UpdateDivisionResponseDTO{
		ID:        division.ID,
		Name:      division.Name,
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
	}

	return response, nil
}

func (u *DivisionCommandUsecase) UpdatePartial(dto *UpdatePartialDivisionRequestDTO) (*UpdatePartialDivisionResponseDTO, error) {
	division := &domain.Division{
		ID: dto.ID,
	}

	if dto.Name != nil {
		division.Name = dto.Name
	}

	if err := u.DivisionService.UpdatePartial(division); err != nil {
		return nil, err
	}

	response := &UpdatePartialDivisionResponseDTO{
		ID:        division.ID,
		Name:      division.Name,
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
	}

	return response, nil
}

func (u *DivisionCommandUsecase) Delete(dto *DeleteDivisionRequestDTO) (*DeleteDivisionResponseDTO, error) {
	division := &domain.Division{
		ID: dto.ID,
	}

	if err := u.DivisionService.Delete(division); err != nil {
		return nil, err
	}

	response := &DeleteDivisionResponseDTO{
		ID: division.ID,
	}

	return response, nil
}
