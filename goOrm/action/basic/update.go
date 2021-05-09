package basicOperation

import (
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Update1() {
	var basicModel model.Basic
	//使用save,建议使用Model
	var resultModel model.Basic
	db := connection.GormDB.Debug().Model(&basicModel).First(&resultModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", resultModel)
		resultModel.Embeded.Info = "this is update"
		//如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)。
		connection.GormDB.Debug().Save(&resultModel)
	}
}
func Update2() {
	var basicModel model.Basic
	//使用save,建议使用Model
	db := connection.GormDB.Debug().Model(&basicModel).Where("id=?", 1).Update("embed_info", "this is update")
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
	// 根据 `struct` 更新属性，只会更新非零值的字段
	updateInfo := model.Basic{
		Embeded: model.Embed{
			Info: "this is a test",
		},
	}
	db = connection.GormDB.Debug().Model(&basicModel).Where("id=?", 1).Updates(&updateInfo)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func Update3() {
	var basicModel model.Basic
	updateInfoMap := map[string]interface{}{"embed_info": "this is update", "only_create": "this is update"}
	db := connection.GormDB.Debug().Model(&basicModel).Select("embed_info").Where("id=?", 1).Updates(&updateInfoMap)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func Update4() {
	var basicModel model.Basic
	updateInfoMap := map[string]interface{}{"embed_info": "this is update", "only_create": "this is update", "only_update": nil}
	db := connection.GormDB.Debug().Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&basicModel).Updates(&updateInfoMap)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func Update5() {
	var basicModel model.Basic
	//从源码得知,UpdateColumn会跳过钩子函数,Update不会
	//UpdateColumn也会跳过时间追踪
	db := connection.GormDB.Debug().Model(&basicModel).Select("embed_info").Where("id=?", 1).UpdateColumn("embed_info", gorm.Expr("upper(embed_info)"))
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func Update6() {
	var basicModel model.Basic
	connection.GormDB.Session(&gorm.Session{DryRun: true}).Debug().Model(&basicModel).Select("embed_info").Where("id=?", 1).Update("embed_info", connection.GormDB.Table(basicModel.TableName()).Select("embed_info").Where("id=?", 2))
}

func Update7() {
	var updateResults []model.Basic
	updateInfoMap := map[string]interface{}{"embed_info": "this is update", "only_create": "this is update"}
	db := connection.GormDB.Debug().Model(&updateResults).Clauses(clause.Returning{}).Select("embed_info").Where("id=?", 1).Updates(&updateInfoMap)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
	//返回被更新的数据，仅当数据库支持回写功能时才能正常运行 MYSQL不支持
	fmt.Printf("%+v\n", updateResults)
}
func UpdateDemo() {
	dropTable()
	autoMigrateTable()
	Create3()
	Update7()
}
