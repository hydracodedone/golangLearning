package migrate

import (
	"database_demo/connection"
	"database_demo/model"
	"log"
)

func AutoMigrateModel(modelList []interface{}) {
	if connection.GormDB == nil {
		return
	}
	if modelList == nil {
		modelList = []interface{}{
			&model.Basic{},
		}
	}
	err := connection.GormDB.AutoMigrate(modelList...)
	if err != nil {
		log.Fatal(err)
	}
}
