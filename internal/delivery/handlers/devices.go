package handlers

import (
	"net/http"
	"strconv"

	"github.com/banggibima/be-itam/modules/devices/application/command"
	"github.com/banggibima/be-itam/modules/devices/application/query"
	"github.com/banggibima/be-itam/pkg/config"
	"github.com/banggibima/be-itam/pkg/response"
	"github.com/banggibima/be-itam/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type DeviceHandler struct {
	Config               *config.Config
	DeviceCommandUsecase *command.DeviceCommandUsecase
	DeviceQueryUsecase   *query.DeviceQueryUsecase
}

func NewDeviceHandler(
	config *config.Config,
	deviceCommandUsecase *command.DeviceCommandUsecase,
	deviceQueryUsecase *query.DeviceQueryUsecase,
) *DeviceHandler {
	return &DeviceHandler{
		Config:               config,
		DeviceCommandUsecase: deviceCommandUsecase,
		DeviceQueryUsecase:   deviceQueryUsecase,
	}
}

func (h *DeviceHandler) FindAll(c *fiber.Ctx) error {
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

	total, err := h.DeviceQueryUsecase.CountAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	data, err := h.DeviceQueryUsecase.FindAll(offset, limit, sort, order, filters)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	meta.Total = total
	meta.Count = len(data)

	return c.Status(http.StatusOK).JSON(response.ResponsePagination(data, meta))
}

func (h *DeviceHandler) FindByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ResponseError(err))
	}

	data, err := h.DeviceQueryUsecase.FindByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *DeviceHandler) Create(c *fiber.Ctx) error {
	dto := new(command.CreateDeviceRequestDTO)

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

	data, err := h.DeviceCommandUsecase.Create(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusCreated).JSON(response.ResponseData(data))
}

func (h *DeviceHandler) Update(c *fiber.Ctx) error {
	dto := new(command.UpdateDeviceRequestDTO)

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

	data, err := h.DeviceCommandUsecase.Update(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *DeviceHandler) UpdatePartial(c *fiber.Ctx) error {
	dto := new(command.UpdatePartialDeviceRequestDTO)

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

	data, err := h.DeviceCommandUsecase.UpdatePartial(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusOK).JSON(response.ResponseData(data))
}

func (h *DeviceHandler) Delete(c *fiber.Ctx) error {
	dto := new(command.DeleteDeviceRequestDTO)

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

	data, err := h.DeviceCommandUsecase.Delete(dto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ResponseError(err))
	}

	return c.Status(http.StatusNoContent).JSON(response.ResponseData(data))
}
