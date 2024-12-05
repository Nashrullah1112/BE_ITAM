package domain

type PostgresDivisionRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Division, error)
	FindByID(id int) (*Division, error)
	Create(division *Division) error
	Update(division *Division) error
	UpdatePartial(division *Division) error
	Delete(division *Division) error
}
