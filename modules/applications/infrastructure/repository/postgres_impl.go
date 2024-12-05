package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/applications/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresApplicationRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresApplicationRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresApplicationRepository {
	return &PostgresApplicationRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresApplicationRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM applications"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresApplicationRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Application, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM applications"

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

	applications := []*domain.Application{}
	for rows.Next() {
		application := domain.Application{}
		err := rows.Scan(
			&application.ID,
			&application.CreatedAt,
			&application.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		applications = append(applications, &application)
	}

	return applications, nil
}

func (r *PostgresApplicationRepository) FindByID(id int) (*domain.Application, error) {
	query := "SELECT * FROM applications WHERE id = $1"

	application := &domain.Application{}

	err := r.DB.QueryRow(query, id).Scan(
		&application.ID,
		&application.ApplicationName,
		&application.CreationDate,
		&application.AcceptanceDate,
		&application.StorageServerLocation,
		&application.ApplicationType,
		&application.ApplicationLink,
		&application.ApplicationCertification,
		&application.ActivationDate,
		&application.ExpirationDate,
		&application.AssetID,
		&application.CreatedAt,
		&application.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return application, nil
}

func (r *PostgresApplicationRepository) Create(application *domain.Application) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO applications (application_name, creation_date, acceptance_date, storage_server_location, application_type, application_link, application_certification, activation_date, expiration_date, asset_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING *"

	application.CreatedAt = time.Now()
	application.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		application.ApplicationName,
		application.CreationDate,
		application.AcceptanceDate,
		application.StorageServerLocation,
		application.ApplicationType,
		application.ApplicationLink,
		application.ApplicationCertification,
		application.ActivationDate,
		application.ExpirationDate,
		application.AssetID,
		application.CreatedAt,
		application.UpdatedAt,
	).Scan(
		&application.ID,
		&application.ApplicationName,
		&application.CreationDate,
		&application.AcceptanceDate,
		&application.StorageServerLocation,
		&application.ApplicationType,
		&application.ApplicationLink,
		&application.ApplicationCertification,
		&application.ActivationDate,
		&application.ExpirationDate,
		&application.AssetID,
		&application.CreatedAt,
		&application.UpdatedAt,
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

func (r *PostgresApplicationRepository) Update(application *domain.Application) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE applications SET application_name = $2, creation_date = $3, acceptance_date = $4, storage_server_location = $5, application_type = $6, application_link = $7, application_certification = $8, activation_date = $9, expiration_date = $10, asset_id = $11, updated_at = $12 WHERE id = $1 RETURNING *"

	application.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		application.ID,
		application.ApplicationName,
		application.CreationDate,
		application.AcceptanceDate,
		application.StorageServerLocation,
		application.ApplicationType,
		application.ApplicationLink,
		application.ApplicationCertification,
		application.ActivationDate,
		application.ExpirationDate,
		application.AssetID,
		application.UpdatedAt,
	).Scan(
		&application.ID,
		&application.ApplicationName,
		&application.CreationDate,
		&application.AcceptanceDate,
		&application.StorageServerLocation,
		&application.ApplicationType,
		&application.ApplicationLink,
		&application.ApplicationCertification,
		&application.ActivationDate,
		&application.ExpirationDate,
		&application.AssetID,
		&application.CreatedAt,
		&application.UpdatedAt,
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

func (r *PostgresApplicationRepository) UpdatePartial(application *domain.Application) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE applications SET "

	if application.ApplicationName != nil {
		query += "application_name = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ApplicationName)
	}

	if application.CreationDate != nil {
		query += "creation_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.CreationDate)
	}

	if application.AcceptanceDate != nil {
		query += "acceptance_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.AcceptanceDate)
	}

	if application.StorageServerLocation != nil {
		query += "storage_server_location = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.StorageServerLocation)
	}

	if application.ApplicationType != nil {
		query += "application_type = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ApplicationType)
	}

	if application.ApplicationLink != nil {
		query += "application_link = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ApplicationLink)
	}

	if application.ApplicationCertification != nil {
		query += "application_certification = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ApplicationCertification)
	}

	if application.ActivationDate != nil {
		query += "activation_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ActivationDate)
	}

	if application.ExpirationDate != nil {
		query += "expiration_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.ExpirationDate)
	}

	if application.AssetID != nil {
		query += "asset_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *application.AssetID)
	}

	application.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, application.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, application.ID)

	err = tx.QueryRow(query, args...).Scan(
		&application.ID,
		&application.ApplicationName,
		&application.CreationDate,
		&application.AcceptanceDate,
		&application.StorageServerLocation,
		&application.ApplicationType,
		&application.ApplicationLink,
		&application.ApplicationCertification,
		&application.ActivationDate,
		&application.ExpirationDate,
		&application.AssetID,
		&application.CreatedAt,
		&application.UpdatedAt,
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

func (r *PostgresApplicationRepository) Delete(application *domain.Application) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM applications WHERE id = $1"

	result, err := tx.Exec(query, application.ID)
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
