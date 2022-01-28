package model_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yiaw/go-example/go-clean-arc/model"
	"github.com/yiaw/go-example/go-clean-arc/model/mocks"
)

func TestUserCreate(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	t.Run("유저 생성 성공", func(t *testing.T) {

		mockuser := model.User{
			Username: "smlee123",
			Password: "smlee1234",
		}
		mockRepo.On("Create", &mockuser).Once().Return(nil)
		mockRepo.On("Get", mockuser.Username).Once().Return(nil, errors.New("not found user"))

		u := model.NewUser(mockRepo)

		err := u.UserCreate(&mockuser)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("아이디 5자리 미만 입력 시 실패", func(t *testing.T) {

		mockuser := model.User{
			Username: "sml",
			Password: "12345",
		}

		u := model.NewUser(mockRepo)

		err := u.UserCreate(&mockuser)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("아이디 16자리 초과 입력 시 실패", func(t *testing.T) {
		mockuser := model.User{
			Username: "yiaw1234567890123456",
			Password: "12345",
		}

		u := model.NewUser(mockRepo)

		err := u.UserCreate(&mockuser)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

/*
func TestUserGet(t *testing.T) {
	mockRepo := new(mocks.UserRepository)

	t.Run("error-filaed", func(t *testing.T) {
		mockRepo.On("Get").Return(nil, errors.New("not found user")).Once()

		u := model.NewUser(mockRepo)
		gu, err := u.UserGet("smlee")
		if err == nil {
			t.Errorf("user get fail %v", err)
		}
		assert.Error(t, err)
		assert.Nil(t, gu)

		mockRepo.AssertExpectations(t)
	})
}



func TestUserCreate_Success(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	mockuser := model.User{
		Username: "smlee123",
		Password: "smlee1234",
	}
	mockRepo.On("Create", &mockuser).Once().Return(nil)
	mockRepo.On("Get", mockuser.Username).Once().Return(nil, errors.New("not found user"))

	u := model.NewUser(mockRepo)

	err := u.UserCreate(&mockuser)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

*/
