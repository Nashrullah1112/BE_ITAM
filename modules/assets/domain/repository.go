package domain

type PostgresAssetRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*Asset, error)
	FindByID(id int) (*Asset, error)
	Create(asset *Asset) error
	Update(asset *Asset) error
	UpdatePartial(asset *Asset) error
	Delete(asset *Asset) error
}
