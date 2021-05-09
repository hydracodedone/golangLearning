package hook

import (
	"database_demo/migrate"
	"database_demo/connection"
	"database_demo/model"
	"fmt"
	"log"
)

func hookDrop() {
	var data model.HookModel
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", data.TableName()))
}
func HookCreate() {
	var hookModel model.HookModel
	var createData model.HookModel = model.HookModel{
		Info: nil,
	}
	db := connection.GormDB.Debug().Model(&hookModel).Create(&createData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	fmt.Printf("%+v\n", *(createData.Info))
}
func HookSearch() {
	var searchData model.HookModel

	db := connection.GormDB.Debug().Table("hook_model").Take(&searchData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	fmt.Printf("%+v\n", *(searchData.Info))

	var searchDataStruct struct {
		Info *string
	}
	//struct will skip hook
	db = connection.GormDB.Debug().Table("hook_model").Take(&searchDataStruct)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	fmt.Printf("%+v\n", *(searchDataStruct.Info))
}
func HookUpdate() {
	//Changed它只是检查 Model 对象字段的值与 Update、Updates 的值是否相等，如果值有变更，且字段没有被忽略，则返回 true
	//
	var hookModel model.HookModel
	updateString := "this is update"
	var updateData = model.HookModel{
		Info: &updateString,
	}
	db := connection.GormDB.Debug().Model(&hookModel).Where("id=?", 1).Updates(&updateData)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	fmt.Printf("%+v\n", *(updateData.Info))

}
func HookDelete() {
	info := ""
	var hookModel model.HookModel = model.HookModel{
		Info: &info,
	}

	db := connection.GormDB.Debug().Unscoped().Delete(&hookModel, []int{1})
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}
	fmt.Printf("%+v\n", hookModel)
}
func HookDemo() {
	hookDrop()
	migrate.AutoMigrateModel([]interface{}{&model.HookModel{}})
	fmt.Println("create")
	HookCreate()
	fmt.Println("delete")
	HookUpdate()
}
