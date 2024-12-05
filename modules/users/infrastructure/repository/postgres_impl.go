package repository

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/banggibima/be-itam/modules/users/domain"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PostgresUserRepository struct {
	Config *config.Config
	DB     *sql.DB
	Logger *logrus.Logger
}

func NewPostgresUserRepository(
	config *config.Config,
	db *sql.DB,
	logger *logrus.Logger,
) *PostgresUserRepository {
	return &PostgresUserRepository{
		Config: config,
		DB:     db,
		Logger: logger,
	}
}

func (r *PostgresUserRepository) CountAll() (int, error) {
	query := "SELECT COUNT(*) FROM users"
	total := 0

	err := r.DB.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (r *PostgresUserRepository) FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*domain.User, error) {
	args := []interface{}{}
	conditions := []string{}

	query := "SELECT * FROM users"

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

	users := []*domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.NIP,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.JoinDate,
			&user.DivisionID,
			&user.PositionID,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (r *PostgresUserRepository) FindByID(id int) (*domain.User, error) {
	query := "SELECT * FROM users WHERE id = $1"

	user := &domain.User{}

	err := r.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.NIP,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.JoinDate,
		&user.DivisionID,
		&user.PositionID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByEmail(email string) (*domain.User, error) {
	query := "SELECT * FROM users WHERE email = $1"

	user := &domain.User{}

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.NIP,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.JoinDate,
		&user.DivisionID,
		&user.PositionID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("data tidak ditemukan")
		}
		return nil, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "INSERT INTO users (nip, name, email, password, role, join_date, division_id, position_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *"

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		user.NIP,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.JoinDate,
		user.DivisionID,
		user.PositionID,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(
		&user.ID,
		&user.NIP,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.JoinDate,
		&user.DivisionID,
		&user.PositionID,
		&user.CreatedAt,
		&user.UpdatedAt,
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

func (r *PostgresUserRepository) Update(user *domain.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "UPDATE users SET nip = $2, name = $3, email = $4, password = $5, role = $6, join_date = $7, division_id = $8, position_id = $9, updated_at = $10 WHERE id = $1 RETURNING *"

	user.UpdatedAt = time.Now()

	err = tx.QueryRow(
		query,
		user.ID,
		user.NIP,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
		user.JoinDate,
		user.DivisionID,
		user.PositionID,
		user.UpdatedAt,
	).Scan(
		&user.ID,
		&user.NIP,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.JoinDate,
		&user.DivisionID,
		&user.PositionID,
		&user.CreatedAt,
		&user.UpdatedAt,
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

func (r *PostgresUserRepository) UpdatePartial(user *domain.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	args := []interface{}{}
	query := "UPDATE users SET "

	if user.NIP != nil {
		query += "nip = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.NIP)
	}

	if user.Name != nil {
		query += "name = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.Name)
	}

	if user.Email != nil {
		query += "email = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.Email)
	}

	if user.Password != nil {
		query += "password = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.Password)
	}

	if user.Role != nil {
		query += "role = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.Role)
	}

	if user.JoinDate != nil {
		query += "join_date = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.JoinDate)
	}

	if user.DivisionID != nil {
		query += "division_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.DivisionID)
	}

	if user.PositionID != nil {
		query += "position_id = $" + strconv.Itoa(len(args)+1) + ", "
		args = append(args, *user.PositionID)
	}

	user.UpdatedAt = time.Now()
	query += "updated_at = $" + strconv.Itoa(len(args)+1)
	args = append(args, user.UpdatedAt)

	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " RETURNING *"
	args = append(args, user.ID)

	err = tx.QueryRow(query, args...).Scan(
		&user.ID,
		&user.NIP,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.JoinDate,
		&user.DivisionID,
		&user.PositionID,
		&user.CreatedAt,
		&user.UpdatedAt,
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

func (r *PostgresUserRepository) Delete(user *domain.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := "DELETE FROM users WHERE id = $1"

	result, err := tx.Exec(query, user.ID)
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
