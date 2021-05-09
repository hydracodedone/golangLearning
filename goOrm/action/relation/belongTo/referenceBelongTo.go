package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateReferenceTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceBelongToCompany{}, &model.ReferenceBelongToUser{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceBelongToUser{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceBelongToCompany{}).TableName()))
}

func BelongToReferenceDemo() {
	dropReferenceTable()
	autoMigrateReferenceTable()
}
