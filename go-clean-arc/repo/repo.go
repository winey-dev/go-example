package repo

import (
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
	user := userdb[name]
	return user, nil
}

func (u *userRepo) Update(user *model.User) error {
	userdb[user.Username] = user
	return nil
}

func (u *userRepo) Delete(name string) error {
	delete(userdb, name)
	return nil
}

func NewRepo() model.UserRepository {
	return &userRepo{}
}
