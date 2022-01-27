package repoimp

import (
	"errors"
	"sync"

	"github.com/yiaw/go-example/go-clean-arc/model"
)

type userRepo struct{}

var userdb map[string]*model.User
var once sync.Once

func InitDB() {
	once.Do(func() {
		userdb = make(map[string]*model.User)
	})
}

func (u *userRepo) Create(user *model.User) error {
	userdb[user.Username] = user
	return nil
}

func (u *userRepo) Get(name string) (*model.User, error) {
	user, ok := userdb[name]
	if !ok {
		return nil, errors.New("not found user")
	}
	return user, nil
}

func (u *userRepo) Update(user *model.User) error {
	if _, ok := userdb[user.Username]; ok {
		delete(userdb, user.Username)
	}

	userdb[user.Username] = user
	return nil
}

func (u *userRepo) Delete(name string) error {
	_, ok := userdb[name]
	if !ok {
		return errors.New("not found user")
	}
	delete(userdb, name)
	return nil
}

func NewRepo() model.UserRepository {
	return &userRepo{}
}
