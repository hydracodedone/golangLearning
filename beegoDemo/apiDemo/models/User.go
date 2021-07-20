package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	db = InitDB("./gorm.db")
	SetDB(db)
	MigrateDB(db, &User{})
}

type User struct {
	gorm.Model
	Name string
	Age  int
}

func InsertUser(name string, age int) *User {
	defer func() {
		CloseDB(db)
	}()
	user := new(User)
	user.Name = name
	user.Age = age
	err := db.Create(user).Error
	if err != nil {
		log.Fatalf("Insert User Fail:%s\n", err.Error())
	}
	return user
}
