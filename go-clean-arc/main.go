package main

import (
	"log"

	"github.com/yiaw/go-example/go-clean-arc/model"
	"github.com/yiaw/go-example/go-clean-arc/repo"
)

func main() {
	repo.InitDB()
	repoInterface := repo.NewRepo()
	u := model.NewUser(repoInterface)

	err := u.UserCreate(&model.User{Username: "smlee", Password: "smlee1234"})
	if err != nil {
		log.Fatalf("user create fail")
	}

	user, err := u.UserGet("smlee")
	if err != nil {
		log.Fatalf("user get fail")
	}
	log.Printf("user : %v", user)

	user.Password = "qwe582"

	err = u.UserUpdate(user)
	if err != nil {
		log.Fatalf("user update fail")
	}

	user2, err := u.UserGet("smlee")
	if err != nil {
		log.Fatalf("user get fail")
	}
	log.Printf("user2 : %v", user2)

	err = u.UserDelete("smlee")
	if err != nil {
		log.Fatalf("user delete fail")
	}

	user3, err := u.UserGet("smlee")
	if err != nil {
		log.Fatalf("user get fail")
	}
	log.Printf("user3 : %v", user3)
}
