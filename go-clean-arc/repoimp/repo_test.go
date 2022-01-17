package repoimp_test

import (
	"testing"

	"github.com/yiaw/go-example/go-clean-arc/model"
	"github.com/yiaw/go-example/go-clean-arc/repoimp"
)

func TestCreate(t *testing.T) {
	u := model.User{
		Username: "test",
		Password: "test",
	}
	err := repoimp.Create(&u)
	if err != nil {
		t.Errorf("Create fail err=%v\n", err)
	}
}

func TestGet(t *testing.T) {
	u := model.User{
		Username: "test",
		Password: "test",
	}
	err := repoimp.Create(&u)
	if err != nil {
		t.Errorf("Create fail err=%v\n", err)
		return
	}

	user, err := repoimp.Get("test")
	if err != nil {
		t.Errorf("Get fail err=%v\n", err)
		return
	}

	if user.Username != "test" || user.Password != "test" {
		t.Errorf("Get Result Value Fail\n")
		return
	}
}
