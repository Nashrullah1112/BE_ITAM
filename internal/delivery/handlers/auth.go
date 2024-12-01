package handlers

import (
	"net/http"

	"github.com/banggibima/be-itam/modules/auth/application/command"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/response"
	"github.com/banggibima/be-itam/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Config             *config.Config
	AuthCommandUsecase *command.AuthCommandUsecase
}

func NewAuthHandler(
	config *config.Config,
	authCommandUsecase *command.AuthCommandUsecase,
) *AuthHandler {
	return &AuthHandler{
		Config:             config,
		AuthCommandUsecase: authCommandUsecase,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	dto := new(command.AuthRegisterRequestDTO)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	if err := validator.New().Struct(dto); err != nil {
		fields := []string{}
		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, v := range verr {
				fields = append(fields, utils.FormatValidationError(v))
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseValidation(fields))
	}

	data, err := h.AuthCommandUsecase.Register(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusCreated).JSON(response.ResponseData(data))
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	dto := new(command.AuthLoginRequestDTO)

	if err := c.BodyParser(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	if err := validator.New().Struct(dto); err != nil {
		fields := []string{}
		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, v := range verr {
				fields = append(fields, utils.FormatValidationError(v))
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseValidation(fields))
	}

	data, err := h.AuthCommandUsecase.Login(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}
