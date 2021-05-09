package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"hydracode.com/ormDemo/common"
)

type User struct {
	ID    int
	UName string
}
type CreditCard struct {
	ID        int
	CName     string
	UserRefer int
	User      User `gorm:"foreignKey:UserRefer"`
}

func insertData(db *gorm.DB) {
	creditCard := new(CreditCard)
	creditCard.CName = "card1"
	creditCard.User = User{
		UName: "u1",
	}
	db.Debug().Create(&creditCard)
}

func queryDataPreload(db *gorm.DB) {
	queryData := new(CreditCard)
	db.Debug().Model(queryData).Preload("User").First(queryData)
	fmt.Printf("The Query Data Is %+v\n", queryData)
}
func queryDataRelated(db *gorm.DB) {
	queryData := new(CreditCard)
	db.Debug().Model(queryData).First(queryData)
	fmt.Printf("The Query Data Is %+v\n", queryData)
	db.Debug().Model(queryData).Related(&queryData.User, "UserRefer")
	fmt.Printf("The Query Data Is %+v\n", queryData)
}

func queryDataReference(db *gorm.DB) {
	queryData := new(CreditCard)
	db.Debug().Model(queryData).First(queryData)
	fmt.Printf("The Query Data Is %+v\n", queryData)
	db.Debug().Model(queryData).Association("User").Find(&queryData.User)
	fmt.Printf("The Query Data Is %+v\n", queryData)
}

func main() {
	db := common.InitDB("./belongToForeignKey/gorm.db")
	common.MigrateDB(db, &CreditCard{}, &User{})
	common.SetDB(db)
	queryDataReference(db)
	common.CloseDB(db)
}
