package query

import (
	"github.com/banggibima/be-itam/modules/hardwares/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type HardwareQueryUsecase struct {
	Config          *config.Config
	HardwareService *domain.HardwareService
}

func NewHardwareQueryUsecase(
	config *config.Config,
	hardwareService *domain.HardwareService,
) *HardwareQueryUsecase {
	return &HardwareQueryUsecase{
		Config:          config,
		HardwareService: hardwareService,
	}
}

func (u *HardwareQueryUsecase) CountAll() (int, error) {
	response, err := u.HardwareService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *HardwareQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*HardwareResponseDTO, error) {
	hardwares, err := u.HardwareService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*HardwareResponseDTO, 0)
	for _, hardware := range hardwares {
		response = append(response, &HardwareResponseDTO{
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
		})
	}

	return response, nil
}

func (u *HardwareQueryUsecase) FindByID(id int) (*HardwareResponseDTO, error) {
	hardware, err := u.HardwareService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &HardwareResponseDTO{
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
