package main

import (
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID         int
	UName      string
	CreditCard CreditCard `gorm:"foreignKey:UserRefer"`
}
type CreditCard struct {
	ID        int
	CName     string
	UserRefer int
}

func main() {
	db := common.InitDB("./hasOneForeignKey/gorm.db")
	common.MigrateDB(db, &CreditCard{}, &User{})
	common.SetDB(db)
	common.CloseDB(db)
}
