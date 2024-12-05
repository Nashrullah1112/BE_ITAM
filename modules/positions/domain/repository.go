package domain

type PostgresPositionRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Position, error)
	FindByID(id int) (*Position, error)
	Create(position *Position) error
	Update(position *Position) error
	UpdatePartial(position *Position) error
	Delete(position *Position) error
}
