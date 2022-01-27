package model

type userusecase struct {
	repo UserRepository
}

type UserUsecase interface {
	UserCreate(u *User) error
	UserGet(name string) (*User, error)
	UserUpdate(u *User) error
	UserDelete(name string) error
}

func (u *userusecase) UserCreate(user *User) error {
	return u.repo.Create(user)
}

func (u *userusecase) UserGet(name string) (*User, error) {
	return u.repo.Get(name)
}

func (u *userusecase) UserUpdate(user *User) error {
	return u.repo.Update(user)
}

func (u *userusecase) UserDelete(name string) error {
	return u.repo.Delete(name)
}

func NewUseCase(repo UserRepository) UserUsecase {
	return &userusecase{repo: repo}
}
