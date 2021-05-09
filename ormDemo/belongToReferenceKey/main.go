package main

import (
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID    int
	UName string
	UCode int
}
type CreditCard struct {
	ID     int
	CName  string
	UserID int
	User   User `gorm:"referenceKey:UCode"`
}

func main() {
	db := common.InitDB("./belongToReferenceKey/gorm.db")
	common.MigrateDB(db, &CreditCard{}, &User{})
	common.SetDB(db)
	common.CloseDB(db)
}
