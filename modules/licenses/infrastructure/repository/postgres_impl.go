package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/licenses/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresLicenseRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresLicenseRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresLicenseRepository {
	return &PostgresLicenseRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresLicenseRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM licenses"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresLicenseRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.License, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM licenses"

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

	licenses := []*domain.License{}
	for rows.Next() {
		license := domain.License{}
		err := rows.Scan(
			&license.ID,
			&license.PurchaseDate,
			&license.InstalledDeviceSN,
			&license.ActivationDate,
			&license.ExpirationDate,
			&license.AssetOwnershipType,
			&license.LicenseCategory,
			&license.LicenseVersion,
			&license.MaxApplicationUsers,
			&license.MaxDeviceLicenses,
			&license.LicenseType,
			&license.AssetID,
			&license.CreatedAt,
			&license.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		licenses = append(licenses, &license)
	}

	return licenses, nil
}

func (r *PostgresLicenseRepository) FindByID(id int) (*domain.License, error) {
	query := "SELECT * FROM licenses WHERE id = $1"

	license := &domain.License{}

	err := r.DB.QueryRow(query, id).Scan(
		&license.ID,
		&license.PurchaseDate,
		&license.InstalledDeviceSN,
		&license.ActivationDate,
		&license.ExpirationDate,
		&license.AssetOwnershipType,
		&license.LicenseCategory,
		&license.LicenseVersion,
		&license.MaxApplicationUsers,
		&license.MaxDeviceLicenses,
		&license.LicenseType,
		&license.AssetID,
		&license.CreatedAt,
		&license.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return license, nil
}

func (r *PostgresLicenseRepository) Create(license *domain.License) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO licenses (purchase_date, installed_device_sn, activation_date, expiration_date, asset_ownership_type, license_category, license_version, max_application_users, max_device_licenses, license_type, asset_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING *"

	license.CreatedAt = time.Now()
	license.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		license.PurchaseDate,
		license.InstalledDeviceSN,
		license.ActivationDate,
		license.ExpirationDate,
		license.AssetOwnershipType,
		license.LicenseCategory,
		license.LicenseVersion,
		license.MaxApplicationUsers,
		license.MaxDeviceLicenses,
		license.LicenseType,
		license.AssetID,
		license.CreatedAt,
		license.UpdatedAt,
	).Scan(
		&license.ID,
		&license.PurchaseDate,
		&license.InstalledDeviceSN,
		&license.ActivationDate,
		&license.ExpirationDate,
		&license.AssetOwnershipType,
		&license.LicenseCategory,
		&license.LicenseVersion,
		&license.MaxApplicationUsers,
		&license.MaxDeviceLicenses,
		&license.LicenseType,
		&license.AssetID,
		&license.CreatedAt,
		&license.UpdatedAt,
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

func (r *PostgresLicenseRepository) Update(license *domain.License) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE licenses SET purchase_date = $2, installed_device_sn = $3, activation_date = $4, expiration_date = $5, asset_ownership_type = $6, license_category = $7, license_version = $8, max_application_users = $9, max_device_licenses = $10, license_type = $11, asset_id = $12, updated_at = $13 WHERE id = $1 RETURNING *"

	license.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		license.ID,
		license.PurchaseDate,
		license.InstalledDeviceSN,
		license.ActivationDate,
		license.ExpirationDate,
		license.AssetOwnershipType,
		license.LicenseCategory,
		license.LicenseVersion,
		license.MaxApplicationUsers,
		license.MaxDeviceLicenses,
		license.LicenseType,
		license.AssetID,
		license.UpdatedAt,
	).Scan(
		&license.ID,
		&license.PurchaseDate,
		&license.InstalledDeviceSN,
		&license.ActivationDate,
		&license.ExpirationDate,
		&license.AssetOwnershipType,
		&license.LicenseCategory,
		&license.LicenseVersion,
		&license.MaxApplicationUsers,
		&license.MaxDeviceLicenses,
		&license.LicenseType,
		&license.AssetID,
		&license.CreatedAt,
		&license.UpdatedAt,
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

func (r *PostgresLicenseRepository) UpdatePartial(license *domain.License) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE licenses SET "

	if license.PurchaseDate != nil {
		query += "purchase_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.PurchaseDate)
	}

	if license.InstalledDeviceSN != nil {
		query += "installed_device_sn = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.InstalledDeviceSN)
	}

	if license.ActivationDate != nil {
		query += "activation_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.ActivationDate)
	}

	if license.ExpirationDate != nil {
		query += "expiration_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.ExpirationDate)
	}

	if license.AssetOwnershipType != nil {
		query += "asset_ownership_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.AssetOwnershipType)
	}

	if license.LicenseCategory != nil {
		query += "license_category = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.LicenseCategory)
	}

	if license.LicenseVersion != nil {
		query += "license_version = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.LicenseVersion)
	}

	if license.MaxApplicationUsers != nil {
		query += "max_application_users = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.MaxApplicationUsers)
	}

	if license.MaxDeviceLicenses != nil {
		query += "max_device_licenses = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.MaxDeviceLicenses)
	}

	if license.LicenseType != nil {
		query += "license_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.LicenseType)
	}

	if license.AssetID != nil {
		query += "asset_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *license.AssetID)
	}

	license.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, license.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, license.ID)

	err = tx.QueryRow(query, args...).Scan(
		&license.ID,
		&license.PurchaseDate,
		&license.InstalledDeviceSN,
		&license.ActivationDate,
		&license.ExpirationDate,
		&license.AssetOwnershipType,
		&license.LicenseCategory,
		&license.LicenseVersion,
		&license.MaxApplicationUsers,
		&license.MaxDeviceLicenses,
		&license.LicenseType,
		&license.AssetID,
		&license.CreatedAt,
		&license.UpdatedAt,
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

func (r *PostgresLicenseRepository) Delete(license *domain.License) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM licenses WHERE id = $1"

	result, err := tx.Exec(query, license.ID)
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
