package basicOperation

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"database_demo/connection"
	"database_demo/model"
)

func autoMigrateTable() {
	connection.GormDB.AutoMigrate(&model.Basic{})
}
func dropTable() {
	basicInstance := model.Basic{}
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", basicInstance.TableName()))
}

// 创建
func Create1() {
	basicInstance := model.Basic{
		Embeded: model.Embed{
			Info: "",
		},
		Authority: model.Authority{
			OnlyRead:   "OnlyRead",
			OnlyUpdate: "OnlyUpdate",
			OnlyCreate: "OnlyCreate",
		},
		NullAndNotNUllData: model.NullAndNotNUllData{
			NullBool: sql.NullBool{
				Bool:  true,
				Valid: false,
			},
			NormalBool: false,
			//在指定了某个字段的默认值后,如果想将该字段设置为零值,传入了零值,会被默认值覆盖掉,可以使用指针类型或者Sql.Null类型实现
			NullString: sql.NullString{
				String: "",
				Valid:  true,
			},
			NormalString:        "",
			DefaultNormalString: "",
		},
	}
	result := connection.GormDB.Debug().Create(&basicInstance)
	log.Printf("result error %v,result affected row %v\n", result.Error, result.RowsAffected)
}

// 指定字段创建
func Create2() {
	basicInstance := model.Basic{
		Embeded: model.Embed{
			Info: "hello world",
		},
	}
	db := connection.GormDB.Session(&gorm.Session{SkipHooks: true}).Select("embed_info").Create(&basicInstance)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
}

// 批量创建
func Create3() {
	createList := make([]map[string]interface{}, 3)
	createList[0] = map[string]interface{}{"embed_info": "batch_create_1", "deleted_flag": 0}
	createList[1] = map[string]interface{}{"embed_info": "batch_create_2", "deleted_flag": 0}
	createList[2] = map[string]interface{}{"embed_info": "batch_create_3", "deleted_flag": 0}
	db := connection.GormDB.Model(&model.Basic{}).Create(createList)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
}

func Create4() {
	createList := make([]map[string]interface{}, 4)
	createList[0] = map[string]interface{}{"embed_info": "batch_create_1"}
	createList[1] = map[string]interface{}{"embed_info": "batch_create_2"}
	createList[2] = map[string]interface{}{"embed_info": "batch_create_3"}
	createList[3] = map[string]interface{}{"embed_info": "batch_create_4"}
	db := connection.GormDB.Debug().Model(&model.Basic{}).CreateInBatches(createList, 2)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
}

// 创建时候指定SQL表达式
func Create5() {
	var param []interface{} = make([]interface{}, 3)
	param[0] = "-"
	param[1] = "demo"
	param[2] = "1"
	db := connection.GormDB.Debug().Model(&model.Basic{}).Create(
		map[string]interface{}{
			"embed_info": clause.Expr{
				SQL:  "CONCAT_WS(?,?,?)",
				Vars: param,
			}},
	)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
}
func CreateDemo() {
	dropTable()
	autoMigrateTable()
	Create1()
}
