package utils

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func BcryptComparePassword(hashed, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	if err != nil {
		return errors.New("kata sandi tidak cocok")
	}
	return nil
}

func AsString(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func AsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

func AsFloat(key string, fallback float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
			return floatValue
		}
	}
	return fallback
}

func AsBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

func FormatValidationError(field validator.FieldError) string {
	lower := strings.ToLower(field.Field())
	param := field.Param()

	switch field.Tag() {
	case "required":
		return lower + " wajib diisi"
	case "email":
		return lower + " bukan alamat email yang valid"
	case "min":
		return lower + " harus memiliki setidaknya " + param + " karakter"
	case "max":
		return lower + " harus memiliki paling banyak " + param + " karakter"
	case "eqlower":
		return lower + " harus sama dengan " + param
	case "nelower":
		return lower + " tidak boleh sama dengan " + param
	case "gte":
		return lower + " harus lebih besar atau sama dengan " + param
	case "gt":
		return lower + " harus lebih besar dari " + param
	case "lte":
		return lower + " harus lebih kecil atau sama dengan " + param
	case "lt":
		return lower + " harus lebih kecil dari " + param
	case "uuid":
		return lower + " bukan UUID yang valid"
	default:
		return lower + " tidak valid"
	}
}

func FormatLogRequest(logger *logrus.Logger, method, uri string, status, latency int) {
	switch {
	case status >= 200 && status < 300:
		logger.Infof("request: method=" + method + " uri=" + uri + " status=" + strconv.Itoa(status) + " latency=" + strconv.Itoa(latency))
	case status >= 300 && status < 400:
		logger.Warnf("client error: method=" + method + " uri=" + uri + " status=" + strconv.Itoa(status) + " latency=" + strconv.Itoa(latency))
	case status >= 400 && status < 500:
		logger.Warnf("client error: method=" + method + " uri=" + uri + " status=" + strconv.Itoa(status) + " latency=" + strconv.Itoa(latency))
	case status >= 500:
		logger.Errorf("server error: method=" + method + " uri=" + uri + " status=" + strconv.Itoa(status) + " latency=" + strconv.Itoa(latency))
	default:
		logger.Infof("request: method=" + method + " uri=" + uri + " status=" + strconv.Itoa(status) + " latency=" + strconv.Itoa(latency))
	}
}
