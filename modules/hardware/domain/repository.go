package domain

type PostgresHardwareRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Hardware, error)
	FindByID(id int) (*Hardware, error)
	Create(hardware *Hardware) error
	Update(hardware *Hardware) error
	UpdatePartial(hardware *Hardware) error
	Delete(hardware *Hardware) error
}
