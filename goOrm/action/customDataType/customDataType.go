package customDataType

import (
	"database_demo/migrate"
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func customDrop() {
	var data model.CustomDataTable
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", data.TableName()))
}

func customCreate() {
	var data model.CustomDataTable = model.CustomDataTable{
		CustomInfo: model.CustomInfo{
			Name: "Hydra",
			Age:  "23",
		},
	}
	connection.GormDB.Debug().Create(&data)
}

func customSearch() {
	var data model.CustomDataTable
	connection.GormDB.Debug().Model(&data).First(&data)
	fmt.Printf("%+v\n", data)
}
func customUpdate() {
	var data model.CustomDataTable
	customUpdateData := model.CustomInfo{
		Name: "Hydra2",
		Age:  "24",
	}
	db := connection.GormDB.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&data).Update("custom_info", &customUpdateData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func CustmDataTypeDemo() {
	customDrop()
	migrate.AutoMigrateModel([]interface{}{&model.CustomDataTable{}})
	customCreate()
	customSearch()
	customUpdate()
}
