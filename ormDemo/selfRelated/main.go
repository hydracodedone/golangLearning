package main

import (
	"hydracode.com/ormDemo/common"
)

type Employer struct {
	ID         int
	Name       string
	Code       int
	Supervisor *Employer
}

func main() {
	db := common.InitDB("./selfRelated/gorm.db")
	common.MigrateDB(db, &Employer{})
	common.SetDB(db)
	common.CloseDB(db)
}
