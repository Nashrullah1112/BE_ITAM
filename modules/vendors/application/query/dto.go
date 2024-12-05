package query

import (
	"time"
)

type VendorResponseDTO struct {
	ID            int       `json:"id"`
	ContactPerson *string   `json:"contact_person"`
	Email         *string   `json:"email"`
	ContactNumber *string   `json:"contact_number"`
	Location      *string   `json:"location"`
	SIUPNumber    *string   `json:"siup_number"`
	NIBNumber     *string   `json:"nib_number"`
	NPWPNumber    *string   `json:"npwp_number"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
