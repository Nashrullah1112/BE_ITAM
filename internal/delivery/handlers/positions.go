package handlers

import (
	"net/http"
	"strconv"

	"github.com/banggibima/be-itam/modules/positions/application/command"
	"github.com/banggibima/be-itam/modules/positions/application/query"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/response"
	"github.com/banggibima/be-itam/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PositionHandler struct {
	Config                 *config.Config
	PositionCommandUsecase *command.PositionCommandUsecase
	PositionQueryUsecase   *query.PositionQueryUsecase
}

func NewPositionHandler(
	config *config.Config,
	positionCommandUsecase *command.PositionCommandUsecase,
	positionQueryUsecase *query.PositionQueryUsecase,
) *PositionHandler {
	return &PositionHandler{
		Config:                 config,
		PositionCommandUsecase: positionCommandUsecase,
		PositionQueryUsecase:   positionQueryUsecase,
	}
}

func (h *PositionHandler) FindAll(c *fiber.Ctx) error {
	meta := response.Meta{}

	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	meta.Page = max(0, page)
	meta.Size = max(0, size)

	meta.Sort = c.Query("sort")
	if meta.Sort == "" {
		meta.Sort = "created_at"
	}
	meta.Order = c.Query("order")
	if meta.Order == "" {
		meta.Order = "asc"
	}

	offset := (meta.Page - 1) * meta.Size
	limit := meta.Size
	sort := meta.Sort
	order := meta.Order

	filters := make(map[string]interface{})

	total, err := h.PositionQueryUsecase.CountAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	data, err := h.PositionQueryUsecase.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	meta.Total = total
	meta.Count = len(data)

	return c.Status(http.StatusOK).JSON(response.ResponsePagination(data, meta))
}

func (h *PositionHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	data, err := h.PositionQueryUsecase.FindByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *PositionHandler) Create(c *fiber.Ctx) error {
	dto := new(command.CreatePositionRequestDTO)

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

	data, err := h.PositionCommandUsecase.Create(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusCreated).JSON(response.ResponseData(data))
}

func (h *PositionHandler) Update(c *fiber.Ctx) error {
	dto := new(command.UpdatePositionRequestDTO)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	if err := c.BodyParser(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	dto.ID = id

	if err := validator.New().Struct(dto); err != nil {
		fields := []string{}
		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, v := range verr {
				fields = append(fields, utils.FormatValidationError(v))
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseValidation(fields))
	}

	data, err := h.PositionCommandUsecase.Update(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *PositionHandler) UpdatePartial(c *fiber.Ctx) error {
	dto := new(command.UpdatePartialPositionRequestDTO)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	if err := c.BodyParser(dto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	dto.ID = id

	if err := validator.New().Struct(dto); err != nil {
		fields := []string{}
		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, v := range verr {
				fields = append(fields, utils.FormatValidationError(v))
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseValidation(fields))
	}

	data, err := h.PositionCommandUsecase.UpdatePartial(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *PositionHandler) Delete(c *fiber.Ctx) error {
	dto := new(command.DeletePositionRequestDTO)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	dto.ID = id

	if err := validator.New().Struct(dto); err != nil {
		fields := []string{}
		if verr, ok := err.(validator.ValidationErrors); ok {
			for _, v := range verr {
				fields = append(fields, utils.FormatValidationError(v))
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(response.ResponseValidation(fields))
	}

	data, err := h.PositionCommandUsecase.Delete(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusNoContent).JSON(response.ResponseData(data))
}
