package main

import (
	"time"
)

type Gender string

const (
	Male   = Gender("M")
	FeMale = Gender("F")
)

var emailList = [...]string{"naver.com", "gmail.com", "daum.net"}

type User struct {
	ID                 string
	Password           string
	Name               string
	Email              string
	Gender             string
	PasswordExpireTime time.Time
}

func SetPasswordExpireTime() time.Time {
	t := time.Now()
	return t.AddDate(1, 0, 0)
}
