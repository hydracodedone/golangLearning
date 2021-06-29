package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID       int //default foreignkey
	Name     string
	UserCode int

	CompanyID int //default reference key
}

type Company struct {
	ID   int
	Name string
	User User
}

//CompanyWithCustomForeignKey Abandon User Table  CompanyID Column
type CompanyWithCustomForeignKey struct {
	ID   int
	Name string
	User User `gorm:"foreignKey:UserCode"`
}

type CompanyWithCustomReferenceKey struct {
	ID          int
	CompanyName string
	UID         int
	User        User `gorm:"references:UID"`
}

func insertCompany(db *gorm.DB) {
	insertCompany := new(Company)
	insertCompany.Name = "cmp1"
	insertCompany.User = User{
		Name: "u1",
	}
	db.Debug().Create(insertCompany)
	fmt.Printf("The Insert Result Is %+v\n", insertCompany)

}
func preloadQueryCompany(db *gorm.DB) {
	queryCompany := new(Company)
	db.Debug().Model(queryCompany).Preload("User").First(queryCompany)
	fmt.Printf("The Query Result Is %+v\n", queryCompany)
}
func insertCompanyWithCustomForeignKey(db *gorm.DB) {
	insertCompany := new(CompanyWithCustomForeignKey)
	insertCompany.Name = "cmp2"
	insertCompany.User = User{
		Name:     "u2",
		UserCode: 1002,
	}
	db.Debug().Create(insertCompany)
	fmt.Printf("The Insert Result Is %+v\n", insertCompany)

}
func preloadQueryCompanyWithCustomForeignKey(db *gorm.DB) {
	queryCompany := new(CompanyWithCustomForeignKey)
	db.Debug().Model(queryCompany).Preload("User").Where(map[string]interface{}{"company_name": "cmp2"}).Find(queryCompany)
	fmt.Printf("The Query Result Is %+v\n", queryCompany)
}
func preloadQueryCompanyWithCustomReferenceKey(db *gorm.DB) {
	queryCompany := new(CompanyWithCustomReferenceKey)
	db.Debug().Model(queryCompany).Preload("User").Where(map[string]interface{}{"company_name": "cmp1"}).Find(queryCompany)
	fmt.Printf("The Query Result Is %+v\n", queryCompany)
}
func relatedQueryCompanyWithCustomReferenceKey(db *gorm.DB) {
	queryCompany := new(CompanyWithCustomReferenceKey)
	db.Debug().Model(queryCompany).Where(map[string]interface{}{"company_name": "cmp1"}).Find(queryCompany)
	db.Debug().Model(queryCompany).Related(&queryCompany.User, "UID")
	fmt.Printf("The Query Result Is %+v\n", queryCompany)
}

func main() {
	db := common.InitDB("./hasOne/gorm.db")
	common.MigrateDB(db, &Company{}, &User{}, &CompanyWithCustomForeignKey{}, &CompanyWithCustomReferenceKey{})
	common.SetDB(db)
	insertCompanyWithCustomForeignKey(db)
	common.CloseDB(db)
}
