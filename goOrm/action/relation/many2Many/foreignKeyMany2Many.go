package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ForeignKeyMany2ManyLanguage{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.ForeignKeyMany2ManyUser{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", "foreign_key_many_2_many_mid_table"))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyMany2ManyLanguage{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ForeignKeyMany2ManyUser{}).TableName()))
}

func Many2ManyForeignKeyDemo() {
	dropForeignKeyTable()
	autoMigrateForeignKeyTable()
}
