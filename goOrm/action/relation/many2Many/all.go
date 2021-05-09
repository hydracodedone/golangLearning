package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func autoMigrateAllTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ForeignKeyMany2ManyLanguage{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.ForeignKeyMany2ManyUser{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropAllTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", "foreign_key_many_2_many_mid_table"))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyMany2ManyLanguage{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyMany2ManyUser{}).TableName()))
}

func Create1() {
	user1 := model.ForeignKeyMany2ManyUser{
		ID:        1,
		UserName:  "user1",
		UserInfo:  "this is user1",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	user2 := model.ForeignKeyMany2ManyUser{
		ID:        2,
		UserName:  "user2",
		UserInfo:  "this is user2",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	language1 := model.ForeignKeyMany2ManyLanguage{
		ID:           1,
		LanguageName: "language1",
		LanguageInfo: "this is language1",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}
	language2 := model.ForeignKeyMany2ManyLanguage{
		ID:           2,
		LanguageName: "language2",
		LanguageInfo: "this is language2",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}
	user1.Languages = append(user1.Languages, language1, language2)
	user2.Languages = append(user2.Languages, language1)
	db := connection.GormDB.Debug().Create(user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	db = connection.GormDB.Debug().Create(user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
}
func Create2() {
	user1 := model.ForeignKeyMany2ManyUser{
		ID:        1,
		UserName:  "user1",
		UserInfo:  "this is user1",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	user2 := model.ForeignKeyMany2ManyUser{
		ID:        2,
		UserName:  "user2",
		UserInfo:  "this is user2",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	language1 := model.ForeignKeyMany2ManyLanguage{
		ID:           1,
		LanguageName: "language1",
		LanguageInfo: "this is language1",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}
	language2 := model.ForeignKeyMany2ManyLanguage{
		ID:           2,
		LanguageName: "language2",
		LanguageInfo: "this is language2",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}

	db := connection.GormDB.Debug().Create(user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	db = connection.GormDB.Debug().Create(user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
	err := connection.GormDB.Debug().Model(&user1).Association("Languages").Append(&language1, &language2)
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().Model(&user2).Association("Languages").Append(&language1)
	if err != nil {
		log.Fatalln(err)
	}
}

func Create3() {
	//不建议使用save
	user1 := model.ForeignKeyMany2ManyUser{
		ID:        1,
		UserName:  "user1",
		UserInfo:  "this is user1",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	user2 := model.ForeignKeyMany2ManyUser{
		ID:        2,
		UserName:  "user2",
		UserInfo:  "this is user2",
		Languages: []model.ForeignKeyMany2ManyLanguage{},
	}
	language1 := model.ForeignKeyMany2ManyLanguage{
		ID:           1,
		LanguageName: "language1",
		LanguageInfo: "this is language1",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}
	language2 := model.ForeignKeyMany2ManyLanguage{
		ID:           2,
		LanguageName: "language2",
		LanguageInfo: "this is language2",
		Users:        []model.ForeignKeyMany2ManyUser{},
	}

	db := connection.GormDB.Debug().Create(user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	db = connection.GormDB.Debug().Create(user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
	user1.Languages = []model.ForeignKeyMany2ManyLanguage{language1, language2}
	user2.Languages = []model.ForeignKeyMany2ManyLanguage{language2}
	fmt.Println(111)
	db = connection.GormDB.Debug().Save(&user1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user1)
	}
	fmt.Println(222)

	db = connection.GormDB.Debug().Save(&user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user2)
	}
}
func Delete1() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Delete(&user, 1)
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
func Delete2() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Delete(&user, 1)
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
func Delete3() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Find(&user, 1)
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
	err := connection.GormDB.Model(&user).Association("Languages").Clear()
	if err != nil {
		log.Fatalln(err)
	}
}
func Delete4() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Find(&user, 1)
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
	err := connection.GormDB.Model(&user).Association("Languages").Clear()
	if err != nil {
		log.Fatalln(err)
	}
	db = connection.GormDB.Debug().Delete(&user, 1)
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
func Delete5() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Preload("Languages").Find(&user, 1)
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
	db = connection.GormDB.Debug().Delete(&user.Languages[0])
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	err := connection.GormDB.Debug().Model(&user.Languages[0]).Association("Users").Clear()
	if err != nil {
		log.Fatalln(err)
	}
}
func Query1() {
	var user model.ForeignKeyMany2ManyUser
	db := connection.GormDB.Debug().Model(&user).Preload("Languages").First(&user)
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
func Query2() {
	var language model.ForeignKeyMany2ManyLanguage
	db := connection.GormDB.Debug().Model(&language).Preload("Users", 1).First(&language)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", language)
	}
}
func Many2ManyAllDemo() {
	dropAllTable()
	autoMigrateAllTable()
	Create2()
	Query2()
}
