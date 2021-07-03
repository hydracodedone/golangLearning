package main

import (
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID    int
	UName string
}
type CreditCard struct {
	ID     int
	CName  string
	UserID int
	User   User
}

func main() {
	db := common.InitDB("./belongTo/gorm.db")
	common.MigrateDB(db, &CreditCard{}, &User{})
	common.SetDB(db)
	common.CloseDB(db)
}
