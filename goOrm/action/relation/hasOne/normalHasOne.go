package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateNormalTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.NormalHasOneUser{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.NormalHasOneCard{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropNormalTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalHasOneCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalHasOneUser{}).TableName()))
}

func HasOneNormalDemo() {
	dropNormalTable()
	autoMigrateNormalTable()
}
