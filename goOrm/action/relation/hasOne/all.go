package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func autoMigrateReferenceForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceForeignKeyHasOneUser{}, &model.ReferenceForeignKeyHasOneCard{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyHasOneCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyHasOneUser{}).TableName()))
}

func HasOneCreate1() {
	user := model.ReferenceForeignKeyHasOneCard{
		ID:       1,
		CardName: "card1",
		UserName: nil,
	}
	db := connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func HasOneCreate2() {
	card := model.ReferenceForeignKeyHasOneCard{
		ID:       1,
		CardName: "card1",
		UserName: nil,
	}
	user := model.ReferenceForeignKeyHasOneUser{
		ID:       1,
		UserName: "user1",
		Card:     card,
	}
	db := connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func HasOneCreate3() {
	user := model.ReferenceForeignKeyHasOneUser{
		ID:       1,
		UserName: "user1",
	}
	card1 := model.ReferenceForeignKeyHasOneCard{
		ID:       1,
		CardName: "card1",
		UserName: &user.UserName,
	}
	card2 := model.ReferenceForeignKeyHasOneCard{
		ID:       2,
		CardName: "card2",
		UserName: nil,
	}
	db := connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}

	db = connection.GormDB.Debug().Create(&card1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", card1)
	}

	db = connection.GormDB.Debug().Create(&card2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", card2)
	}
}

func HasOneQuery1() {
	user := model.ReferenceForeignKeyHasOneUser{}
	db := connection.GormDB.Model(&user).Debug().Find(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasOneQuery2() {
	user := model.ReferenceForeignKeyHasOneUser{}
	db := connection.GormDB.Model(&user).Debug().Preload("Card").Find(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasOneQuery3() {
	user := model.ReferenceForeignKeyHasOneUser{}
	db := connection.GormDB.Model(&user).Debug().Joins("Card").Find(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasOneQuery4() {
	var cards []model.ReferenceForeignKeyHasOneCard
	db := connection.GormDB.Model(&model.ReferenceForeignKeyHasOneCard{}).Debug().Joins("Join `reference_foreign_key_has_one_user` on `reference_foreign_key_has_one_card`.`user_name`=`reference_foreign_key_has_one_user`.`user_name`").Find(&cards)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", cards)
	}
}
func HasOneUpdate1() {
	user := model.ReferenceForeignKeyHasOneUser{}
	db := connection.GormDB.Model(&user).Debug().Joins("Card").Find(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	user.Card.CardInfo = "card info"
	user.Card.CardName = "new card"
	connection.GormDB.Save(&user.Card)
}

func HasOneUpdate2() {
	user := model.ReferenceForeignKeyHasOneUser{}
	db := connection.GormDB.Model(&user).Debug().Joins("Card").Find(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	user.UserName = "new user"
	connection.GormDB.Save(&user)
}
func HasOneReferenceForeignKeyDemo() {
	dropReferenceForeignKeyTable()
	autoMigrateReferenceForeignKeyTable()
	HasOneCreate3()
	HasOneUpdate2()
}
