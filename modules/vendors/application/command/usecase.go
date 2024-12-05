package command

import (
	"github.com/banggibima/be-itam/modules/vendors/domain"
	"github.com/banggibima/be-itam/pkg/config"
)

type VendorCommandUsecase struct {
	Config        *config.Config
	VendorService *domain.VendorService
}

func NewVendorCommandUsecase(
	config *config.Config,
	vendorService *domain.VendorService,
) *VendorCommandUsecase {
	return &VendorCommandUsecase{
		Config:        config,
		VendorService: vendorService,
	}
}

func (u *VendorCommandUsecase) Create(dto *CreateVendorRequestDTO) (*CreateVendorResponseDTO, error) {
	vendor := &domain.Vendor{
		ContactPerson: &dto.ContactPerson,
		Email:         &dto.Email,
		ContactNumber: &dto.ContactNumber,
		Location:      &dto.Location,
		SIUPNumber:    &dto.SIUPNumber,
		NIBNumber:     &dto.NIBNumber,
		NPWPNumber:    &dto.NPWPNumber,
	}

	if err := u.VendorService.Create(vendor); err != nil {
		return nil, err
	}

	response := &CreateVendorResponseDTO{
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

func (u *VendorCommandUsecase) Update(dto *UpdateVendorRequestDTO) (*UpdateVendorResponseDTO, error) {
	vendor := &domain.Vendor{
		ID:            dto.ID,
		ContactPerson: &dto.ContactPerson,
		Email:         &dto.Email,
		ContactNumber: &dto.ContactNumber,
		Location:      &dto.Location,
		SIUPNumber:    &dto.SIUPNumber,
		NIBNumber:     &dto.NIBNumber,
		NPWPNumber:    &dto.NPWPNumber,
	}

	if err := u.VendorService.Update(vendor); err != nil {
		return nil, err
	}

	response := &UpdateVendorResponseDTO{
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

func (u *VendorCommandUsecase) UpdatePartial(dto *UpdatePartialVendorRequestDTO) (*UpdatePartialVendorResponseDTO, error) {
	vendor := &domain.Vendor{
		ID: dto.ID,
	}

	if dto.ContactPerson != nil {
		vendor.ContactPerson = dto.ContactPerson
	}

	if dto.Email != nil {
		vendor.Email = dto.Email
	}

	if dto.ContactNumber != nil {
		vendor.ContactNumber = dto.ContactNumber
	}

	if dto.Location != nil {
		vendor.Location = dto.Location
	}

	if dto.SIUPNumber != nil {
		vendor.SIUPNumber = dto.SIUPNumber
	}

	if dto.NIBNumber != nil {
		vendor.NIBNumber = dto.NIBNumber
	}

	if dto.NPWPNumber != nil {
		vendor.NPWPNumber = dto.NPWPNumber
	}

	if err := u.VendorService.UpdatePartial(vendor); err != nil {
		return nil, err
	}

	response := &UpdatePartialVendorResponseDTO{
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

func (u *VendorCommandUsecase) Delete(dto *DeleteVendorRequestDTO) (*DeleteVendorResponseDTO, error) {
	vendor := &domain.Vendor{
		ID: dto.ID,
	}

	if err := u.VendorService.Delete(vendor); err != nil {
		return nil, err
	}

	response := &DeleteVendorResponseDTO{
		ID: vendor.ID,
	}

	return response, nil
}
