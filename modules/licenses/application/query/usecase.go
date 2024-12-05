package query

import (
	"github.com/banggibima/be-itam/modules/licenses/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type LicenseQueryUsecase struct {
	Config         *config.Config
	LicenseService *domain.LicenseService
}

func NewLicenseQueryUsecase(
	config *config.Config,
	licenseService *domain.LicenseService,
) *LicenseQueryUsecase {
	return &LicenseQueryUsecase{
		Config:         config,
		LicenseService: licenseService,
	}
}

func (u *LicenseQueryUsecase) CountAll() (int, error) {
	response, err := u.LicenseService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *LicenseQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*LicenseResponseDTO, error) {
	licenses, err := u.LicenseService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*LicenseResponseDTO, 0)
	for _, license := range licenses {
		response = append(response, &LicenseResponseDTO{
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
			AssetID:             license.AssetID,
			CreatedAt:           license.CreatedAt,
			UpdatedAt:           license.UpdatedAt,
		})
	}

	return response, nil
}

func (u *LicenseQueryUsecase) FindByID(id int) (*LicenseResponseDTO, error) {
	license, err := u.LicenseService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &LicenseResponseDTO{
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
		AssetID:             license.AssetID,
		CreatedAt:           license.CreatedAt,
		UpdatedAt:           license.UpdatedAt,
	}

	return response, nil
}
