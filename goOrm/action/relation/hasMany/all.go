package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func autoMigrateReferenceForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceForeignKeyHasManyUser{}, &model.ReferenceForeignKeyHasManyCard{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyHasManyCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyHasManyUser{}).TableName()))
}

func HasManyCreate1() {
	user := model.ReferenceForeignKeyHasManyCard{
		ID:        1,
		CardName:  "card1",
		UserNames: nil,
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

func HasManyCreate2() {
	card1 := model.ReferenceForeignKeyHasManyCard{
		ID:        1,
		CardName:  "card1",
		UserNames: nil,
	}
	card2 := model.ReferenceForeignKeyHasManyCard{
		ID:        2,
		CardName:  "card2",
		UserNames: nil,
	}
	user := model.ReferenceForeignKeyHasManyUser{
		ID:       1,
		UserName: "user1",
		Cards:    []model.ReferenceForeignKeyHasManyCard{card1, card2},
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

func HasManyCreate3() {
	card1 := model.ReferenceForeignKeyHasManyCard{
		ID:        1,
		CardName:  "card1",
		UserNames: nil,
	}
	card2 := model.ReferenceForeignKeyHasManyCard{
		ID:        2,
		CardName:  "card2",
		UserNames: nil,
	}
	card3 := model.ReferenceForeignKeyHasManyCard{
		ID:        3,
		CardName:  "card3",
		UserNames: nil,
	}
	user := model.ReferenceForeignKeyHasManyUser{
		ID:       1,
		UserName: "user1",
		Cards:    nil,
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
	err := connection.GormDB.Debug().Model(&user).Association("Cards").Append([]model.ReferenceForeignKeyHasManyCard{card1, card2})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
	db = connection.GormDB.Debug().Create(&card3)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}
func HasManyCreate4() {
	card1 := model.ReferenceForeignKeyHasManyCard{
		ID:        1,
		CardName:  "card1",
		UserNames: nil,
	}
	card2 := model.ReferenceForeignKeyHasManyCard{
		ID:        2,
		CardName:  "card2",
		UserNames: nil,
	}
	user := model.ReferenceForeignKeyHasManyUser{
		ID:       1,
		UserName: "user1",
		Cards:    []model.ReferenceForeignKeyHasManyCard{card1, card2},
	}
	db := connection.GormDB.Debug().Omit("Cards").Create(&user)

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
func HasManyCreate5() {
	card1 := model.ReferenceForeignKeyHasManyCard{
		ID:        1,
		CardName:  "card1",
		UserNames: nil,
	}
	card2 := model.ReferenceForeignKeyHasManyCard{
		ID:        2,
		CardName:  "card2",
		UserNames: nil,
	}
	user := model.ReferenceForeignKeyHasManyUser{
		ID:       1,
		UserName: "user1",
		Cards:    []model.ReferenceForeignKeyHasManyCard{card1, card2},
	}
	db := connection.GormDB.Debug().Omit(clause.Associations).Create(&user)
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

func HasManyQuery1() {
	user := model.ReferenceForeignKeyHasManyUser{}
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
func HasManyQuery2() {
	user := model.ReferenceForeignKeyHasManyUser{}
	db := connection.GormDB.Model(&user).Debug().Preload("Cards").Find(&user)
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
func HasManyQuery3() {
	user := model.ReferenceForeignKeyHasManyUser{}
	db := connection.GormDB.Debug().Model(&user).First(&user)
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
	err := connection.GormDB.Debug().Model(&user).Association("Cards").Find(&user.Cards)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasManyQuery4() {
	var userInfo []struct {
		UserId   int
		UserName string
		CardId   int
		CardName string
	}
	db := connection.GormDB.Debug().Raw("select user.id,user.user_name,card.id,card.card_name from reference_foreign_key_has_many_user as user join reference_foreign_key_has_many_card as card on user.user_name=card.user_name where user.id=1;").Scan(&userInfo)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", userInfo)
	}
}

func HasManyQuery5() {
	user := model.ReferenceForeignKeyHasManyUser{UserName: "user1"}
	//带上外键对应的主表的信息,就能查询出
	err := connection.GormDB.Debug().Model(&user).Where("card_name=?", "card1").Association("Cards").Find(&user.Cards)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasManyQuery6() {
	user := model.ReferenceForeignKeyHasManyUser{UserName: "user1"}
	//带上外键对应的主表的信息,就能查询出
	result := connection.GormDB.Debug().Model(&user).Where("card_name=?", "card1").Association("Cards").Count()
	fmt.Printf("%+v\n", result)
}

func HasManyQuery7() {
	var card model.ReferenceForeignKeyHasManyCard
	db := connection.GormDB.Debug().Model(&model.ReferenceForeignKeyHasManyCard{}).Preload("User").Find(&card, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", card)
	}
}
func HasManyUpdate1() {
	card3 := model.ReferenceForeignKeyHasManyCard{ID: 3}
	user := model.ReferenceForeignKeyHasManyUser{UserName: "user1"}
	//全部替换
	err := connection.GormDB.Debug().Model(&user).Association("Cards").Replace(&card3)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func HasManyDelete1() {
	user := model.ReferenceForeignKeyHasManyUser{UserName: "user1"}
	deleteCard := []model.ReferenceForeignKeyHasManyCard{{ID: 1}, {ID: 2}}
	err := connection.GormDB.Debug().Model(&user).Association("Cards").Delete(&deleteCard)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func HasManyDelete2() {
	user := model.ReferenceForeignKeyHasManyUser{UserName: "user1"}
	err := connection.GormDB.Debug().Model(&user).Association("Cards").Clear()
	if err != nil {
		log.Fatalln(err)
	}
}
func HasManyDelete3() {
	user := model.ReferenceForeignKeyHasManyUser{}
	db := connection.GormDB.Debug().Model(&user).Where("user_name=?", "user1").Delete(&user)
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
func HasManyDelete4() {
	user := model.ReferenceForeignKeyHasManyUser{}
	// db := connection.GormDB.Debug().Model(&user).Select("Cards").Where("user_name=?", "user1").Delete(&user)

	db := connection.GormDB.Debug().Model(&user).Select(clause.Associations).Where("user_name=?", "user1").Delete(&user)
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
func HasManyReferenceForeignKeyDemo() {
	dropReferenceForeignKeyTable()
	autoMigrateReferenceForeignKeyTable()
	HasManyCreate3()
	HasManyQuery7()
}
