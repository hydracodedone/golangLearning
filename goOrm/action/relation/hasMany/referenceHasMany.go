package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func autoMigrateReferenceTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceHasManyUser{}, &model.ReferenceHasManyCard{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceHasManyCard{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceHasManyUser{}).TableName()))
}

func HasManyReferenceDemo() {
	dropReferenceTable()
	autoMigrateReferenceTable()
}
