package domain

type PostgresVendorRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Vendor, error)
	FindByID(id int) (*Vendor, error)
	Create(vendor *Vendor) error
	Update(vendor *Vendor) error
	UpdatePartial(vendor *Vendor) error
	Delete(vendor *Vendor) error
}
