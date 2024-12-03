package domain

type PostgresLicenseRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*License, error)
	FindByID(id int) (*License, error)
	Create(license *License) error
	Update(license *License) error
	UpdatePartial(license *License) error
	Delete(license *License) error
}
