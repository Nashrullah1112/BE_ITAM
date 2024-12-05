package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/divisions/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresDivisionRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresDivisionRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresDivisionRepository {
	return &PostgresDivisionRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresDivisionRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM divisions"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresDivisionRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Division, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM divisions"

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

	divisions := []*domain.Division{}
	for rows.Next() {
		division := domain.Division{}
		err := rows.Scan(
			&division.ID,
			&division.Name,
			&division.CreatedAt,
			&division.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		divisions = append(divisions, &division)
	}

	return divisions, nil
}

func (r *PostgresDivisionRepository) FindByID(id int) (*domain.Division, error) {
	query := "SELECT * FROM divisions WHERE id = $1"

	division := &domain.Division{}

	err := r.DB.QueryRow(query, id).Scan(
		&division.ID,
		&division.Name,
		&division.CreatedAt,
		&division.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return division, nil
}

func (r *PostgresDivisionRepository) Create(division *domain.Division) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO divisions (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING *"

	division.CreatedAt = time.Now()
	division.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		division.Name,
		division.CreatedAt,
		division.UpdatedAt,
	).Scan(
		&division.ID,
		&division.Name,
		&division.CreatedAt,
		&division.UpdatedAt,
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

func (r *PostgresDivisionRepository) Update(division *domain.Division) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE divisions SET name = $2, updated_at = $3 WHERE id = $1 RETURNING *"

	division.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		division.ID,
		division.Name,
		division.UpdatedAt,
	).Scan(
		&division.ID,
		&division.Name,
		&division.CreatedAt,
		&division.UpdatedAt,
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

func (r *PostgresDivisionRepository) UpdatePartial(division *domain.Division) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE divisions SET "

	if division.Name != nil {
		query += "name = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *division.Name)
	}

	division.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, division.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, division.ID)

	err = tx.QueryRow(query, args...).Scan(
		&division.ID,
		&division.Name,
		&division.CreatedAt,
		&division.UpdatedAt,
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

func (r *PostgresDivisionRepository) Delete(division *domain.Division) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM divisions WHERE id = $1"

	result, err := tx.Exec(query, division.ID)
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
