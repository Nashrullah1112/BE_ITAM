package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/banggibima/be-itam/pkg/config"
	jwtmiddleware "github.com/banggibima/be-itam/pkg/middleware"
	"github.com/banggibima/be-itam/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTMiddleware struct {
	Config *config.Config
}

func NewJWTMiddleware(
	config *config.Config,
) *JWTMiddleware {
	return &JWTMiddleware{
		Config: config,
	}
}

func (j *JWTMiddleware) Authentication(c *fiber.Ctx) error {
	authorization := c.Get("Authorization")
	if authorization == "" {
		return c.Status(http.StatusUnauthorized).JSON(response.ResponseError(errors.New("header otorisasi diperlukan")))
	}

	if !strings.HasPrefix(authorization, "Bearer ") {
		return c.Status(http.StatusUnauthorized).JSON(response.ResponseError(errors.New("header otorisasi harus berupa token bearer")))
	}

	token := strings.TrimPrefix(authorization, "Bearer ")

	claims, err := jwtmiddleware.DecodedAccess(&jwtmiddleware.JWT{
		Secret: j.Config.JWT.SecretAccess,
	}, token)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(response.ResponseError(err))
	}

	c.Locals("claims", claims)

	return c.Next()
}

func (j *JWTMiddleware) Authorization(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("claims").(*jwt.Token).Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(http.StatusUnauthorized).JSON(response.ResponseError(errors.New("klaim token tidak valid")))
		}

		for _, role := range roles {
			if claims["role"].(string) == role {
				return c.Next()
			}
		}

		return c.Status(http.StatusForbidden).JSON(response.ResponseError(errors.New("tidak diizinkan mengakses resource ini")))
	}
}
