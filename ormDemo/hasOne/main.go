package main

import (
	"hydracode.com/ormDemo/common"
)

type company struct {
	ID          int
	CompanyName string
	User        user
}
type user struct {
	ID        int //default foreignkey
	Name      string
	UserCode  string
	companyID string //default reference key
}

type companyWithCustomForeignKey struct {
	ID          int
	CompanyName string
	User        user `gorm:"foreignKey:UserCode"`
}

type companyWithCustomReferenceKey struct {
	ID          int
	CompanyName string
	UserName    string
	User        user `gorm:"references:UserName"`
}

func main() {
	db := common.InitDB("./hasOne/gorm.db")
	common.MigrateDB(db, &company{}, &user{}, &companyWithCustomForeignKey{}, &companyWithCustomReferenceKey{})
	common.SetDB(db)
	common.CloseDB(db)
}
