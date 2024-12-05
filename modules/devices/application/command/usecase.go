package command

import (
	"github.com/banggibima/be-itam/modules/devices/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type DeviceCommandUsecase struct {
	Config        *config.Config
	DeviceService *domain.DeviceService
}

func NewDeviceCommandUsecase(
	config *config.Config,
	deviceService *domain.DeviceService,
) *DeviceCommandUsecase {
	return &DeviceCommandUsecase{
		Config:        config,
		DeviceService: deviceService,
	}
}

func (u *DeviceCommandUsecase) Create(dto *CreateDeviceRequestDTO) (*CreateDeviceResponseDTO, error) {
	device := &domain.Device{
		RecipientLocation:     &dto.RecipientLocation,
		ReceiptTime:           &dto.ReceiptTime,
		ReceiptProof:          &dto.ReceiptProof,
		AssetType:             &dto.AssetType,
		AssetActivationTime:   &dto.AssetActivationTime,
		AssetInspectionResult: &dto.AssetInspectionResult,
		SerialNumber:          &dto.SerialNumber,
		Model:                 &dto.Model,
		WarrantyStartTime:     &dto.WarrantyStartTime,
		WarrantyCardNumber:    &dto.WarrantyCardNumber,
		Processor:             &dto.Processor,
		RAMCapacity:           &dto.RAMCapacity,
		ROMCapacity:           &dto.ROMCapacity,
		RAMType:               &dto.RAMType,
		StorageType:           &dto.StorageType,
		AssetStatus:           &dto.AssetStatus,
		AssetValue:            &dto.AssetValue,
		DepreciationValue:     &dto.DepreciationValue,
		UsagePeriod:           &dto.UsagePeriod,
		AssetOutTime:          &dto.AssetOutTime,
		AssetConditionOnExit:  &dto.AssetConditionOnExit,
		PurchaseReceipt:       &dto.PurchaseReceipt,
		AssetID:               dto.AssetID,
		DivisionID:            dto.DivisionID,
		UserID:                dto.UserID,
	}

	if err := u.DeviceService.Create(device); err != nil {
		return nil, err
	}

	response := &CreateDeviceResponseDTO{
		ID:                    device.ID,
		RecipientLocation:     device.RecipientLocation,
		ReceiptTime:           device.ReceiptTime,
		ReceiptProof:          device.ReceiptProof,
		AssetType:             device.AssetType,
		AssetActivationTime:   device.AssetActivationTime,
		AssetInspectionResult: device.AssetInspectionResult,
		SerialNumber:          device.SerialNumber,
		Model:                 device.Model,
		WarrantyStartTime:     device.WarrantyStartTime,
		WarrantyCardNumber:    device.WarrantyCardNumber,
		Processor:             device.Processor,
		RAMCapacity:           device.RAMCapacity,
		ROMCapacity:           device.ROMCapacity,
		RAMType:               device.RAMType,
		StorageType:           device.StorageType,
		AssetStatus:           device.AssetStatus,
		AssetValue:            device.AssetValue,
		DepreciationValue:     device.DepreciationValue,
		UsagePeriod:           device.UsagePeriod,
		AssetOutTime:          device.AssetOutTime,
		AssetConditionOnExit:  device.AssetConditionOnExit,
		PurchaseReceipt:       device.PurchaseReceipt,
		AssetID:               device.AssetID,
		DivisionID:            device.DivisionID,
		UserID:                device.UserID,
		CreatedAt:             device.CreatedAt,
		UpdatedAt:             device.UpdatedAt,
	}

	return response, nil
}

func (u *DeviceCommandUsecase) Update(dto *UpdateDeviceRequestDTO) (*UpdateDeviceResponseDTO, error) {
	device := &domain.Device{
		ID:                    dto.ID,
		RecipientLocation:     &dto.RecipientLocation,
		ReceiptTime:           &dto.ReceiptTime,
		ReceiptProof:          &dto.ReceiptProof,
		AssetType:             &dto.AssetType,
		AssetActivationTime:   &dto.AssetActivationTime,
		AssetInspectionResult: &dto.AssetInspectionResult,
		SerialNumber:          &dto.SerialNumber,
		Model:                 &dto.Model,
		WarrantyStartTime:     &dto.WarrantyStartTime,
		WarrantyCardNumber:    &dto.WarrantyCardNumber,
		Processor:             &dto.Processor,
		RAMCapacity:           &dto.RAMCapacity,
		ROMCapacity:           &dto.ROMCapacity,
		RAMType:               &dto.RAMType,
		StorageType:           &dto.StorageType,
		AssetStatus:           &dto.AssetStatus,
		AssetValue:            &dto.AssetValue,
		DepreciationValue:     &dto.DepreciationValue,
		UsagePeriod:           &dto.UsagePeriod,
		AssetOutTime:          &dto.AssetOutTime,
		AssetConditionOnExit:  &dto.AssetConditionOnExit,
		PurchaseReceipt:       &dto.PurchaseReceipt,
		AssetID:               dto.AssetID,
		DivisionID:            dto.DivisionID,
		UserID:                dto.UserID,
	}

	if err := u.DeviceService.Update(device); err != nil {
		return nil, err
	}

	response := &UpdateDeviceResponseDTO{
		ID:                    device.ID,
		RecipientLocation:     device.RecipientLocation,
		ReceiptTime:           device.ReceiptTime,
		ReceiptProof:          device.ReceiptProof,
		AssetType:             device.AssetType,
		AssetActivationTime:   device.AssetActivationTime,
		AssetInspectionResult: device.AssetInspectionResult,
		SerialNumber:          device.SerialNumber,
		Model:                 device.Model,
		WarrantyStartTime:     device.WarrantyStartTime,
		WarrantyCardNumber:    device.WarrantyCardNumber,
		Processor:             device.Processor,
		RAMCapacity:           device.RAMCapacity,
		ROMCapacity:           device.ROMCapacity,
		RAMType:               device.RAMType,
		StorageType:           device.StorageType,
		AssetStatus:           device.AssetStatus,
		AssetValue:            device.AssetValue,
		DepreciationValue:     device.DepreciationValue,
		UsagePeriod:           device.UsagePeriod,
		AssetOutTime:          device.AssetOutTime,
		AssetConditionOnExit:  device.AssetConditionOnExit,
		PurchaseReceipt:       device.PurchaseReceipt,
		AssetID:               device.AssetID,
		DivisionID:            device.DivisionID,
		UserID:                device.UserID,
		CreatedAt:             device.CreatedAt,
		UpdatedAt:             device.UpdatedAt,
	}

	return response, nil
}

func (u *DeviceCommandUsecase) UpdatePartial(dto *UpdatePartialDeviceRequestDTO) (*UpdatePartialDeviceResponseDTO, error) {
	device := &domain.Device{
		ID: dto.ID,
	}

	if dto.RecipientLocation != nil {
		device.RecipientLocation = dto.RecipientLocation
	}

	if dto.ReceiptTime != nil {
		device.ReceiptTime = dto.ReceiptTime
	}

	if dto.ReceiptProof != nil {
		device.ReceiptProof = dto.ReceiptProof
	}

	if dto.AssetType != nil {
		device.AssetType = dto.AssetType
	}

	if dto.AssetActivationTime != nil {
		device.AssetActivationTime = dto.AssetActivationTime
	}

	if dto.AssetInspectionResult != nil {
		device.AssetInspectionResult = dto.AssetInspectionResult
	}

	if dto.SerialNumber != nil {
		device.SerialNumber = dto.SerialNumber
	}

	if dto.Model != nil {
		device.Model = dto.Model
	}

	if dto.WarrantyStartTime != nil {
		device.WarrantyStartTime = dto.WarrantyStartTime
	}

	if dto.WarrantyCardNumber != nil {
		device.WarrantyCardNumber = dto.WarrantyCardNumber
	}

	if dto.Processor != nil {
		device.Processor = dto.Processor
	}

	if dto.RAMCapacity != nil {
		device.RAMCapacity = dto.RAMCapacity
	}

	if dto.ROMCapacity != nil {
		device.ROMCapacity = dto.ROMCapacity
	}

	if dto.RAMType != nil {
		device.RAMType = dto.RAMType
	}

	if dto.StorageType != nil {
		device.StorageType = dto.StorageType
	}

	if dto.AssetStatus != nil {
		device.AssetStatus = dto.AssetStatus
	}

	if dto.AssetValue != nil {
		device.AssetValue = dto.AssetValue
	}

	if dto.DepreciationValue != nil {
		device.DepreciationValue = dto.DepreciationValue
	}

	if dto.UsagePeriod != nil {
		device.UsagePeriod = dto.UsagePeriod
	}

	if dto.AssetOutTime != nil {
		device.AssetOutTime = dto.AssetOutTime
	}

	if dto.AssetConditionOnExit != nil {
		device.AssetConditionOnExit = dto.AssetConditionOnExit
	}

	if dto.PurchaseReceipt != nil {
		device.PurchaseReceipt = dto.PurchaseReceipt
	}

	if dto.AssetID != nil {
		device.AssetID = dto.AssetID
	}

	if dto.DivisionID != nil {
		device.DivisionID = dto.DivisionID
	}

	if dto.UserID != nil {
		device.UserID = dto.UserID
	}

	if err := u.DeviceService.UpdatePartial(device); err != nil {
		return nil, err
	}

	response := &UpdatePartialDeviceResponseDTO{
		ID:                    device.ID,
		RecipientLocation:     device.RecipientLocation,
		ReceiptTime:           device.ReceiptTime,
		ReceiptProof:          device.ReceiptProof,
		AssetType:             device.AssetType,
		AssetActivationTime:   device.AssetActivationTime,
		AssetInspectionResult: device.AssetInspectionResult,
		SerialNumber:          device.SerialNumber,
		Model:                 device.Model,
		WarrantyStartTime:     device.WarrantyStartTime,
		WarrantyCardNumber:    device.WarrantyCardNumber,
		Processor:             device.Processor,
		RAMCapacity:           device.RAMCapacity,
		ROMCapacity:           device.ROMCapacity,
		RAMType:               device.RAMType,
		StorageType:           device.StorageType,
		AssetStatus:           device.AssetStatus,
		AssetValue:            device.AssetValue,
		DepreciationValue:     device.DepreciationValue,
		UsagePeriod:           device.UsagePeriod,
		AssetOutTime:          device.AssetOutTime,
		AssetConditionOnExit:  device.AssetConditionOnExit,
		PurchaseReceipt:       device.PurchaseReceipt,
		AssetID:               device.AssetID,
		DivisionID:            device.DivisionID,
		UserID:                device.UserID,
		CreatedAt:             device.CreatedAt,
		UpdatedAt:             device.UpdatedAt,
	}

	return response, nil
}

func (u *DeviceCommandUsecase) Delete(dto *DeleteDeviceRequestDTO) (*DeleteDeviceResponseDTO, error) {
	device := &domain.Device{
		ID: dto.ID,
	}

	if err := u.DeviceService.Delete(device); err != nil {
		return nil, err
	}

	response := &DeleteDeviceResponseDTO{
		ID: device.ID,
	}

	return response, nil
}
