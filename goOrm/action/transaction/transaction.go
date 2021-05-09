package transaction

import (
	"database_demo/connection"
	"database_demo/migrate"
	"database_demo/model"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func transanctionDrop() {
	var data model.TransactionModel
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", data.TableName()))
}
func Transaction1() {
	connection.GormDB.Debug().Session(&gorm.Session{SkipDefaultTransaction: true}).Transaction(func(tx *gorm.DB) error {
		var transationModel model.TransactionModel
		var createData = []model.TransactionModel{
			{Info: "this is a test"},
			{Info: "this is a test"},
		}
		db := tx.Debug().Model(&transationModel).CreateInBatches(&createData, 2)
		fmt.Printf("RowAffected %v\n", db.RowsAffected)
		if db.Error != nil {
			fmt.Println(db.Error)
			return db.Error
		} else {
			fmt.Printf("%+v\n", createData)
			return nil
		}
	})
}

func Transaction2() {
	connection.GormDB.Debug().Session(&gorm.Session{SkipDefaultTransaction: true}).Transaction(func(tx *gorm.DB) error {
		var transationModel model.TransactionModel
		var createData = []model.TransactionModel{
			{Info: "this is a test1"},
		}
		db := tx.Model(&transationModel).CreateInBatches(&createData, 2)
		fmt.Printf("RowAffected %v\n", db.RowsAffected)
		if db.Error != nil {
			fmt.Println(db.Error)
			return db.Error
		} else {
			fmt.Printf("%+v\n", createData)
		}
		//事务嵌套
		tx.Transaction(func(tx *gorm.DB) error {
			var createData = []model.TransactionModel{
				{Info: "this is a test1"},
			}
			db := tx.Model(&transationModel).CreateInBatches(&createData, 2)
			fmt.Printf("RowAffected %v\n", db.RowsAffected)
			if db.Error != nil {
				fmt.Println(db.Error)
				return db.Error
			} else {
				fmt.Printf("%+v\n", createData)
				return nil
			}
		})
		return nil
	})
}
func Transaction3() {
	//支持隔离级别
	tx := connection.GormDB.Debug().Session(&gorm.Session{SkipDefaultTransaction: true}).Begin()
	var transationModel model.TransactionModel
	var createData = model.TransactionModel{
		Info: "this is a test1",
	}
	db := tx.Model(&transationModel).Create(&createData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		fmt.Println(db.Error)
		fmt.Println("rollback to begin")
		tx.Rollback()
	} else {
		tx.SavePoint("create1")
	}
	db = tx.Model(&transationModel).Create(&createData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		fmt.Println(db.Error)
		fmt.Println("rollback to create1")
		tx.RollbackTo("create1")
	} else {
		tx.SavePoint("create2")
	}
	fmt.Println("commit")
	db = tx.Commit()
	if db.Error != nil {
		log.Fatalln(db.Error)
	}
}
func TransactionDemo() {
	transanctionDrop()
	migrate.AutoMigrateModel([]interface{}{&model.TransactionModel{}})
	Transaction3()
}
