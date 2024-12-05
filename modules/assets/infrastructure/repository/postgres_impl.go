package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/assets/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresAssetRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresAssetRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresAssetRepository {
	return &PostgresAssetRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresAssetRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM assets"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresAssetRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Asset, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM assets"

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

	assets := []*domain.Asset{}
	for rows.Next() {
		asset := domain.Asset{}
		err := rows.Scan(
			&asset.ID,
			&asset.SerialNumber,
			&asset.Brand,
			&asset.Model,
			&asset.ReceiptNumber,
			&asset.Status,
			&asset.VendorID,
			&asset.CreatedAt,
			&asset.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}

func (r *PostgresAssetRepository) FindByID(id int) (*domain.Asset, error) {
	query := "SELECT * FROM assets WHERE id = $1"

	asset := &domain.Asset{}

	err := r.DB.QueryRow(query, id).Scan(
		&asset.ID,
		&asset.SerialNumber,
		&asset.Brand,
		&asset.Model,
		&asset.ReceiptNumber,
		&asset.Status,
		&asset.VendorID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return asset, nil
}

func (r *PostgresAssetRepository) Create(asset *domain.Asset) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO assets (serial_number, brand, model, receipt_number, status, vendor_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *"

	asset.CreatedAt = time.Now()
	asset.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		asset.SerialNumber,
		asset.Brand,
		asset.Model,
		asset.ReceiptNumber,
		asset.Status,
		asset.VendorID,
		asset.CreatedAt,
		asset.UpdatedAt,
	).Scan(
		&asset.ID,
		&asset.SerialNumber,
		&asset.Brand,
		&asset.Model,
		&asset.ReceiptNumber,
		&asset.Status,
		&asset.VendorID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
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

func (r *PostgresAssetRepository) Update(asset *domain.Asset) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE assets SET serial_number = $2, brand = $3, model = $4, receipt_number = $5, status = $6, vendor_id = $7, updated_at = $8 WHERE id = $1 RETURNING *"

	asset.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		asset.ID,
		asset.SerialNumber,
		asset.Brand,
		asset.Model,
		asset.ReceiptNumber,
		asset.Status,
		asset.VendorID,
		asset.UpdatedAt,
	).Scan(
		&asset.ID,
		&asset.SerialNumber,
		&asset.Brand,
		&asset.Model,
		&asset.ReceiptNumber,
		&asset.Status,
		&asset.VendorID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
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

func (r *PostgresAssetRepository) UpdatePartial(asset *domain.Asset) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE assets SET "

	if asset.SerialNumber != nil {
		query += "serial_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.SerialNumber)
	}

	if asset.Brand != nil {
		query += "brand = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.Brand)
	}

	if asset.Model != nil {
		query += "model = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.Model)
	}

	if asset.ReceiptNumber != nil {
		query += "receipt_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.ReceiptNumber)
	}

	if asset.Status != nil {
		query += "status = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.Status)
	}

	if asset.VendorID != nil {
		query += "vendor_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *asset.VendorID)
	}

	asset.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, asset.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, asset.ID)

	err = tx.QueryRow(query, args...).Scan(
		&asset.ID,
		&asset.SerialNumber,
		&asset.Brand,
		&asset.Model,
		&asset.ReceiptNumber,
		&asset.Status,
		&asset.VendorID,
		&asset.CreatedAt,
		&asset.UpdatedAt,
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

func (r *PostgresAssetRepository) Delete(asset *domain.Asset) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM assets WHERE id = $1"

	result, err := tx.Exec(query, asset.ID)
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
