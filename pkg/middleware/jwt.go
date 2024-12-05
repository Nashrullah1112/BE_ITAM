package middleware

import (
	"errors"
	"strconv"
	"time"

	userquery "github.com/banggibima/be-itam/modules/users/application/query"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Secret   string
	Expire   int
	Audience string
	Issuer   string
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func EncodedAccess(j *JWT, user *userquery.UserResponseDTO) (*jwt.Token, error) {
	id := strconv.Itoa(user.ID)

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			Issuer:    j.Issuer,
			Audience:  jwt.ClaimStrings{j.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role: *user.Role,
	}

	raw, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodedAccess(j *JWT, raw string) (*jwt.Token, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token telah kedaluwarsa")
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("signature token tidak valid")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return token, nil
}

func EncodedRefresh(j *JWT, user *userquery.UserResponseDTO) (*jwt.Token, error) {
	id := strconv.Itoa(user.ID)

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   id,
			Issuer:    j.Issuer,
			Audience:  jwt.ClaimStrings{j.Audience},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.Expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Role: *user.Role,
	}

	raw, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.Secret))
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodedRefresh(j *JWT, raw string) (*jwt.Token, error) {
	token, err := jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token telah kedaluwarsa")
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("signature token tidak valid")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return token, nil
}
