package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateErrorReferenceForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ErrorReferenceForeignKeyBelongToUser{}, &model.ErrorReferenceForeignKeyBelongToCompany{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropErrorReferenceForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ErrorReferenceForeignKeyBelongToCompany{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ErrorReferenceForeignKeyBelongToUser{}).TableName()))

}

func BelongToErrorReferenceForeignKeyDemo() {
	dropErrorReferenceForeignKeyTable()
	autoMigrateErrorReferenceForeignKeyTable()
}
