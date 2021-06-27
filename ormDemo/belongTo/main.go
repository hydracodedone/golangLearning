package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"hydracode.com/ormDemo/common"
)

type company struct {
	ID   int
	Name string
	Code int
}
type user struct {
	Name      string
	CompanyID int
	Company   company
}
type userWithCustomForeignKey struct {
	Name         string
	CompanyRefer int
	Company      company `gorm:"foreignKey:CompanyRefer"` // 使用 CompanyRefer 作为外键
}
type userWithCustomReferenceKey struct {
	Name      string
	CompanyID int
	Company   company `gorm:"references:Code"` // 使用 Code 作为引用
}

func preloadQueryUser(db *gorm.DB) {
	queryUser := new(user)
	db.Debug().Model(queryUser).Preload("Company").First(queryUser)
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}
func preloadQueryUserWithCustomForeignKey(db *gorm.DB) {
	queryUser := new(userWithCustomForeignKey)
	db.Debug().Model(queryUser).Preload("Company").First(queryUser)
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}
func preloadQueryUserWithCustomReferenceKey(db *gorm.DB) {
	queryUser := new(userWithCustomReferenceKey)
	db.Debug().Model(queryUser).Preload("Company").First(queryUser) // reference change can not use preload
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}
func relatedQueryUser(db *gorm.DB) {
	queryUser := new(user)
	db.Debug().Model(queryUser).First(queryUser)
	db.Debug().Model(queryUser).Related(&queryUser.Company)
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}
func relatedQueryUserWithCustomForeignKey(db *gorm.DB) {
	queryUser := new(userWithCustomForeignKey)
	db.Debug().Model(queryUser).First(queryUser)
	db.Debug().Model(queryUser).Related(&queryUser.Company, "CompanyRefer")
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}
func relatedQueryUserWithCustomReferenceKey(db *gorm.DB) {
	queryUser := new(userWithCustomReferenceKey)
	db.Debug().Model(queryUser).Where(map[string]interface{}{"name": "testUser"}).Find(queryUser)
	db.Debug().Model(queryUser).Related(&queryUser.Company)
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}

func relatedInsertUserWithCustomReferenceKey(db *gorm.DB) {
	queryUser := new(userWithCustomReferenceKey)
	queryUser.Name = "testUser"
	queryUser.Company = company{
		Name: "testCompany",
		Code: 2001,
	}
	db.Debug().Create(queryUser)
	fmt.Printf("The Query Result Is %+v\n", queryUser)
}

func main() {
	db := common.InitDB("./belongTo/gorm.db")
	common.MigrateDB(db, &company{}, &user{}, &userWithCustomForeignKey{}, &userWithCustomReferenceKey{})
	preloadQueryUserWithCustomReferenceKey(db)
	common.SetDB(db)
	common.CloseDB(db)
}
