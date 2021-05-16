package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"hydracode.com/ormDemo/common"
	. "hydracode.com/ormDemo/model"
)

func queryDBAllCompany(db *gorm.DB) {
	var companyInstance []Company
	db.Debug().Find(&companyInstance)
	fmt.Printf("The result is %#v\n", companyInstance)
}
func queryDBCompanyByUser(db *gorm.DB) {
	var userInstance User
	var companyInstanceSingle Company
	companyInstanceSingle = Company{}
	userInstance = User{
		Model: gorm.Model{
			ID: 1,
		},
	}
	err := db.Debug().Model(&userInstance).Association("Company").Find(&companyInstanceSingle)
	if err != nil {
		log.Fatalf("Query DB Fail :<%s>\n", err.Error())
	}
	fmt.Printf("The result is %+v\n", companyInstanceSingle)
}
func main() {

	db := common.InitDB()
	common.SetDB(db)
	common.MigrateDB(db, &User{}, &User2{}, &User3{}, &User4{}, &Company{}, &Company2{}, &Company3{}, &Company4{})
	queryDBAllCompany(db)
	queryDBCompanyByUser(db)
	common.CleanDB(db)
	common.CloseDB(db)
}
