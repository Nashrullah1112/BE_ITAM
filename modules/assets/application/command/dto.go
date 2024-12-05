package command

import (
	"time"
)

type CreateAssetRequestDTO struct {
	SerialNumber  string `json:"serial_number" validate:"required"`
	Brand         string `json:"brand" validate:"required"`
	Model         string `json:"model" validate:"required"`
	ReceiptNumber string `json:"receipt_number" validate:"required"`
	Status        string `json:"status" validate:"required"`
	VendorID      *int   `json:"vendor_id" validate:"omitempty"`
}

type UpdateAssetRequestDTO struct {
	ID            int    `json:"id" validate:"required"`
	SerialNumber  string `json:"serial_number" validate:"required"`
	Brand         string `json:"brand" validate:"required"`
	Model         string `json:"model" validate:"required"`
	ReceiptNumber string `json:"receipt_number" validate:"required"`
	Status        string `json:"status" validate:"required"`
	VendorID      *int   `json:"vendor_id" validate:"omitempty"`
}

type UpdatePartialAssetRequestDTO struct {
	ID            int     `json:"id" validate:"required"`
	SerialNumber  *string `json:"serial_number" validate:"omitempty"`
	Brand         *string `json:"brand" validate:"omitempty"`
	Model         *string `json:"model" validate:"omitempty"`
	ReceiptNumber *string `json:"receipt_number" validate:"omitempty"`
	Status        *string `json:"status" validate:"omitempty"`
	VendorID      *int    `json:"vendor_id" validate:"omitempty"`
}

type DeleteAssetRequestDTO struct {
	ID int `json:"id" validate:"required"`
}

type CreateAssetResponseDTO struct {
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

type UpdateAssetResponseDTO struct {
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

type UpdatePartialAssetResponseDTO struct {
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

type DeleteAssetResponseDTO struct {
	ID int `json:"id"`
}
