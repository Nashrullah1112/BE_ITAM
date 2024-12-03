package domain

type PostgresDeviceRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Device, error)
	FindByID(id int) (*Device, error)
	Create(device *Device) error
	Update(device *Device) error
	UpdatePartial(device *Device) error
	Delete(device *Device) error
}
