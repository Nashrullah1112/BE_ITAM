package domain

import (
	"time"
)

type Asset struct {
	ID            int
	SerialNumber  *string
	Brand         *string
	Model         *string
	ReceiptNumber *string
	Status        *string
	VendorID      *int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
