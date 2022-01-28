package model

import (
	"fmt"
)

// Domain 정의
type User struct {
	Username string
	Password string
}

//go:generate mokery --name UserRepository --case undersocre --inpakcage
type UserRepository interface {
	Create(u *User) error
	Get(name string) (*User, error)
	Update(u *User) error
	Delete(name string) error
}

type UserInterface interface {
	UserCreate(u *User) error
	UserGet(name string) (*User, error)
	UserUpdate(u *User) error
	UserDelete(name string) error
}

// Usecase 구현
type user struct {
	repo UserRepository
}

func (u *user) UserCreate(user *User) error {
	if len(user.Username) < 5 || len(user.Username) > 16 {
		return fmt.Errorf("invalid username (%s)", user.Username)
	}

	if len(user.Password) < 8 || len(user.Password) > 32 {
		return fmt.Errorf("invalid password")
	}

	gu, _ := u.repo.Get(user.Username)
	if gu != nil {
		return fmt.Errorf("already user (%s)", user.Username)
	}

	return u.repo.Create(user)
}

func (u *user) UserGet(name string) (*User, error) {
	if len(name) < 5 || len(name) > 16 {
		return nil, fmt.Errorf("invalid username (%s)", name)
	}
	return u.repo.Get(name)
}

func (u *user) UserUpdate(user *User) error {
	if len(user.Username) < 5 || len(user.Username) > 16 {
		return fmt.Errorf("invalid username (%s)", user.Username)
	}

	if len(user.Password) < 8 || len(user.Password) > 32 {
		return fmt.Errorf("invalid password")
	}

	return u.repo.Update(user)
}

func (u *user) UserDelete(name string) error {
	if len(name) < 5 || len(name) > 16 {
		return fmt.Errorf("invalid username (%s)", name)
	}
	return u.repo.Delete(name)
}

func NewUser(repo UserRepository) UserInterface {
	return &user{repo: repo}
}
