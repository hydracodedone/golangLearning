package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateCustomTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.CustomMany2ManyLanguage{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.CustomMany2ManyUser{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.CustomMany2ManyUserLanguageMidTable{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropCustomTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.CustomMany2ManyUserLanguageMidTable{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.CustomMany2ManyUser{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.CustomMany2ManyLanguage{}).TableName()))
}

func Many2ManyCustomDemo() {
	dropCustomTable()
	autoMigrateCustomTable()
}
