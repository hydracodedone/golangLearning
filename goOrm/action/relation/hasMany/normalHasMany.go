package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateNormalTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.NormalHasManyUser{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate(&model.NormalHasManyCard{})
	if err != nil {
		log.Fatalln(err)
	}
}
func dropNormalTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalHasManyCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.NormalHasManyUser{}).TableName()))
}

func HasManyNormalDemo() {
	dropNormalTable()
	autoMigrateNormalTable()
}
