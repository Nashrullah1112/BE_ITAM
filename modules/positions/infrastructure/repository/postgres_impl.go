package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/positions/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresPositionRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresPositionRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresPositionRepository {
	return &PostgresPositionRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresPositionRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM positions"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresPositionRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.Position, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM positions"

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

	positions := []*domain.Position{}
	for rows.Next() {
		position := domain.Position{}
		err := rows.Scan(
			&position.ID,
			&position.Name,
			&position.CreatedAt,
			&position.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		positions = append(positions, &position)
	}

	return positions, nil
}

func (r *PostgresPositionRepository) FindByID(id int) (*domain.Position, error) {
	query := "SELECT * FROM positions WHERE id = $1"

	position := &domain.Position{}

	err := r.DB.QueryRow(query, id).Scan(
		&position.ID,
		&position.Name,
		&position.CreatedAt,
		&position.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return position, nil
}

func (r *PostgresPositionRepository) Create(position *domain.Position) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO positions (name, created_at, updated_at) VALUES ($1, $2, $3) RETURNING *"

	position.CreatedAt = time.Now()
	position.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		position.Name,
		position.CreatedAt,
		position.UpdatedAt,
	).Scan(
		&position.ID,
		&position.Name,
		&position.CreatedAt,
		&position.UpdatedAt,
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

func (r *PostgresPositionRepository) Update(position *domain.Position) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE positions SET name = $2, updated_at = $3 WHERE id = $1 RETURNING *"

	position.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		position.ID,
		position.Name,
		position.UpdatedAt,
	).Scan(
		&position.ID,
		&position.Name,
		&position.CreatedAt,
		&position.UpdatedAt,
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

func (r *PostgresPositionRepository) UpdatePartial(position *domain.Position) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE positions SET "

	if position.Name != nil {
		query += "name = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *position.Name)
	}

	position.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, position.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, position.ID)

	err = tx.QueryRow(query, args...).Scan(
		&position.ID,
		&position.Name,
		&position.CreatedAt,
		&position.UpdatedAt,
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

func (r *PostgresPositionRepository) Delete(position *domain.Position) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM positions WHERE id = $1"

	result, err := tx.Exec(query, position.ID)
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
