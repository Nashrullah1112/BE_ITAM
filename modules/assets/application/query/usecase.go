package query

import (
	"github.com/banggibima/be-itam/modules/assets/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type AssetQueryUsecase struct {
	Config       *config.Config
	AssetService *domain.AssetService
}

func NewAssetQueryUsecase(
	config *config.Config,
	assetService *domain.AssetService,
) *AssetQueryUsecase {
	return &AssetQueryUsecase{
		Config:       config,
		AssetService: assetService,
	}
}

func (u *AssetQueryUsecase) CountAll() (int, error) {
	response, err := u.AssetService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *AssetQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*AssetResponseDTO, error) {
	assets, err := u.AssetService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*AssetResponseDTO, 0)
	for _, asset := range assets {
		response = append(response, &AssetResponseDTO{
			ID:            asset.ID,
			SerialNumber:  asset.SerialNumber,
			Brand:         asset.Brand,
			Model:         asset.Model,
			ReceiptNumber: asset.ReceiptNumber,
			Status:        asset.Status,
			VendorID:      asset.VendorID,
			CreatedAt:     asset.CreatedAt,
			UpdatedAt:     asset.UpdatedAt,
		})
	}

	return response, nil
}

func (u *AssetQueryUsecase) FindByID(id int) (*AssetResponseDTO, error) {
	asset, err := u.AssetService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &AssetResponseDTO{
		ID:            asset.ID,
		SerialNumber:  asset.SerialNumber,
		Brand:         asset.Brand,
		Model:         asset.Model,
		ReceiptNumber: asset.ReceiptNumber,
		Status:        asset.Status,
		VendorID:      asset.VendorID,
		CreatedAt:     asset.CreatedAt,
		UpdatedAt:     asset.UpdatedAt,
	}

	return response, nil
}
