package main

import (
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID         int
	UName      string
	CreditCard CreditCard
}
type CreditCard struct {
	ID     int
	CName  string
	UserID int
}

func main() {
	db := common.InitDB("./hasOne/gorm.db")
	common.MigrateDB(db, &CreditCard{}, &User{})
	common.SetDB(db)
	common.CloseDB(db)
}
