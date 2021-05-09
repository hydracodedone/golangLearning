package basicOperation

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"database_demo/connection"
	"database_demo/model"
)

func Delete1() {
	/*
		如果在没有任何条件的情况下执行批量删除，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误
		1 增加条件
		2 使用原生SQL
		3 启用AllowGlobalUpdate
	*/
	var basicModel model.Basic
	//通过AllowGlobalUpdate避免全局删除
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: false, AllowGlobalUpdate: false}).Unscoped().Delete(&basicModel)
	//使用Unscoped()避免软删除
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: false, AllowGlobalUpdate: true}).Unscoped().Delete(&basicModel)
	//内联条件
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true, AllowGlobalUpdate: false}).Delete(&basicModel, "id=?", 1)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true, AllowGlobalUpdate: false}).Delete(&basicModel, []int{1, 2, 3})
	connection.GormDB.Debug().Where("id=?", 1).Delete(&model.Basic{})
	connection.GormDB.Debug().Exec("DELETE FROM base_model;")
}

func Delete2() {
	//返回被删除的数据，仅当数据库支持回写功能时才能正常运行 MYSQL不支持
	var basicModel model.Basic
	var deleteResults []model.Basic
	db := connection.GormDB.Debug().Session(&gorm.Session{DryRun: false, AllowGlobalUpdate: true}).Clauses(clause.Returning{}).Unscoped().Model(&basicModel).Delete(&deleteResults)
	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	} else {
		fmt.Printf("%+v\n", deleteResults)
	}
}

func DeleteDemo() {
	dropTable()
	autoMigrateTable()
	Create3()
	Delete2()
}
