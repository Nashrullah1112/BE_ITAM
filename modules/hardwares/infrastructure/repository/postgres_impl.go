package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/hardwares/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresHardwareRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresHardwareRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresHardwareRepository {
	return &PostgresHardwareRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresHardwareRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM hardwares"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresHardwareRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Hardware, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM hardwares"

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

	hardwares := []*domain.Hardware{}
	for rows.Next() {
		hardware := domain.Hardware{}
		err := rows.Scan(
			&hardware.ID,
			&hardware.ReceiptDate,
			&hardware.ReceiptProof,
			&hardware.AssetType,
			&hardware.DeviceActivationDate,
			&hardware.InspectionResult,
			&hardware.SerialNumber,
			&hardware.Model,
			&hardware.WarrantyStartDate,
			&hardware.WarrantyEndDate,
			&hardware.WarrantyCardNumber,
			&hardware.DeviceSpecifications,
			&hardware.AssetStatus,
			&hardware.AssetResponsible,
			&hardware.StorageLocation,
			&hardware.UsagePeriod,
			&hardware.AssetOutTime,
			&hardware.AssetCondition,
			&hardware.PurchaseReceipt,
			&hardware.AssetID,
			&hardware.DivisionID,
			&hardware.CreatedAt,
			&hardware.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		hardwares = append(hardwares, &hardware)
	}

	return hardwares, nil
}

func (r *PostgresHardwareRepository) FindByID(id int) (*domain.Hardware, error) {
	query := "SELECT * FROM hardwares WHERE id = $1"

	hardware := &domain.Hardware{}

	err := r.DB.QueryRow(query, id).Scan(
		&hardware.ID,
		&hardware.ReceiptDate,
		&hardware.ReceiptProof,
		&hardware.AssetType,
		&hardware.DeviceActivationDate,
		&hardware.InspectionResult,
		&hardware.SerialNumber,
		&hardware.Model,
		&hardware.WarrantyStartDate,
		&hardware.WarrantyEndDate,
		&hardware.WarrantyCardNumber,
		&hardware.DeviceSpecifications,
		&hardware.AssetStatus,
		&hardware.AssetResponsible,
		&hardware.StorageLocation,
		&hardware.UsagePeriod,
		&hardware.AssetOutTime,
		&hardware.AssetCondition,
		&hardware.PurchaseReceipt,
		&hardware.AssetID,
		&hardware.DivisionID,
		&hardware.CreatedAt,
		&hardware.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return hardware, nil
}

func (r *PostgresHardwareRepository) Create(hardware *domain.Hardware) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO hardwares (receipt_date, receipt_proof, asset_type, device_activation_date, inspection_result, serial_number, model, warranty_start_date, warranty_end_date, warranty_card_number, device_specifications, asset_status, asset_responsible, storage_location, usage_period, asset_out_time, asset_condition, purchase_receipt, asset_id, division_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22) RETURNING *"

	hardware.CreatedAt = time.Now()
	hardware.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		hardware.ReceiptDate,
		hardware.ReceiptProof,
		hardware.AssetType,
		hardware.DeviceActivationDate,
		hardware.InspectionResult,
		hardware.SerialNumber,
		hardware.Model,
		hardware.WarrantyStartDate,
		hardware.WarrantyEndDate,
		hardware.WarrantyCardNumber,
		hardware.DeviceSpecifications,
		hardware.AssetStatus,
		hardware.AssetResponsible,
		hardware.StorageLocation,
		hardware.UsagePeriod,
		hardware.AssetOutTime,
		hardware.AssetCondition,
		hardware.PurchaseReceipt,
		hardware.AssetID,
		hardware.DivisionID,
		hardware.CreatedAt,
		hardware.UpdatedAt,
	).Scan(
		&hardware.ID,
		&hardware.ReceiptDate,
		&hardware.ReceiptProof,
		&hardware.AssetType,
		&hardware.DeviceActivationDate,
		&hardware.InspectionResult,
		&hardware.SerialNumber,
		&hardware.Model,
		&hardware.WarrantyStartDate,
		&hardware.WarrantyEndDate,
		&hardware.WarrantyCardNumber,
		&hardware.DeviceSpecifications,
		&hardware.AssetStatus,
		&hardware.AssetResponsible,
		&hardware.StorageLocation,
		&hardware.UsagePeriod,
		&hardware.AssetOutTime,
		&hardware.AssetCondition,
		&hardware.PurchaseReceipt,
		&hardware.AssetID,
		&hardware.DivisionID,
		&hardware.CreatedAt,
		&hardware.UpdatedAt,
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

func (r *PostgresHardwareRepository) Update(hardware *domain.Hardware) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE hardwares SET receipt_date = $2, receipt_proof = $3, asset_type = $4, device_activation_date = $5, inspection_result = $6, serial_number = $7, model = $8, warranty_start_date = $9, warranty_end_date = $10, warranty_card_number = $11, device_specifications = $12, asset_status = $13, asset_responsible = $14, storage_location = $15, usage_period = $16, asset_out_time = $17, asset_condition = $18, purchase_receipt = $19, asset_id = $20, division_id = $21, updated_at = $22 WHERE id = $1 RETURNING *"

	hardware.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		hardware.ID,
		hardware.ReceiptDate,
		hardware.ReceiptProof,
		hardware.AssetType,
		hardware.DeviceActivationDate,
		hardware.InspectionResult,
		hardware.SerialNumber,
		hardware.Model,
		hardware.WarrantyStartDate,
		hardware.WarrantyEndDate,
		hardware.WarrantyCardNumber,
		hardware.DeviceSpecifications,
		hardware.AssetStatus,
		hardware.AssetResponsible,
		hardware.StorageLocation,
		hardware.UsagePeriod,
		hardware.AssetOutTime,
		hardware.AssetCondition,
		hardware.PurchaseReceipt,
		hardware.AssetID,
		hardware.DivisionID,
		hardware.UpdatedAt,
	).Scan(
		&hardware.ID,
		&hardware.ReceiptDate,
		&hardware.ReceiptProof,
		&hardware.AssetType,
		&hardware.DeviceActivationDate,
		&hardware.InspectionResult,
		&hardware.SerialNumber,
		&hardware.Model,
		&hardware.WarrantyStartDate,
		&hardware.WarrantyEndDate,
		&hardware.WarrantyCardNumber,
		&hardware.DeviceSpecifications,
		&hardware.AssetStatus,
		&hardware.AssetResponsible,
		&hardware.StorageLocation,
		&hardware.UsagePeriod,
		&hardware.AssetOutTime,
		&hardware.AssetCondition,
		&hardware.PurchaseReceipt,
		&hardware.AssetID,
		&hardware.DivisionID,
		&hardware.CreatedAt,
		&hardware.UpdatedAt,
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

func (r *PostgresHardwareRepository) UpdatePartial(hardware *domain.Hardware) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE hardwares SET "

	if hardware.ReceiptDate != nil {
		query += "receipt_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.ReceiptDate)
	}

	if hardware.ReceiptProof != nil {
		query += "receipt_proof = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.ReceiptProof)
	}

	if hardware.AssetType != nil {
		query += "asset_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetType)
	}

	if hardware.DeviceActivationDate != nil {
		query += "device_activation_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.DeviceActivationDate)
	}

	if hardware.InspectionResult != nil {
		query += "inspection_result = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.InspectionResult)
	}

	if hardware.SerialNumber != nil {
		query += "serial_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.SerialNumber)
	}

	if hardware.Model != nil {
		query += "model = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.Model)
	}

	if hardware.WarrantyStartDate != nil {
		query += "warranty_start_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.WarrantyStartDate)
	}

	if hardware.WarrantyEndDate != nil {
		query += "warranty_end_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.WarrantyEndDate)
	}

	if hardware.WarrantyCardNumber != nil {
		query += "warranty_card_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.WarrantyCardNumber)
	}

	if hardware.DeviceSpecifications != nil {
		query += "device_specifications = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.DeviceSpecifications)
	}
	if hardware.AssetStatus != nil {
		query += "asset_status = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetStatus)
	}

	if hardware.AssetResponsible != nil {
		query += "asset_responsible = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetResponsible)
	}

	if hardware.StorageLocation != nil {
		query += "storage_location = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.StorageLocation)
	}

	if hardware.UsagePeriod != nil {
		query += "usage_period = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.UsagePeriod)
	}

	if hardware.AssetOutTime != nil {
		query += "asset_out_time = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetOutTime)
	}

	if hardware.AssetCondition != nil {
		query += "asset_condition = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetCondition)
	}

	if hardware.PurchaseReceipt != nil {
		query += "purchase_receipt = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.PurchaseReceipt)
	}

	if hardware.AssetID != nil {
		query += "asset_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.AssetID)
	}

	if hardware.DivisionID != nil {
		query += "division_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *hardware.DivisionID)
	}

	hardware.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, hardware.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, hardware.ID)

	err = tx.QueryRow(query, args...).Scan(
		&hardware.ID,
		&hardware.ReceiptDate,
		&hardware.ReceiptProof,
		&hardware.AssetType,
		&hardware.DeviceActivationDate,
		&hardware.InspectionResult,
		&hardware.SerialNumber,
		&hardware.Model,
		&hardware.WarrantyStartDate,
		&hardware.WarrantyEndDate,
		&hardware.WarrantyCardNumber,
		&hardware.DeviceSpecifications,
		&hardware.AssetStatus,
		&hardware.AssetResponsible,
		&hardware.StorageLocation,
		&hardware.UsagePeriod,
		&hardware.AssetOutTime,
		&hardware.AssetCondition,
		&hardware.PurchaseReceipt,
		&hardware.AssetID,
		&hardware.DivisionID,
		&hardware.CreatedAt,
		&hardware.UpdatedAt,
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

func (r *PostgresHardwareRepository) Delete(hardware *domain.Hardware) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM hardwares WHERE id = $1"

	result, err := tx.Exec(query, hardware.ID)
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
