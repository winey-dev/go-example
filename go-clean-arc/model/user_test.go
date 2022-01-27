package model_test

import (
	"testing"

	"github.com/yiaw/go-example/go-clean-arc/model"
	"github.com/yiaw/go-example/go-clean-arc/repo"
)

func TestUserCreate(t *testing.T) {
	repo.InitDB()
	repoInterface := repo.NewRepo()
	u := model.NewUser(repoInterface)

	user := &model.User{
		Username: "smlee",
		Password: "smlee1234",
	}

	err := u.UserCreate(user)
	if err != nil {
		t.Errorf("user create fail %v", err)
	}
}

func TestUserGet(t *testing.T) {
	repo.InitDB()
	repoInterface := repo.NewRepo()
	u := model.NewUser(repoInterface)

	_, err := u.UserGet("smlee")
	if err == nil {
		t.Errorf("user get fail %v", err)
	}
}
