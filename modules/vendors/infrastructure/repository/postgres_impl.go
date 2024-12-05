package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/vendors/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresVendorRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresVendorRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresVendorRepository {
	return &PostgresVendorRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresVendorRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM vendors"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresVendorRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Vendor, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM vendors"

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

	vendors := []*domain.Vendor{}
	for rows.Next() {
		vendor := domain.Vendor{}
		err := rows.Scan(
			&vendor.ID,
			&vendor.ContactPerson,
			&vendor.Email,
			&vendor.ContactNumber,
			&vendor.Location,
			&vendor.SIUPNumber,
			&vendor.NIBNumber,
			&vendor.NPWPNumber,
			&vendor.CreatedAt,
			&vendor.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		vendors = append(vendors, &vendor)
	}

	return vendors, nil
}

func (r *PostgresVendorRepository) FindByID(id int) (*domain.Vendor, error) {
	query := "SELECT * FROM vendors WHERE id = $1"

	vendor := &domain.Vendor{}

	err := r.DB.QueryRow(query, id).Scan(
		&vendor.ID,
		&vendor.ContactPerson,
		&vendor.Email,
		&vendor.ContactNumber,
		&vendor.Location,
		&vendor.SIUPNumber,
		&vendor.NIBNumber,
		&vendor.NPWPNumber,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return vendor, nil
}

func (r *PostgresVendorRepository) Create(vendor *domain.Vendor) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO vendors (contact_person, email, contact_number, location, siup_number, nib_number, npwp_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *"

	vendor.CreatedAt = time.Now()
	vendor.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		vendor.ContactPerson,
		vendor.Email,
		vendor.ContactNumber,
		vendor.Location,
		vendor.SIUPNumber,
		vendor.NIBNumber,
		vendor.NPWPNumber,
		vendor.CreatedAt,
		vendor.UpdatedAt,
	).Scan(
		&vendor.ID,
		&vendor.ContactPerson,
		&vendor.Email,
		&vendor.ContactNumber,
		&vendor.Location,
		&vendor.SIUPNumber,
		&vendor.NIBNumber,
		&vendor.NPWPNumber,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
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

func (r *PostgresVendorRepository) Update(vendor *domain.Vendor) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE vendors SET contact_person = $2, email = $3, contact_number = $4, location = $5, siup_number = $6, nib_number = $7, npwp_number = $8, updated_at = $9 WHERE id = $1 RETURNING *"

	vendor.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		vendor.ID,
		vendor.ContactPerson,
		vendor.Email,
		vendor.ContactNumber,
		vendor.Location,
		vendor.SIUPNumber,
		vendor.NIBNumber,
		vendor.NPWPNumber,
		vendor.UpdatedAt,
	).Scan(
		&vendor.ID,
		&vendor.ContactPerson,
		&vendor.Email,
		&vendor.ContactNumber,
		&vendor.Location,
		&vendor.SIUPNumber,
		&vendor.NIBNumber,
		&vendor.NPWPNumber,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
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

func (r *PostgresVendorRepository) UpdatePartial(vendor *domain.Vendor) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE vendors SET "

	if vendor.ContactPerson != nil {
		query += "contact_person = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.ContactPerson)
	}

	if vendor.Email != nil {
		query += "email = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.Email)
	}

	if vendor.ContactNumber != nil {
		query += "contact_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.ContactNumber)
	}

	if vendor.Location != nil {
		query += "location = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.Location)
	}

	if vendor.SIUPNumber != nil {
		query += "siup_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.SIUPNumber)
	}

	if vendor.NIBNumber != nil {
		query += "nib_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.NIBNumber)
	}

	if vendor.NPWPNumber != nil {
		query += "npwp_number = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *vendor.NPWPNumber)
	}

	vendor.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, vendor.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, vendor.ID)

	err = tx.QueryRow(query, args...).Scan(
		&vendor.ID,
		&vendor.ContactPerson,
		&vendor.Email,
		&vendor.ContactNumber,
		&vendor.Location,
		&vendor.SIUPNumber,
		&vendor.NIBNumber,
		&vendor.NPWPNumber,
		&vendor.CreatedAt,
		&vendor.UpdatedAt,
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

func (r *PostgresVendorRepository) Delete(vendor *domain.Vendor) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM vendors WHERE id = $1"

	result, err := tx.Exec(query, vendor.ID)
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
