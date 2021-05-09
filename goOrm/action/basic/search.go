package basicOperation

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"database_demo/connection"
	"database_demo/model"
)

func Query1() {
	var basicModel model.Basic
	//不需要指定表或者模型
	db := connection.GormDB.Debug().Take(&basicModel, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}

	db = connection.GormDB.Debug().First(&basicModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}

	db = connection.GormDB.Debug().Last(&basicModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
}
func Query2() {
	var basicModel model.Basic
	data := map[string]interface{}{}
	// SELECT * FROM `base_model` WHERE `base_model`.`deleted_at` IS NULL ORDER BY `base_model`.`id` LIMIT 1
	db := connection.GormDB.Debug().Model(&basicModel).First(&data)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", data)
	}

	//SELECT * FROM `base_model` WHERE `base_model`.`deleted_at` IS NULL LIMIT 1
	db = connection.GormDB.Debug().Model(&basicModel).Take(&data)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", data)
	}

	//SELECT * FROM `base_model` LIMIT 1
	db = connection.GormDB.Debug().Table("base_model").Take(&data)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", data)
	}
}

func Query3() {
	//GORM 允许扫描结果至 map[string]interface{} 或 []map[string]interface{}，此时别忘了指定 Model 或 Table
	var basicModel model.Basic
	var sliceMapResult []map[string]interface{}
	basicModelMap := map[string]interface{}{
		"embed_info": "batch_create_3",
	}
	//查询结果对于空值不会填充
	//[map[create_milli:<nil> create_nano:<nil> create_time_stamp:<nil> created_at:<nil> deleted_at:<nil> embed_info:batch_create_3 id:3 only_create:<nil> only_read:<nil> only_update:<nil> update_time_stamp:<nil> updated_at:<nil> updated_milli:<nil> updated_nano:<nil>]]
	db := connection.GormDB.Debug().Model(&basicModel).Where(basicModelMap).Find(&sliceMapResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", sliceMapResult)
	}

	//[{Model:{ID:3 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Embeded:{Info:batch_create_3} Update:{UpdatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedNano:0 UpdatedMilli:0 UpdateTimeStamp:0} Create:{CreatedAt:0 CreateNano:0 CreateMilli:0 CreateTimeStamp:0} Delete:{DeletedAt:0001-01-01 00:00:00 +0000 UTC} Authority:{OnlyCreate: OnlyUpdate: OnlyRead:}}]
	var sliceResult []model.Basic
	db = connection.GormDB.Debug().Model(&basicModel).Where(basicModelMap).Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", sliceResult)
	}

	primaryKeySlice := []int{1, 2, 3}
	db = connection.GormDB.Debug().Model(&basicModel).Where(primaryKeySlice).Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", sliceResult)
	}
}

func Query4() {
	//优先级从高到低为NOT AND OR
	//OR在条件最前面是无效的
	var basicModel model.Basic
	//SELECT * FROM `base_model` WHERE (NOT id=100 OR `id` = 1 AND id>10) AND `base_model`.`deleted_at` IS NULL
	//等价于 (NOT id=100) OR (`id` = 1 AND id>10)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Not("id=?", 100).Or(map[string]interface{}{"id": 1}).Where("id>?", 10).Find(&basicModel)

	//SELECT * FROM `base_model` WHERE (NOT id=100 AND id>10 OR `id` = 1) AND `base_model`.`deleted_at` IS NULL
	//等价于 (NOT id=100 AND id>10) OR (`id` = 1)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Not("id=?", 100).Where("id>?", 10).Or(map[string]interface{}{"id": 1}).Find(&basicModel)

	//SELECT * FROM `base_model` WHERE (`id` = 1 AND NOT id=100 AND id>10) AND `base_model`.`deleted_at` IS NULL
	//等价于(`id` = 1) AND (NOT id=100) AND (id>10)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Or(map[string]interface{}{"id": 1}).Not("id=?", 100).Where("id>?", 10).Find(&basicModel)

	//SELECT * FROM `base_model` WHERE (`id` = 1 AND id>10 AND NOT id=100) AND `base_model`.`deleted_at` IS NULL
	//等价于 (`id` = 1) AND (id>10) AND (NOT id=100)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Or(map[string]interface{}{"id": 1}).Where("id>?", 10).Not("id=?", 100).Find(&basicModel)

	//SELECT * FROM `base_model` WHERE (id>10 OR `id` = 1 AND NOT id=100) AND `base_model`.`deleted_at` IS NULL
	//(id>10) OR (`id` = 1 AND (NOT id=100))
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Where("id>?", 10).Or(map[string]interface{}{"id": 1}).Not("id=?", 100).Find(&basicModel)

	//SELECT * FROM `base_model` WHERE (id>10 AND NOT id=100 OR `id` = 1) AND `base_model`.`deleted_at` IS NULL
	//(id>10 AND (NOT id=100)) OR (`id` = 1)
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Where("id>?", 10).Not("id=?", 100).Or(map[string]interface{}{"id": 1}).Find(&basicModel)

	// SELECT * FROM `base_model` WHERE (id=1 AND (id=(1,2,3) OR id=(4,5,6))) AND `base_model`.`deleted_at` IS NULL
	//( id=1 AND ( id=(1,2,3) OR id=(4,5,6) ) )
	//推荐使用该写法
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Where(connection.GormDB.Where("id=?", 1).Where(connection.GormDB.Where("id=?", []int{1, 2, 3}).Or("id=?", []int{4, 5, 6}))).Find(&basicModel)
	//SELECT * FROM `base_model` WHERE id=1 AND `base_model`.`deleted_at` IS NULL
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Where("id=@id", map[string]interface{}{"id": 1}).Find(&basicModel)
	//SELECT * FROM `base_model` WHERE id=1 AND `base_model`.`deleted_at` IS NULL
	connection.GormDB.Debug().Session(&gorm.Session{DryRun: true}).Model(&basicModel).Where("id=@id", sql.Named("id", 1)).Find(&basicModel)
}

func Query5() {
	var basicModel model.Basic
	var sliceResult []model.Basic
	//用了Distinct再用Select就无效了
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Distinct("embed_info").Select("only_create").Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", sliceResult)
	}
}

func Query6() {
	var basicModel model.Basic
	var sliceResult []struct {
		Count     int
		EmbedInfo string
	}
	//自定义字段通过Find映射
	//Find查询不到不会报错
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Select("count(*) as count, embed_info").Where("id<?", 20).Group("embed_info").Having("count>?", 0).Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", sliceResult)
	}
}
func Query7() {
	var basicModel model.Basic
	db := connection.GormDB.Debug().Model(&basicModel).Where("id=?", 10).Scan(&basicModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", basicModel)
	}

	var structResult struct {
		ID          string
		CreatedAt   time.Time
		OnlyCreate  string
		UpdatedNano int64
	}
	//Scan,Find查询不到数据不会报错
	db = connection.GormDB.Debug().Model(&basicModel).Where("id=?", 10).Scan(&structResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", structResult)
	}
}
func Query8() {
	var basicModel model.Basic
	var sliceResult []model.Basic
	//构造Where子查询
	//SELECT * FROM `base_model` WHERE id in(SELECT `id` FROM `base_model` WHERE `base_model`.`deleted_at` IS NULL) AND `base_model`.`deleted_at` IS NULL
	db := connection.GormDB.Debug().Model(&basicModel).Where("id in(?)", connection.GormDB.Model(&basicModel).Select("id")).Find(&sliceResult)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", sliceResult)
	}
	//构造from子查询
	var resultStruct []struct {
		ID string
	}
	db = connection.GormDB.Debug().Table("(?) as u", connection.GormDB.Model(&basicModel).Select("id,embed_info")).Select("id").Where("ID IS NOT NULL").Find(&resultStruct)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", resultStruct)
	}
	var anotherResultStruct []struct {
		AVGID float64
	}
	//having后面的子查寻记得加括号
	db = connection.GormDB.Debug().Model(&basicModel).Select("AVG(id) as avg_id").Group("id").Having("avg_id<(?)", connection.GormDB.Model(&basicModel).Select("SUM(id) as sum_id").Where("id>?", 1)).Find(&anotherResultStruct)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", anotherResultStruct)
	}
}
func Query9() {
	var basicModel model.Basic
	var resultModel model.Basic
	//找不到,则初始化
	db := connection.GormDB.Debug().Model(&basicModel).Where("id =?", 1).Where("embed_info=?", "batch_create_2").Attrs("OnlyCreate", "OnlyCreate").FirstOrInit(&resultModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", resultModel)
	} else {
		fmt.Printf("%+v\n", resultModel)
	}
	//不管是否找到记录Assign 都会将属性赋值给 struct，但这些属性不会被用于生成查询 SQL，也不会被保存到数据库
	db = connection.GormDB.Debug().Model(&basicModel).Where("id =?", 1).Where("embed_info=?", "batch_create_1").Assign("OnlyCreate", "OnlyCreate1").FirstOrInit(&resultModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", resultModel)
	} else {
		fmt.Printf("%+v\n", resultModel)
	}
	//找到会不会更新,找到会更新
	db = connection.GormDB.Debug().Model(&basicModel).Where("id =?", 4).Where("embed_info=?", "batch_create_4").Attrs(map[string]interface{}{"id": 4, "OnlyCreate": "OnlyCreate4", "embed_info": "batch_create_4"}).FirstOrCreate(&resultModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", resultModel)
	} else {
		fmt.Printf("%+v\n", resultModel)
	}
	//都会更新
	db = connection.GormDB.Debug().Model(&basicModel).Where("id =?", 4).Where("embed_info=?", "batch_create_4").Assign(map[string]interface{}{"id": 4, "OnlyCreate": "OnlyCreate4", "embed_info": "batch_create_4"}).FirstOrCreate(&resultModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", resultModel)
	} else {
		fmt.Printf("%+v\n", resultModel)
	}
}
func Query10() {
	var basicModel model.Basic
	var Id interface{}
	var EmbedInfo interface{}
	var OnlyCreate interface{}

	row := connection.GormDB.Debug().Model(&basicModel).Select("id", "embed_info", "only_create").Where("id=?", 3).Row()
	err := row.Err()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		//If more than one row matches the query, Scan uses the first row and discards the rest
		err := row.Scan(&Id, &EmbedInfo, &OnlyCreate)
		if err != nil {
			log.Fatal(err)
		} else {
			if stringId, ok := Id.(int64); ok {
				fmt.Printf("Id:[%d] \n", stringId)
			} else {
				fmt.Printf("Id:[%v] \n", stringId)
			}
			if stringEmbedInfo, ok := EmbedInfo.([]byte); ok {
				fmt.Printf("EmbedInfo:[%s] \n", stringEmbedInfo)
			} else {
				fmt.Printf("EmbedInfo:[%v] \n", EmbedInfo)
			}
			if stringOnlyCreate, ok := OnlyCreate.([]byte); ok {
				fmt.Printf("OnlyCreate:[%s] \n", stringOnlyCreate)
			} else {
				fmt.Printf("OnlyCreate:[%v] \n", OnlyCreate)
			}
		}
	}
}
func Query11() {
	var basicModel model.Basic
	var resultModel model.Basic

	var structResult struct {
		ID         string
		EmbedInfo  string
		OnlyCreate string
	}

	rows, err := connection.GormDB.Debug().Model(&basicModel).Rows()
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			err := connection.GormDB.ScanRows(rows, &structResult)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%+v\n", structResult)
			}
		}
	}
	rows, err = connection.GormDB.Debug().Model(&basicModel).Rows()
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			err := connection.GormDB.ScanRows(rows, &resultModel)
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("%+v\n", resultModel)
			}
		}
	}
}
func Query12() {
	var basicModel model.Basic
	var resultModel []model.Basic
	handleBatchFunc := func(tx *gorm.DB, batch int) error {
		fmt.Printf("BATCH RowAffected %v\n", tx.RowsAffected)
		fmt.Printf("BATCH ERROR %v\n", tx.Error)
		for _, value := range resultModel {
			fmt.Printf("%+v\n", value)
		}
		return nil
	}
	db := connection.GormDB.Debug().Model(&basicModel).Where("id<?", 10).FindInBatches(&resultModel, 2, handleBatchFunc)
	if db.Error != nil {
		log.Fatal(db.Error)
	}
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
}
func Query13() {
	// 超过一列的查询，应该使用 `Scan` 或者 `Find`
	//查询单列,可以使用Pluck
	var embedInfo []string
	var basicModel model.Basic
	db := connection.GormDB.Debug().Model(&basicModel).Distinct().Pluck("embed_info", &embedInfo)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else if db.RowsAffected != 0 {
		fmt.Printf("%+v\n", embedInfo)
	} else {
		fmt.Printf("%+v\n", embedInfo)
	}
}
func Query14() {
	var basicModel model.Basic
	var count int64
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Where("id=?", 5).Count(&count)
	//RowsAffected总为1
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", count)
	}
}
func SkipID(db *gorm.DB) *gorm.DB {
	return db.Where("id <> ?", 1)
}
func SkipIDS(idList []int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id not in ?", idList)
	}
}
func Query15() {
	var basicModel model.Basic
	//不需要指定表或者模型
	db := connection.GormDB.Debug().Scopes(SkipID, SkipIDS([]int{1, 2})).First(&basicModel)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
}
func Query16() {
	var basicModel model.Basic
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Where("null_string is NULL").Where("normal_bool =?", 0).First(&basicModel)
	//RowsAffected总为1
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
	db = connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Where("null_string is NULL").Where("normal_bool =?", false).First(&basicModel)
	//RowsAffected总为1
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", basicModel)
	}
}

func Query17() {
	var basicModel model.Basic
	var queryResult []model.Basic
	db := connection.GormDB.Debug().Session(&gorm.Session{SkipHooks: true}).Model(&basicModel).Where("(id,embed_info) in ?", [][]interface{}{{1, "batch_create_1"}, {2, "batch_create_2"}, {3, "batch_create_3"}}).Scan(&queryResult)
	//RowsAffected总为1
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", queryResult)
	}

}
func SearchDemo() {
	dropTable()
	autoMigrateTable()
	Create1()
	Query1()
}
