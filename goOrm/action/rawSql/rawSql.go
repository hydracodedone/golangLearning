package rawSql

import (
	"database/sql"
	"database_demo/connection"
	"database_demo/migrate"
	"database_demo/model"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RawSQLDrop() {
	var data model.NormalModel
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", data.TableName()))
}
func RawSQLInsert() {
	var normalModel model.NormalModel
	var createData = []model.NormalModel{
		{Info: "this is a test 1"},
		{Info: "this is a test 2"},
	}
	db := connection.GormDB.Debug().Model(&normalModel).CreateInBatches(&createData, 2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error)
	}
}

func RawSQL1() {
	var results []model.NormalModel
	connection.GormDB.Raw("SELECT * FROM normal_model").Scan(&results)
	fmt.Printf("result %#v\n", results)
	internalSQL := clause.Expr{
		SQL:  "CONCAT_WS(?,?,?)",
		Vars: []interface{}{"-", "basic", "info"},
	}
	db := connection.GormDB.Debug().Exec("update normal_model set info=@info where id=@id", sql.NamedArg{Name: "info", Value: internalSQL}, sql.Named("id", 1))
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error)
	}
}
func RawSQL2() {
	internalSQL := clause.Expr{
		SQL:  "CONCAT_WS(?,?,?)",
		Vars: []interface{}{"-", "basic", "info"},
	}
	sqlString := connection.GormDB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Exec("update normal_model set info=@info where id=@id", sql.NamedArg{Name: "info", Value: internalSQL}, sql.Named("id", 1))
	})
	fmt.Println(sqlString)
}

func RawSQL3() {
	var info string
	var id int64
	sqlRow := connection.GormDB.Raw("SELECT id, info FROM normal_model").Row()
	fmt.Printf("result %#v\n", sqlRow)
	for {
		err := sqlRow.Scan(&id, &info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("id:%d,info:%s\n", id, info)
	}
}

func RawSQL4() {
	var info string
	var id int64
	sqlRows, err := connection.GormDB.Raw("SELECT id, info FROM normal_model").Rows()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("result %#v\n", sqlRows)
	defer sqlRows.Close()
	for sqlRows.Next() {
		err := sqlRows.Scan(&id, &info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("id:%d,info:%s\n", id, info)
	}
}

func RawSQL5() {
	var info struct {
		Info string
		ID   int64
	}
	sqlRows, err := connection.GormDB.Raw("SELECT id, info FROM normal_model").Rows()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("result %#v\n", sqlRows)
	defer sqlRows.Close()
	for sqlRows.Next() {
		err := connection.GormDB.ScanRows(sqlRows, &info)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%+v\n", info)
	}
}
func RawSQLDemo() {
	RawSQLDrop()
	migrate.AutoMigrateModel([]interface{}{&model.NormalModel{}})
	RawSQLInsert()
	RawSQL5()
}
