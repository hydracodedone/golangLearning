package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateReferenceTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceHasOneUser{}, &model.ReferenceHasOneCard{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceHasOneCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceHasOneUser{}).TableName()))
}

func HasOneReferenceDemo() {
	dropReferenceTable()
	autoMigrateReferenceTable()
}
