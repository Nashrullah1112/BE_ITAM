package command

import (
	"github.com/banggibima/be-itam/modules/applications/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type ApplicationCommandUsecase struct {
	Config             *config.Config
	ApplicationService *domain.ApplicationService
}

func NewApplicationCommandUsecase(
	config *config.Config,
	applicationService *domain.ApplicationService,
) *ApplicationCommandUsecase {
	return &ApplicationCommandUsecase{
		Config:             config,
		ApplicationService: applicationService,
	}
}

func (u *ApplicationCommandUsecase) Create(dto *CreateApplicationRequestDTO) (*CreateApplicationResponseDTO, error) {
	application := &domain.Application{
		ApplicationName:          &dto.ApplicationName,
		CreationDate:             &dto.CreationDate,
		AcceptanceDate:           &dto.AcceptanceDate,
		StorageServerLocation:    &dto.StorageServerLocation,
		ApplicationType:          &dto.ApplicationType,
		ApplicationLink:          &dto.ApplicationLink,
		ApplicationCertification: &dto.ApplicationCertification,
		ActivationDate:           &dto.ActivationDate,
		ExpirationDate:           &dto.ExpirationDate,
		AssetID:                  dto.AssetID,
	}

	if err := u.ApplicationService.Create(application); err != nil {
		return nil, err
	}

	response := &CreateApplicationResponseDTO{
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

func (u *ApplicationCommandUsecase) Update(dto *UpdateApplicationRequestDTO) (*UpdateApplicationResponseDTO, error) {
	application := &domain.Application{
		ID:                       dto.ID,
		ApplicationName:          &dto.ApplicationName,
		CreationDate:             &dto.CreationDate,
		AcceptanceDate:           &dto.AcceptanceDate,
		StorageServerLocation:    &dto.StorageServerLocation,
		ApplicationType:          &dto.ApplicationType,
		ApplicationLink:          &dto.ApplicationLink,
		ApplicationCertification: &dto.ApplicationCertification,
		ActivationDate:           &dto.ActivationDate,
		ExpirationDate:           &dto.ExpirationDate,
		AssetID:                  dto.AssetID,
	}

	if err := u.ApplicationService.Update(application); err != nil {
		return nil, err
	}

	response := &UpdateApplicationResponseDTO{
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

func (u *ApplicationCommandUsecase) UpdatePartial(dto *UpdatePartialApplicationRequestDTO) (*UpdatePartialApplicationResponseDTO, error) {
	application := &domain.Application{
		ID: dto.ID,
	}

	if dto.ApplicationName != nil {
		application.ApplicationName = dto.ApplicationName
	}

	if dto.CreationDate != nil {
		application.CreationDate = dto.CreationDate
	}

	if dto.AcceptanceDate != nil {
		application.AcceptanceDate = dto.AcceptanceDate
	}

	if dto.StorageServerLocation != nil {
		application.StorageServerLocation = dto.StorageServerLocation
	}

	if dto.ApplicationType != nil {
		application.ApplicationType = dto.ApplicationType
	}

	if dto.ApplicationLink != nil {
		application.ApplicationLink = dto.ApplicationLink
	}

	if dto.ApplicationCertification != nil {
		application.ApplicationCertification = dto.ApplicationCertification
	}

	if dto.ActivationDate != nil {
		application.ActivationDate = dto.ActivationDate
	}

	if dto.ExpirationDate != nil {
		application.ExpirationDate = dto.ExpirationDate
	}

	if dto.AssetID != nil {
		application.AssetID = dto.AssetID
	}

	if err := u.ApplicationService.UpdatePartial(application); err != nil {
		return nil, err
	}

	response := &UpdatePartialApplicationResponseDTO{
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

func (u *ApplicationCommandUsecase) Delete(dto *DeleteApplicationRequestDTO) (*DeleteApplicationResponseDTO, error) {
	application := &domain.Application{
		ID: dto.ID,
	}

	if err := u.ApplicationService.Delete(application); err != nil {
		return nil, err
	}

	response := &DeleteApplicationResponseDTO{
		ID: application.ID,
	}

	return response, nil
}
