package domain

type PostgresUserRepository interface {
	CountAll() (int, error)
	FindAll(offset, limit int, sort, order string, filters map[string]interface{}) ([]*User, error)
	FindByID(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	UpdatePartial(user *User) error
	Delete(user *User) error
}
