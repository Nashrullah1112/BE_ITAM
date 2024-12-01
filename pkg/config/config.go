package config

import (
	"github.com/banggibima/be-itam/pkg/utils"
	"github.com/joho/godotenv"
)

type App struct {
	Name  string
	Env   string
	Debug bool
	Port  int
}

type JWT struct {
	SecretAccess  string
	SecretRefresh string
	ExpireAccess  int
	ExpireRefresh int
	Audience      string
	Issuer        string
}

type Postgres struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type Minio struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type Config struct {
	App      App
	JWT      JWT
	Postgres Postgres
	Minio    Minio
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil && utils.AsString("APP_ENV", "development") != "production" {
		return nil, err
	}

	config := &Config{
		App: App{
			Name:  utils.AsString("APP_NAME", "be-itam"),
			Env:   utils.AsString("APP_ENV", "development"),
			Debug: utils.AsBool("APP_DEBUG", true),
			Port:  utils.AsInt("APP_PORT", 8080),
		},
		JWT: JWT{
			SecretAccess:  utils.AsString("JWT_SECRET_ACCESS", "your-secret-access"),
			SecretRefresh: utils.AsString("JWT_SECRET_REFRESH", "your-secret-refresh"),
			ExpireAccess:  utils.AsInt("JWT_EXPIRE_ACCESS", 3600),
			ExpireRefresh: utils.AsInt("JWT_EXPIRE_REFRESH", 86400),
			Audience:      utils.AsString("JWT_AUDIENCE", "be-itam"),
			Issuer:        utils.AsString("JWT_ISSUER", "be-itam"),
		},
		Postgres: Postgres{
			Host:     utils.AsString("POSTGRES_HOST", "localhost"),
			Port:     utils.AsInt("POSTGRES_PORT", 5432),
			Username: utils.AsString("POSTGRES_USERNAME", "postgres"),
			Password: utils.AsString("POSTGRES_PASSWORD", ""),
			Database: utils.AsString("POSTGRES_DATABASE", "be_itam"),
		},
		Minio: Minio{
			Endpoint:        utils.AsString("MINIO_ENDPOINT", "localhost:9000"),
			AccessKeyID:     utils.AsString("MINIO_ACCESS_KEY_ID", "minio-access-key"),
			SecretAccessKey: utils.AsString("MINIO_SECRET_ACCESS_KEY", "minio-secret-key"),
			UseSSL:          utils.AsBool("MINIO_USE_SSL", false),
		},
	}

	return config, nil
}
