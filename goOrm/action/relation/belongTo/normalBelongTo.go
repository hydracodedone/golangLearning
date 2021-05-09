package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateNormalTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.NormalBelongToCompany{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.NormalBelongToUser{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropNormalTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalBelongToCompany{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalBelongToUser{}).TableName()))
}

func BelongToNormalDemo() {
	dropNormalTable()
	autoMigrateNormalTable()
}
