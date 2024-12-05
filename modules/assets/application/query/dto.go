package query

import (
	"time"
)

type AssetResponseDTO struct {
	ID            int       `json:"id"`
	SerialNumber  *string   `json:"serial_number"`
	Brand         *string   `json:"brand"`
	Model         *string   `json:"model"`
	ReceiptNumber *string   `json:"receipt_number"`
	Status        *string   `json:"status"`
	VendorID      *int      `json:"vendor_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
