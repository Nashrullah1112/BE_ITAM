package command

import (
	"github.com/banggibima/be-itam/modules/assets/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type AssetCommandUsecase struct {
	Config       *config.Config
	AssetService *domain.AssetService
}

func NewAssetCommandUsecase(
	config *config.Config,
	assetService *domain.AssetService,
) *AssetCommandUsecase {
	return &AssetCommandUsecase{
		Config:       config,
		AssetService: assetService,
	}
}

func (u *AssetCommandUsecase) Create(dto *CreateAssetRequestDTO) (*CreateAssetResponseDTO, error) {
	asset := &domain.Asset{
		SerialNumber:  &dto.SerialNumber,
		Brand:         &dto.Brand,
		Model:         &dto.Model,
		ReceiptNumber: &dto.ReceiptNumber,
		Status:        &dto.Status,
		VendorID:      dto.VendorID,
	}

	if err := u.AssetService.Create(asset); err != nil {
		return nil, err
	}

	response := &CreateAssetResponseDTO{
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

func (u *AssetCommandUsecase) Update(dto *UpdateAssetRequestDTO) (*UpdateAssetResponseDTO, error) {
	asset := &domain.Asset{
		ID:            dto.ID,
		SerialNumber:  &dto.SerialNumber,
		Brand:         &dto.Brand,
		Model:         &dto.Model,
		ReceiptNumber: &dto.ReceiptNumber,
		Status:        &dto.Status,
		VendorID:      dto.VendorID,
	}

	if err := u.AssetService.Update(asset); err != nil {
		return nil, err
	}

	response := &UpdateAssetResponseDTO{
		ID:            asset.ID,
		SerialNumber:  asset.SerialNumber,
		Brand:         asset.Brand,
		Model:         asset.Model,
		ReceiptNumber: asset.ReceiptNumber,
		Status:        asset.Status,
		VendorID:      asset.VendorID,
	}

	return response, nil
}

func (u *AssetCommandUsecase) UpdatePartial(dto *UpdatePartialAssetRequestDTO) (*UpdatePartialAssetResponseDTO, error) {
	asset := &domain.Asset{
		ID: dto.ID,
	}

	if dto.SerialNumber != nil {
		asset.SerialNumber = dto.SerialNumber
	}

	if dto.Brand != nil {
		asset.Brand = dto.Brand
	}

	if dto.Model != nil {
		asset.Model = dto.Model
	}

	if dto.ReceiptNumber != nil {
		asset.ReceiptNumber = dto.ReceiptNumber
	}

	if dto.Status != nil {
		asset.Status = dto.Status
	}

	if dto.VendorID != nil {
		asset.VendorID = dto.VendorID
	}

	if err := u.AssetService.UpdatePartial(asset); err != nil {
		return nil, err
	}

	response := &UpdatePartialAssetResponseDTO{
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

func (u *AssetCommandUsecase) Delete(dto *DeleteAssetRequestDTO) (*DeleteAssetResponseDTO, error) {
	asset := &domain.Asset{
		ID: dto.ID,
	}

	if err := u.AssetService.Delete(asset); err != nil {
		return nil, err
	}

	response := &DeleteAssetResponseDTO{
		ID: asset.ID,
	}

	return response, nil
}
