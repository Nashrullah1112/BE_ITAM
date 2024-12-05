package query

import (
	"github.com/banggibima/be-itam/modules/vendors/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type VendorQueryUsecase struct {
	Config        *config.Config
	VendorService *domain.VendorService
}

func NewVendorQueryUsecase(
	config *config.Config,
	vendorService *domain.VendorService,
) *VendorQueryUsecase {
	return &VendorQueryUsecase{
		Config:        config,
		VendorService: vendorService,
	}
}

func (u *VendorQueryUsecase) CountAll() (int, error) {
	response, err := u.VendorService.CountAll()
	if err != nil {
		return 0, err
	}

	return response, nil
}

func (u *VendorQueryUsecase) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*VendorResponseDTO, error) {
	vendors, err := u.VendorService.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return nil, err
	}

	response := make([]*VendorResponseDTO, 0)
	for _, vendor := range vendors {
		response = append(response, &VendorResponseDTO{
			ID:            vendor.ID,
			ContactPerson: vendor.ContactPerson,
			Email:         vendor.Email,
			ContactNumber: vendor.ContactNumber,
			Location:      vendor.Location,
			SIUPNumber:    vendor.SIUPNumber,
			NIBNumber:     vendor.NIBNumber,
			NPWPNumber:    vendor.NPWPNumber,
			CreatedAt:     vendor.CreatedAt,
			UpdatedAt:     vendor.UpdatedAt,
		})
	}

	return response, nil
}

func (u *VendorQueryUsecase) FindByID(id int) (*VendorResponseDTO, error) {
	vendor, err := u.VendorService.FindByID(id)
	if err != nil {
		return nil, err
	}

	response := &VendorResponseDTO{
		ID:            vendor.ID,
		ContactPerson: vendor.ContactPerson,
		Email:         vendor.Email,
		ContactNumber: vendor.ContactNumber,
		Location:      vendor.Location,
		SIUPNumber:    vendor.SIUPNumber,
		NIBNumber:     vendor.NIBNumber,
		NPWPNumber:    vendor.NPWPNumber,
		CreatedAt:     vendor.CreatedAt,
		UpdatedAt:     vendor.UpdatedAt,
	}

	return response, nil
}
