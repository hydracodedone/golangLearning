package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ForeignKeyHasOneUser{}, &model.ForeignKeyHasOneCard{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyHasOneCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyHasOneUser{}).TableName()))
}

func HasOneForeignKeyDemo() {
	dropForeignKeyTable()
	autoMigrateForeignKeyTable()
}
