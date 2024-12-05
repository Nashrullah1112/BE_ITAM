package query

import (
	"github.com/banggibima/be-itam/modules/positions/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type PositionQueryUsecase struct {
	Config          *config.Config
	PositionService *domain.PositionService
}

func NewPositionQueryUsecase(
	config *config.Config,
	positionService *domain.PositionService,
) *PositionQueryUsecase {
	return &PositionQueryUsecase{
		Config:          config,
		PositionService: positionService,
	}
}

func (u *PositionQueryUsecase) CountAll() (int, error) {
	response, err := u.PositionService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *PositionQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*PositionResponseDTO, error) {
	positions, err := u.PositionService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*PositionResponseDTO, 0)
	for _, position := range positions {
		response = append(response, &PositionResponseDTO{
			ID:        position.ID,
			Name:      position.Name,
			CreatedAt: position.CreatedAt,
			UpdatedAt: position.UpdatedAt,
		})
	}

	return response, nil
}

func (u *PositionQueryUsecase) FindByID(id int) (*PositionResponseDTO, error) {
	position, err := u.PositionService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &PositionResponseDTO{
		ID:        position.ID,
		Name:      position.Name,
		CreatedAt: position.CreatedAt,
		UpdatedAt: position.UpdatedAt,
	}

	return response, nil
}
