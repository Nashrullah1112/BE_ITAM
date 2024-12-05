package domain

type PostgresApplicationRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Application, error)
	FindByID(id int) (*Application, error)
	Create(application *Application) error
	Update(application *Application) error
	UpdatePartial(application *Application) error
	Delete(application *Application) error
}
