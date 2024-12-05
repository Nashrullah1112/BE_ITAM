package query

import (
	"github.com/banggibima/be-itam/modules/devices/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type DeviceQueryUsecase struct {
	Config        *config.Config
	DeviceService *domain.DeviceService
}

func NewDeviceQueryUsecase(
	config *config.Config,
	deviceService *domain.DeviceService,
) *DeviceQueryUsecase {
	return &DeviceQueryUsecase{
		Config:        config,
		DeviceService: deviceService,
	}
}

func (u *DeviceQueryUsecase) CountAll() (int, error) {
	response, err := u.DeviceService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *DeviceQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*DeviceResponseDTO, error) {
	devices, err := u.DeviceService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*DeviceResponseDTO, 0)
	for _, device := range devices {
		response = append(response, &DeviceResponseDTO{
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
		})
	}

	return response, nil
}

func (u *DeviceQueryUsecase) FindByID(id int) (*DeviceResponseDTO, error) {
	device, err := u.DeviceService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &DeviceResponseDTO{
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
