package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dns := "root:root@tcp(localhost:3306)/gorm_test?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatalf("gorm.Open failed .. err=%v\n", err)
	}

	// Table 자동 생성 및 변경 ??
	db.AutoMigrate(&User{})

	log.Printf("DB Connection Succ. dns=%s\n", dns)

	// Table에 User Data 추가
	result := db.Create(&User{
		ID:                 "smlee1234",
		Password:           "Password",
		Name:               "lee",
		Gender:             "Male",
		Email:              "lsmin0703@naver.com",
		PasswordExpireTime: SetPasswordExpireTime(),
	})

	if result.Error != nil {
		log.Fatalf("db.Create failed .. err=%v\n", err)
	}

	// Table에서 데이터 읽기
	var user User

	// db.First(&userm "id = ?", "smlee") eqauls
	result = db.Where("id = ?", "smlee").First(&user)
	if result.Error != nil {
		log.Fatalf("db.Where.First failed .. err=%v\n", result.Error)
	}

	log.Printf("First Succ=%v\n", user)

	var users []User
	result = db.Find(&users)
	if result.Error != nil {
		log.Fatalf("db.Where.Find failed .. err=%v\n", result.Error)
	}
	log.Printf("Find Succ=%v\n", users)

	// User Update
	db.Model(&user).Update("name", "hello")

	// User Delete
	db.Delete(&user)

}
