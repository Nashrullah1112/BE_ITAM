package response

import (
	"errors"
)

type Meta struct {
	Page  int    `json:"page"`
	Size  int    `json:"size"`
	Count int    `json:"count"`
	Total int    `json:"total"`
	Sort  string `json:"sort"`
	Order string `json:"order"`
}

type Data struct {
	Data interface{} `json:"data"`
}

type Pagination struct {
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Validation struct {
	Error   string   `json:"error"`
	Details []string `json:"details"`
}

type Error struct {
	Error string `json:"error"`
}

func ResponseData(data interface{}) Data {
	wrapper := Data{
		Data: data,
	}

	return wrapper
}

func ResponsePagination(data interface{}, meta Meta) Pagination {
	wrapper := Pagination{
		Data: data,
		Meta: meta,
	}

	return wrapper
}

func ResponseValidation(details []string) Validation {
	wrapper := Validation{
		Error:   errors.New("kesalahan validasi").Error(),
		Details: details,
	}

	return wrapper
}

func ResponseError(err error) Error {
	wrapper := Error{
		Error: err.Error(),
	}

	return wrapper
}
