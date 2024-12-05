package query

import (
	"github.com/banggibima/be-itam/modules/divisions/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type DivisionQueryUsecase struct {
	Config          *config.Config
	DivisionService *domain.DivisionService
}

func NewDivisionQueryUsecase(
	config *config.Config,
	divisionService *domain.DivisionService,
) *DivisionQueryUsecase {
	return &DivisionQueryUsecase{
		Config:          config,
		DivisionService: divisionService,
	}
}

func (u *DivisionQueryUsecase) CountAll() (int, error) {
	response, err := u.DivisionService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *DivisionQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*DivisionResponseDTO, error) {
	divisions, err := u.DivisionService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*DivisionResponseDTO, 0)
	for _, division := range divisions {
		response = append(response, &DivisionResponseDTO{
			ID:        division.ID,
			Name:      division.Name,
			CreatedAt: division.CreatedAt,
			UpdatedAt: division.UpdatedAt,
		})
	}

	return response, nil
}

func (u *DivisionQueryUsecase) FindByID(id int) (*DivisionResponseDTO, error) {
	division, err := u.DivisionService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &DivisionResponseDTO{
		ID:        division.ID,
		Name:      division.Name,
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
	}

	return response, nil
}
