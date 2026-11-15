package domain

// UserRepository adalah kontrak yang diminta oleh domain/usecase
type UserRepository interface {
	Create(user *User) error
	FindByID(id string) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}
