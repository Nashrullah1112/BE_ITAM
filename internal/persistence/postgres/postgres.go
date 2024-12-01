package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/banggibima/be-itam/pkg/config"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Client(config *config.Config, logger *logrus.Logger) (*sql.DB, error) {
	host := config.Postgres.Host
	port := config.Postgres.Port
	user := config.Postgres.Username
	password := config.Postgres.Password
	dbname := config.Postgres.Database

	sslmode := "disable"

	url := fmt.Sprint("postgres://", user, ":", password, "@", host, ":", port, "/", dbname, "?sslmode=", sslmode)

	client, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	if err := Connect(client); err != nil {
		return nil, err
	}

	if err := Migration(client, logger); err != nil {
		return nil, err
	}

	if err := Seed(client, logger); err != nil {
		return nil, err
	}

	return client, nil
}

func Connect(client *sql.DB) error {
	err := client.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Migration(db *sql.DB, logger *logrus.Logger) error {
	query := "CREATE TABLE IF NOT EXISTS migrations (id UUID PRIMARY KEY, name TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)"

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("migrations")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		count := 0
		query = "SELECT COUNT(*) FROM migrations WHERE name = $1"

		err := db.QueryRow(query, file.Name()).Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		path := filepath.Join("migrations", file.Name())
		sql, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(sql))
		if err != nil {
			return err
		}

		id := uuid.New()
		query = "INSERT INTO migrations (id, name) VALUES ($1, $2)"

		_, err = db.Exec(query, id, file.Name())
		if err != nil {
			return err
		}

		logger.Infof("migration applied: %s", file.Name())
	}

	return nil
}

func Seed(db *sql.DB, logger *logrus.Logger) error {
	query := "CREATE TABLE IF NOT EXISTS seeders (id UUID PRIMARY KEY, name TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)"

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("seeders")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		count := 0
		query = "SELECT COUNT(*) FROM seeders WHERE name = $1"

		err := db.QueryRow(query, file.Name()).Scan(&count)
		if err != nil {
			return err
		}

		if count > 0 {
			continue
		}

		path := filepath.Join("seeders", file.Name())
		sql, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(sql))
		if err != nil {
			return err
		}

		id := uuid.New()
		query = "INSERT INTO seeders (id, name) VALUES ($1, $2)"

		_, err = db.Exec(query, id, file.Name())
		if err != nil {
			return err
		}

		logger.Infof("seeder applied: %s", file.Name())
	}

	return nil
}
