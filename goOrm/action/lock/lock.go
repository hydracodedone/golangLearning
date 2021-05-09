package lock

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	basicOperation "database_demo/action/basic"
	"database_demo/connection"
	"database_demo/model"
)

func autoMigrateTable() {
	connection.GormDB.AutoMigrate(&model.Basic{})
}
func dropTable() {
	basicInstance := model.Basic{}
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", basicInstance.TableName()))
}
func LockQuery1() {
	var basicModel model.Basic
	//FOR SHARE和LOCK IN SHARE MODE是等价的
	db := connection.GormDB.Debug().Clauses(clause.Locking{Strength: "SHARE", Options: "NOWAIT"}).Find(&basicModel, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
}

func LockQuery2() {
	var basicModel model.Basic
	db := connection.GormDB.Debug().Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).Find(&basicModel, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
}
func LockDemo() {
	dropTable()
	autoMigrateTable()
	basicOperation.Create1()
	LockQuery2()
}
