package query

import (
	"github.com/banggibima/be-itam/modules/applications/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type ApplicationQueryUsecase struct {
	Config             *config.Config
	ApplicationService *domain.ApplicationService
}

func NewApplicationQueryUsecase(
	config *config.Config,
	applicationService *domain.ApplicationService,
) *ApplicationQueryUsecase {
	return &ApplicationQueryUsecase{
		Config:             config,
		ApplicationService: applicationService,
	}
}

func (u *ApplicationQueryUsecase) CountAll() (int, error) {
	response, err := u.ApplicationService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *ApplicationQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*ApplicationResponseDTO, error) {
	applications, err := u.ApplicationService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*ApplicationResponseDTO, 0)
	for _, application := range applications {
		response = append(response, &ApplicationResponseDTO{
			ID:                       application.ID,
			ApplicationName:          application.ApplicationName,
			CreationDate:             application.CreationDate,
			AcceptanceDate:           application.AcceptanceDate,
			StorageServerLocation:    application.StorageServerLocation,
			ApplicationType:          application.ApplicationType,
			ApplicationLink:          application.ApplicationLink,
			ApplicationCertification: application.ApplicationCertification,
			ActivationDate:           application.ActivationDate,
			ExpirationDate:           application.ExpirationDate,
			AssetID:                  application.AssetID,
			CreatedAt:                application.CreatedAt,
			UpdatedAt:                application.UpdatedAt,
		})
	}

	return response, nil
}

func (u *ApplicationQueryUsecase) FindByID(id int) (*ApplicationResponseDTO, error) {
	application, err := u.ApplicationService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &ApplicationResponseDTO{
		ID:                       application.ID,
		ApplicationName:          application.ApplicationName,
		CreationDate:             application.CreationDate,
		AcceptanceDate:           application.AcceptanceDate,
		StorageServerLocation:    application.StorageServerLocation,
		ApplicationType:          application.ApplicationType,
		ApplicationLink:          application.ApplicationLink,
		ApplicationCertification: application.ApplicationCertification,
		ActivationDate:           application.ActivationDate,
		ExpirationDate:           application.ExpirationDate,
		AssetID:                  application.AssetID,
		CreatedAt:                application.CreatedAt,
		UpdatedAt:                application.UpdatedAt,
	}

	return response, nil
}
