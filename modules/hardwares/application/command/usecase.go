package command

import (
	"github.com/banggibima/be-itam/modules/hardwares/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type HardwareCommandUsecase struct {
	Config          *config.Config
	HardwareService *domain.HardwareService
}

func NewHardwareCommandUsecase(
	config *config.Config,
	hardwareService *domain.HardwareService,
) *HardwareCommandUsecase {
	return &HardwareCommandUsecase{
		Config:          config,
		HardwareService: hardwareService,
	}
}

func (u *HardwareCommandUsecase) Create(dto *CreateHardwareRequestDTO) (*CreateHardwareResponseDTO, error) {
	hardware := &domain.Hardware{
		ReceiptDate:          &dto.ReceiptDate,
		ReceiptProof:         &dto.ReceiptProof,
		AssetType:            &dto.AssetType,
		DeviceActivationDate: &dto.DeviceActivationDate,
		InspectionResult:     &dto.InspectionResult,
		SerialNumber:         &dto.SerialNumber,
		Model:                &dto.Model,
		WarrantyStartDate:    &dto.WarrantyStartDate,
		WarrantyEndDate:      &dto.WarrantyEndDate,
		WarrantyCardNumber:   &dto.WarrantyCardNumber,
		DeviceSpecifications: &dto.DeviceSpecifications,
		AssetStatus:          &dto.AssetStatus,
		AssetResponsible:     &dto.AssetResponsible,
		StorageLocation:      &dto.StorageLocation,
		UsagePeriod:          &dto.UsagePeriod,
		AssetOutTime:         &dto.AssetOutTime,
		AssetCondition:       &dto.AssetCondition,
		PurchaseReceipt:      &dto.PurchaseReceipt,
		AssetID:              dto.AssetID,
		DivisionID:           dto.DivisionID,
	}

	if err := u.HardwareService.Create(hardware); err != nil {
		return nil, err
	}

	response := &CreateHardwareResponseDTO{
		ID:                   hardware.ID,
		ReceiptDate:          hardware.ReceiptDate,
		ReceiptProof:         hardware.ReceiptProof,
		AssetType:            hardware.AssetType,
		DeviceActivationDate: hardware.DeviceActivationDate,
		InspectionResult:     hardware.InspectionResult,
		SerialNumber:         hardware.SerialNumber,
		Model:                hardware.Model,
		WarrantyStartDate:    hardware.WarrantyStartDate,
		WarrantyEndDate:      hardware.WarrantyEndDate,
		WarrantyCardNumber:   hardware.WarrantyCardNumber,
		DeviceSpecifications: hardware.DeviceSpecifications,
		AssetStatus:          hardware.AssetStatus,
		AssetResponsible:     hardware.AssetResponsible,
		StorageLocation:      hardware.StorageLocation,
		UsagePeriod:          hardware.UsagePeriod,
		AssetOutTime:         hardware.AssetOutTime,
		AssetCondition:       hardware.AssetCondition,
		PurchaseReceipt:      hardware.PurchaseReceipt,
		AssetID:              hardware.AssetID,
		DivisionID:           hardware.DivisionID,
		CreatedAt:            hardware.CreatedAt,
		UpdatedAt:            hardware.UpdatedAt,
	}

	return response, nil
}

func (u *HardwareCommandUsecase) Update(dto *UpdateHardwareRequestDTO) (*UpdateHardwareResponseDTO, error) {
	hardware := &domain.Hardware{
		ID:                   dto.ID,
		ReceiptDate:          &dto.ReceiptDate,
		ReceiptProof:         &dto.ReceiptProof,
		AssetType:            &dto.AssetType,
		DeviceActivationDate: &dto.DeviceActivationDate,
		InspectionResult:     &dto.InspectionResult,
		SerialNumber:         &dto.SerialNumber,
		Model:                &dto.Model,
		WarrantyStartDate:    &dto.WarrantyStartDate,
		WarrantyEndDate:      &dto.WarrantyEndDate,
		WarrantyCardNumber:   &dto.WarrantyCardNumber,
		DeviceSpecifications: &dto.DeviceSpecifications,
		AssetStatus:          &dto.AssetStatus,
		AssetResponsible:     &dto.AssetResponsible,
		StorageLocation:      &dto.StorageLocation,
		UsagePeriod:          &dto.UsagePeriod,
		AssetOutTime:         &dto.AssetOutTime,
		AssetCondition:       &dto.AssetCondition,
		PurchaseReceipt:      &dto.PurchaseReceipt,
		AssetID:              dto.AssetID,
		DivisionID:           dto.DivisionID,
	}

	if err := u.HardwareService.Update(hardware); err != nil {
		return nil, err
	}

	response := &UpdateHardwareResponseDTO{
		ID:                   hardware.ID,
		ReceiptDate:          hardware.ReceiptDate,
		ReceiptProof:         hardware.ReceiptProof,
		AssetType:            hardware.AssetType,
		DeviceActivationDate: hardware.DeviceActivationDate,
		InspectionResult:     hardware.InspectionResult,
		SerialNumber:         hardware.SerialNumber,
		Model:                hardware.Model,
		WarrantyStartDate:    hardware.WarrantyStartDate,
		WarrantyEndDate:      hardware.WarrantyEndDate,
		WarrantyCardNumber:   hardware.WarrantyCardNumber,
		DeviceSpecifications: hardware.DeviceSpecifications,
		AssetStatus:          hardware.AssetStatus,
		AssetResponsible:     hardware.AssetResponsible,
		StorageLocation:      hardware.StorageLocation,
		UsagePeriod:          hardware.UsagePeriod,
		AssetOutTime:         hardware.AssetOutTime,
		AssetCondition:       hardware.AssetCondition,
		PurchaseReceipt:      hardware.PurchaseReceipt,
		AssetID:              hardware.AssetID,
		DivisionID:           hardware.DivisionID,
		CreatedAt:            hardware.CreatedAt,
		UpdatedAt:            hardware.UpdatedAt,
	}

	return response, nil
}

func (u *HardwareCommandUsecase) UpdatePartial(dto *UpdatePartialHardwareRequestDTO) (*UpdatePartialHardwareResponseDTO, error) {
	hardware := &domain.Hardware{
		ID: dto.ID,
	}

	if dto.ReceiptDate != nil {
		hardware.ReceiptDate = dto.ReceiptDate
	}

	if dto.ReceiptProof != nil {
		hardware.ReceiptProof = dto.ReceiptProof
	}

	if dto.AssetType != nil {
		hardware.AssetType = dto.AssetType
	}

	if dto.DeviceActivationDate != nil {
		hardware.DeviceActivationDate = dto.DeviceActivationDate
	}

	if dto.InspectionResult != nil {
		hardware.InspectionResult = dto.InspectionResult
	}

	if dto.SerialNumber != nil {
		hardware.SerialNumber = dto.SerialNumber
	}

	if dto.Model != nil {
		hardware.Model = dto.Model
	}

	if dto.WarrantyStartDate != nil {
		hardware.WarrantyStartDate = dto.WarrantyStartDate
	}

	if dto.WarrantyEndDate != nil {
		hardware.WarrantyEndDate = dto.WarrantyEndDate
	}

	if dto.WarrantyCardNumber != nil {
		hardware.WarrantyCardNumber = dto.WarrantyCardNumber
	}

	if dto.DeviceSpecifications != nil {
		hardware.DeviceSpecifications = dto.DeviceSpecifications
	}

	if dto.AssetStatus != nil {
		hardware.AssetStatus = dto.AssetStatus
	}

	if dto.AssetResponsible != nil {
		hardware.AssetResponsible = dto.AssetResponsible
	}

	if dto.StorageLocation != nil {
		hardware.StorageLocation = dto.StorageLocation
	}

	if dto.UsagePeriod != nil {
		hardware.UsagePeriod = dto.UsagePeriod
	}

	if dto.AssetOutTime != nil {
		hardware.AssetOutTime = dto.AssetOutTime
	}

	if dto.AssetCondition != nil {
		hardware.AssetCondition = dto.AssetCondition
	}

	if dto.PurchaseReceipt != nil {
		hardware.PurchaseReceipt = dto.PurchaseReceipt
	}

	if dto.AssetID != nil {
		hardware.AssetID = dto.AssetID
	}

	if dto.DivisionID != nil {
		hardware.DivisionID = dto.DivisionID
	}

	if err := u.HardwareService.UpdatePartial(hardware); err != nil {
		return nil, err
	}

	response := &UpdatePartialHardwareResponseDTO{
		ID:                   hardware.ID,
		ReceiptDate:          hardware.ReceiptDate,
		ReceiptProof:         hardware.ReceiptProof,
		AssetType:            hardware.AssetType,
		DeviceActivationDate: hardware.DeviceActivationDate,
		InspectionResult:     hardware.InspectionResult,
		SerialNumber:         hardware.SerialNumber,
		Model:                hardware.Model,
		WarrantyStartDate:    hardware.WarrantyStartDate,
		WarrantyEndDate:      hardware.WarrantyEndDate,
		WarrantyCardNumber:   hardware.WarrantyCardNumber,
		DeviceSpecifications: hardware.DeviceSpecifications,
		AssetStatus:          hardware.AssetStatus,
		AssetResponsible:     hardware.AssetResponsible,
		StorageLocation:      hardware.StorageLocation,
		UsagePeriod:          hardware.UsagePeriod,
		AssetOutTime:         hardware.AssetOutTime,
		AssetCondition:       hardware.AssetCondition,
		PurchaseReceipt:      hardware.PurchaseReceipt,
		AssetID:              hardware.AssetID,
		DivisionID:           hardware.DivisionID,
		CreatedAt:            hardware.CreatedAt,
		UpdatedAt:            hardware.UpdatedAt,
	}

	return response, nil
}

func (u *HardwareCommandUsecase) Delete(dto *DeleteHardwareRequestDTO) (*DeleteHardwareResponseDTO, error) {
	hardware := &domain.Hardware{
		ID: dto.ID,
	}

	if err := u.HardwareService.Delete(hardware); err != nil {
		return nil, err
	}

	response := &DeleteHardwareResponseDTO{
		ID: hardware.ID,
	}

	return response, nil
}
