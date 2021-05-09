package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateNormalTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.NormalMany2ManyLanguage{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.NormalMany2ManyUser{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropNormalTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", "nomal_many_2_many_mid_table"))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalMany2ManyLanguage{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalMany2ManyUser{}).TableName()))

}

func Many2ManyNormalDemo() {
	dropNormalTable()
	autoMigrateNormalTable()
}
