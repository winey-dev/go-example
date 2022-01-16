package model

type UserRepository interface {
	Create(u *User) error
	Get(name string) (*User, error)
	Update(u *User) error
	Delete(name string) error
}
