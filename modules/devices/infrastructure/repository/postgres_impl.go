package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/devices/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresDeviceRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresDeviceRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresDeviceRepository {
	return &PostgresDeviceRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresDeviceRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM devices"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresDeviceRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Device, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM devices"

	for key, value := range filters {
		conditions = append(conditions, key+" = $"+strconv.Itoa(len(args)+1))
		args = append(args, value)
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if sort != "" && order != "" {
		query += " ORDER BY " + sort + " " + strings.ToUpper(order)
	}

	if limit > 0 && offset >= 0 {
		query += " LIMIT $" + strconv.Itoa(len(args)+1) + " OFFSET $" + strconv.Itoa(len(args)+2)
		args = append(args, limit, offset)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []*domain.Device{}
	for rows.Next() {
		device := domain.Device{}
		err := rows.Scan(
			&device.ID,
			&device.RecipientLocation,
			&device.ReceiptTime,
			&device.ReceiptProof,
			&device.AssetType,
			&device.AssetActivationTime,
			&device.AssetInspectionResult,
			&device.SerialNumber,
			&device.Model,
			&device.WarrantyStartTime,
			&device.WarrantyCardNumber,
			&device.Processor,
			&device.RAMCapacity,
			&device.ROMCapacity,
			&device.RAMType,
			&device.StorageType,
			&device.AssetStatus,
			&device.AssetValue,
			&device.DepreciationValue,
			&device.UsagePeriod,
			&device.AssetOutTime,
			&device.AssetConditionOnExit,
			&device.PurchaseReceipt,
			&device.AssetID,
			&device.DivisionID,
			&device.UserID,
			&device.CreatedAt,
			&device.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		devices = append(devices, &device)
	}

	return devices, nil
}

func (r *PostgresDeviceRepository) FindByID(id int) (*domain.Device, error) {
	query := "SELECT * FROM devices WHERE id = $1"

	device := &domain.Device{}

	err := r.DB.QueryRow(query, id).Scan(
		&device.ID,
		&device.RecipientLocation,
		&device.ReceiptTime,
		&device.ReceiptProof,
		&device.AssetType,
		&device.AssetActivationTime,
		&device.AssetInspectionResult,
		&device.SerialNumber,
		&device.Model,
		&device.WarrantyStartTime,
		&device.WarrantyCardNumber,
		&device.Processor,
		&device.RAMCapacity,
		&device.ROMCapacity,
		&device.RAMType,
		&device.StorageType,
		&device.AssetStatus,
		&device.AssetValue,
		&device.DepreciationValue,
		&device.UsagePeriod,
		&device.AssetOutTime,
		&device.AssetConditionOnExit,
		&device.PurchaseReceipt,
		&device.AssetID,
		&device.DivisionID,
		&device.UserID,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return device, nil
}

func (r *PostgresDeviceRepository) Create(device *domain.Device) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO devices (recipient_location, receipt_time, receipt_proof, asset_type, asset_activation_time, asset_inspection_result, serial_number, model, warranty_start_time, warranty_card_number, processor, ram_capacity, rom_capacity, ram_type, storage_type, asset_status, asset_value, depreciation_value, usage_period, asset_out_time, asset_condition_on_exit, purchase_receipt, asset_id, division_id, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27) RETURNING *"

	device.CreatedAt = time.Now()
	device.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		device.RecipientLocation,
		device.ReceiptTime,
		device.ReceiptProof,
		device.AssetType,
		device.AssetActivationTime,
		device.AssetInspectionResult,
		device.SerialNumber,
		device.Model,
		device.WarrantyStartTime,
		device.WarrantyCardNumber,
		device.Processor,
		device.RAMCapacity,
		device.ROMCapacity,
		device.RAMType,
		device.StorageType,
		device.AssetStatus,
		device.AssetValue,
		device.DepreciationValue,
		device.UsagePeriod,
		device.AssetOutTime,
		device.AssetConditionOnExit,
		device.PurchaseReceipt,
		device.AssetID,
		device.DivisionID,
		device.UserID,
		device.CreatedAt,
		device.UpdatedAt,
	).Scan(
		&device.ID,
		&device.RecipientLocation,
		&device.ReceiptTime,
		&device.ReceiptProof,
		&device.AssetType,
		&device.AssetActivationTime,
		&device.AssetInspectionResult,
		&device.SerialNumber,
		&device.Model,
		&device.WarrantyStartTime,
		&device.WarrantyCardNumber,
		&device.Processor,
		&device.RAMCapacity,
		&device.ROMCapacity,
		&device.RAMType,
		&device.StorageType,
		&device.AssetStatus,
		&device.AssetValue,
		&device.DepreciationValue,
		&device.UsagePeriod,
		&device.AssetOutTime,
		&device.AssetConditionOnExit,
		&device.PurchaseReceipt,
		&device.AssetID,
		&device.DivisionID,
		&device.UserID,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == pq.ErrorCode("23505") {
			field := strings.Split(pqErr.Constraint, "_")[1]
			return errors.New(field + " sudah ada")
		}
		return err
	}

	return tx.Commit()
}

func (r *PostgresDeviceRepository) Update(device *domain.Device) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE devices SET recipient_location = $2, receipt_time = $3, receipt_proof = $4, asset_type = $5, asset_activation_time = $6, asset_inspection_result = $7, serial_number = $8, model = $9, warranty_start_time = $10, warranty_card_number = $11, processor = $12, ram_capacity = $13, rom_capacity = $14, ram_type = $15, storage_type = $16, asset_status = $17, asset_value = $18, depreciation_value = $19, usage_period = $20, asset_out_time = $21, asset_condition_on_exit = $22, purchase_receipt = $23, asset_id = $24, division_id = $25, user_id = $26, updated_at = $27 WHERE id = $1 RETURNING *"

	device.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		device.ID,
		device.RecipientLocation,
		device.ReceiptTime,
		device.ReceiptProof,
		device.AssetType,
		device.AssetActivationTime,
		device.AssetInspectionResult,
		device.SerialNumber,
		device.Model,
		device.WarrantyStartTime,
		device.WarrantyCardNumber,
		device.Processor,
		device.RAMCapacity,
		device.ROMCapacity,
		device.RAMType,
		device.StorageType,
		device.AssetStatus,
		device.AssetValue,
		device.DepreciationValue,
		device.UsagePeriod,
		device.AssetOutTime,
		device.AssetConditionOnExit,
		device.PurchaseReceipt,
		device.AssetID,
		device.DivisionID,
		device.UserID,
		device.UpdatedAt,
	).Scan(
		&device.ID,
		&device.RecipientLocation,
		&device.ReceiptTime,
		&device.ReceiptProof,
		&device.AssetType,
		&device.AssetActivationTime,
		&device.AssetInspectionResult,
		&device.SerialNumber,
		&device.Model,
		&device.WarrantyStartTime,
		&device.WarrantyCardNumber,
		&device.Processor,
		&device.RAMCapacity,
		&device.ROMCapacity,
		&device.RAMType,
		&device.StorageType,
		&device.AssetStatus,
		&device.AssetValue,
		&device.DepreciationValue,
		&device.UsagePeriod,
		&device.AssetOutTime,
		&device.AssetConditionOnExit,
		&device.PurchaseReceipt,
		&device.AssetID,
		&device.DivisionID,
		&device.UserID,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data tidak ditemukan")
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == pq.ErrorCode("23505") {
			field := strings.Split(pqErr.Constraint, "_")[1]
			return errors.New(field + " sudah ada")
		}
		return err
	}

	return tx.Commit()
}

func (r *PostgresDeviceRepository) UpdatePartial(device *domain.Device) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE devices SET "

	if device.RecipientLocation != nil {
		query += "recipient_location = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.RecipientLocation)
	}

	if device.ReceiptTime != nil {
		query += "receipt_time = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.ReceiptTime)
	}

	if device.ReceiptProof != nil {
		query += "receipt_proof = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.ReceiptProof)
	}

	if device.AssetType != nil {
		query += "asset_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetType)
	}

	if device.AssetActivationTime != nil {
		query += "asset_activation_time = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetActivationTime)
	}

	if device.AssetInspectionResult != nil {
		query += "asset_inspection_result = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetInspectionResult)
	}

	if device.SerialNumber != nil {
		query += "serial_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.SerialNumber)
	}

	if device.Model != nil {
		query += "model = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.Model)
	}

	if device.WarrantyStartTime != nil {
		query += "warranty_start_time = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.WarrantyStartTime)
	}

	if device.WarrantyCardNumber != nil {
		query += "warranty_card_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.WarrantyCardNumber)
	}

	if device.Processor != nil {
		query += "processor = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.Processor)
	}

	if device.RAMCapacity != nil {
		query += "ram_capacity = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.RAMCapacity)
	}

	if device.ROMCapacity != nil {
		query += "rom_capacity = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.ROMCapacity)
	}

	if device.RAMType != nil {
		query += "ram_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.RAMType)
	}

	if device.StorageType != nil {
		query += "storage_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.StorageType)
	}

	if device.AssetStatus != nil {
		query += "asset_status = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetStatus)
	}

	if device.AssetValue != nil {
		query += "asset_value = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetValue)
	}

	if device.DepreciationValue != nil {
		query += "depreciation_value = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.DepreciationValue)
	}

	if device.UsagePeriod != nil {
		query += "usage_period = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.UsagePeriod)
	}

	if device.AssetOutTime != nil {
		query += "asset_out_time = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetOutTime)
	}

	if device.AssetConditionOnExit != nil {
		query += "asset_condition_on_exit = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetConditionOnExit)
	}

	if device.PurchaseReceipt != nil {
		query += "purchase_receipt = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.PurchaseReceipt)
	}

	if device.AssetID != nil {
		query += "asset_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.AssetID)
	}

	if device.DivisionID != nil {
		query += "division_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.DivisionID)
	}

	if device.UserID != nil {
		query += "user_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, device.UserID)
	}

	device.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, device.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, device.ID)

	err = tx.QueryRow(query, args...).Scan(
		&device.ID,
		&device.RecipientLocation,
		&device.ReceiptTime,
		&device.ReceiptProof,
		&device.AssetType,
		&device.AssetActivationTime,
		&device.AssetInspectionResult,
		&device.SerialNumber,
		&device.Model,
		&device.WarrantyStartTime,
		&device.WarrantyCardNumber,
		&device.Processor,
		&device.RAMCapacity,
		&device.ROMCapacity,
		&device.RAMType,
		&device.StorageType,
		&device.AssetStatus,
		&device.AssetValue,
		&device.DepreciationValue,
		&device.UsagePeriod,
		&device.AssetOutTime,
		&device.AssetConditionOnExit,
		&device.PurchaseReceipt,
		&device.AssetID,
		&device.DivisionID,
		&device.UserID,
		&device.CreatedAt,
		&device.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("data tidak ditemukan")
		}
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == pq.ErrorCode("23505") {
			field := strings.Split(pqErr.Constraint, "_")[1]
			return errors.New(field + " sudah ada")
		}
		return err
	}

	return tx.Commit()
}

func (r *PostgresDeviceRepository) Delete(device *domain.Device) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM devices WHERE id = $1"

	result, err := tx.Exec(query, device.ID)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("data tidak ditemukan")
	}

	return tx.Commit()
}
