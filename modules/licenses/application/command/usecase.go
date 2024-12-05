package command

import (
	"github.com/banggibima/be-itam/modules/licenses/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type LicenseCommandUsecase struct {
	Config         *config.Config
	LicenseService *domain.LicenseService
}

func NewLicenseCommandUsecase(
	config *config.Config,
	licenseService *domain.LicenseService,
) *LicenseCommandUsecase {
	return &LicenseCommandUsecase{
		Config:         config,
		LicenseService: licenseService,
	}
}

func (u *LicenseCommandUsecase) Create(dto *CreateLicenseRequestDTO) (*CreateLicenseResponseDTO, error) {
	license := &domain.License{
		PurchaseDate:        &dto.PurchaseDate,
		InstalledDeviceSN:   &dto.InstalledDeviceSN,
		ActivationDate:      &dto.ActivationDate,
		ExpirationDate:      &dto.ExpirationDate,
		AssetOwnershipType:  &dto.AssetOwnershipType,
		LicenseCategory:     &dto.LicenseCategory,
		LicenseVersion:      &dto.LicenseVersion,
		MaxApplicationUsers: &dto.MaxApplicationUsers,
		MaxDeviceLicenses:   &dto.MaxDeviceLicenses,
		LicenseType:         &dto.LicenseType,
	}

	if err := u.LicenseService.Create(license); err != nil {
		return nil, err
	}

	response := &CreateLicenseResponseDTO{
		ID:                  license.ID,
		PurchaseDate:        license.PurchaseDate,
		InstalledDeviceSN:   license.InstalledDeviceSN,
		ActivationDate:      license.ActivationDate,
		ExpirationDate:      license.ExpirationDate,
		AssetOwnershipType:  license.AssetOwnershipType,
		LicenseCategory:     license.LicenseCategory,
		LicenseVersion:      license.LicenseVersion,
		MaxApplicationUsers: license.MaxApplicationUsers,
		MaxDeviceLicenses:   license.MaxDeviceLicenses,
		LicenseType:         license.LicenseType,
		CreatedAt:           license.CreatedAt,
		UpdatedAt:           license.UpdatedAt,
	}

	return response, nil
}

func (u *LicenseCommandUsecase) Update(dto *UpdateLicenseRequestDTO) (*UpdateLicenseResponseDTO, error) {
	license := &domain.License{
		ID:                  dto.ID,
		PurchaseDate:        &dto.PurchaseDate,
		InstalledDeviceSN:   &dto.InstalledDeviceSN,
		ActivationDate:      &dto.ActivationDate,
		ExpirationDate:      &dto.ExpirationDate,
		AssetOwnershipType:  &dto.AssetOwnershipType,
		LicenseCategory:     &dto.LicenseCategory,
		LicenseVersion:      &dto.LicenseVersion,
		MaxApplicationUsers: &dto.MaxApplicationUsers,
		MaxDeviceLicenses:   &dto.MaxDeviceLicenses,
		LicenseType:         &dto.LicenseType,
	}

	if err := u.LicenseService.Update(license); err != nil {
		return nil, err
	}

	response := &UpdateLicenseResponseDTO{
		ID:                  license.ID,
		PurchaseDate:        license.PurchaseDate,
		InstalledDeviceSN:   license.InstalledDeviceSN,
		ActivationDate:      license.ActivationDate,
		ExpirationDate:      license.ExpirationDate,
		AssetOwnershipType:  license.AssetOwnershipType,
		LicenseCategory:     license.LicenseCategory,
		LicenseVersion:      license.LicenseVersion,
		MaxApplicationUsers: license.MaxApplicationUsers,
		MaxDeviceLicenses:   license.MaxDeviceLicenses,
		LicenseType:         license.LicenseType,
		CreatedAt:           license.CreatedAt,
		UpdatedAt:           license.UpdatedAt,
	}

	return response, nil
}

func (u *LicenseCommandUsecase) UpdatePartial(dto *UpdatePartialLicenseRequestDTO) (*UpdatePartialLicenseResponseDTO, error) {
	license := &domain.License{
		ID: dto.ID,
	}

	if dto.PurchaseDate != nil {
		license.PurchaseDate = dto.PurchaseDate
	}

	if dto.InstalledDeviceSN != nil {
		license.InstalledDeviceSN = dto.InstalledDeviceSN
	}

	if dto.ActivationDate != nil {
		license.ActivationDate = dto.ActivationDate
	}

	if dto.ExpirationDate != nil {
		license.ExpirationDate = dto.ExpirationDate
	}

	if dto.AssetOwnershipType != nil {
		license.AssetOwnershipType = dto.AssetOwnershipType
	}

	if dto.LicenseCategory != nil {
		license.LicenseCategory = dto.LicenseCategory
	}

	if dto.LicenseVersion != nil {
		license.LicenseVersion = dto.LicenseVersion
	}

	if dto.MaxApplicationUsers != nil {
		license.MaxApplicationUsers = dto.MaxApplicationUsers
	}

	if dto.MaxDeviceLicenses != nil {
		license.MaxDeviceLicenses = dto.MaxDeviceLicenses
	}

	if dto.LicenseType != nil {
		license.LicenseType = dto.LicenseType
	}

	if err := u.LicenseService.UpdatePartial(license); err != nil {
		return nil, err
	}

	response := &UpdatePartialLicenseResponseDTO{
		ID:                  license.ID,
		PurchaseDate:        license.PurchaseDate,
		InstalledDeviceSN:   license.InstalledDeviceSN,
		ActivationDate:      license.ActivationDate,
		ExpirationDate:      license.ExpirationDate,
		AssetOwnershipType:  license.AssetOwnershipType,
		LicenseCategory:     license.LicenseCategory,
		LicenseVersion:      license.LicenseVersion,
		MaxApplicationUsers: license.MaxApplicationUsers,
		MaxDeviceLicenses:   license.MaxDeviceLicenses,
		LicenseType:         license.LicenseType,
		CreatedAt:           license.CreatedAt,
		UpdatedAt:           license.UpdatedAt,
	}

	return response, nil
}

func (u *LicenseCommandUsecase) Delete(dto *DeleteLicenseRequestDTO) (*DeleteLicenseResponseDTO, error) {
	license := &domain.License{
		ID: dto.ID,
	}

	if err := u.LicenseService.Delete(license); err != nil {
		return nil, err
	}

	response := &DeleteLicenseResponseDTO{
		ID: license.ID,
	}

	return response, nil
}
