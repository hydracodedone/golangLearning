package relation

import (
	"database_demo/connection"
	"database_demo/model"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func autoMigrateReferenceForeignKeyTable() {
	err := connection.GormDB.Debug().AutoMigrate(&model.ReferenceForeignKeyBelongToUser{}, &model.ReferenceForeignKeyBelongToCompany{})
	if err != nil {
		log.Fatalln(err)
	}
	err = connection.GormDB.Debug().AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}
}
func dropReferenceForeignKeyTable() {
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyBelongToUser{}).TableName()))
	connection.GormDB.Debug().Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", (&model.ReferenceForeignKeyBelongToCompany{}).TableName()))
}

func BelongToCreate1() {
	user := model.ReferenceForeignKeyBelongToUser{
		ID:          1,
		UserName:    "user1",
		CompanyInfo: nil,
		Company:     model.ReferenceForeignKeyBelongToCompany{},
	}
	db := connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func BelongToCreate2() {
	company := model.ReferenceForeignKeyBelongToCompany{
		ID:          1,
		CompanyName: "company1",
	}
	company2 := model.ReferenceForeignKeyBelongToCompany{
		ID:          2,
		CompanyName: "company2",
	}
	user := model.ReferenceForeignKeyBelongToUser{
		ID:       1,
		UserName: "user1",
		//在指定从属关系的结构体,可以不用指定外键,可以自动填充
		CompanyInfo: nil,
		Company:     company,
	}
	db := connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	user2 := model.ReferenceForeignKeyBelongToUser{
		ID:       2,
		UserName: "user2",
		//在指定从属关系的结构体,可以不用指定外键,可以自动填充
		CompanyInfo: nil,
		Company:     company,
	}
	db = connection.GormDB.Debug().Create(&user2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	db = connection.GormDB.Debug().Create(&company2)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func BelongToCreate3() {
	company := model.ReferenceForeignKeyBelongToCompany{
		ID:          1,
		CompanyName: "company1",
	}
	db := connection.GormDB.Debug().Create(&company)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", company)
	}
	user := model.ReferenceForeignKeyBelongToUser{
		ID:       1,
		UserName: "user1",
		//在指定从属关系的结构体,可以不用指定外键,可以自动填充
		CompanyInfo: &company.CompanyName,
	}
	db = connection.GormDB.Debug().Create(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func BelongToQuery1() {
	user := model.ReferenceForeignKeyBelongToUser{}
	db := connection.GormDB.Model(&user).Debug().Find(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func BelongToQuery2() {
	user := model.ReferenceForeignKeyBelongToUser{}
	//preload 填的是结构体中的内嵌字段
	db := connection.GormDB.Model(&user).Debug().Preload("Company").First(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}
func BelongToQuery3() {
	users := []model.ReferenceForeignKeyBelongToUser{}
	//preload 填的是结构体中的内嵌字段
	//系统查询子表,然后将外键的集合作为查询字段查询主表,然后在将结果映射到内嵌的结构体中
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Preload("Company").Find(&users)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", users)
	}
}
func BelongToQuery4() {
	users := []model.ReferenceForeignKeyBelongToUser{}
	//joins 填的是结构体中的内嵌字段
	//进行的是联表查询
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Joins("Company").Find(&users)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", users)
	}
}
func BelongToQuery5() {
	companys := []model.ReferenceForeignKeyBelongToCompany{}
	//joins 填的是结构体中的内嵌字段
	//进行的是联表查询
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToCompany{}).Debug().Joins("Join `reference_foreign_key_belong_to_user` on `reference_foreign_key_belong_to_company`.`company_name`=`reference_foreign_key_belong_to_user`.`company_info`").Find(&companys)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", companys)
	}
}
func BelongToUpdate1() {
	user := model.ReferenceForeignKeyBelongToUser{}
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Joins("Company").First(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	user.UserName = "update user name"
	db = db.Session(&gorm.Session{AllowGlobalUpdate: true}).Save(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func BelongToUpdate2() {
	user := model.ReferenceForeignKeyBelongToUser{}
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Joins("Company").First(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	var company model.ReferenceForeignKeyBelongToCompany
	db = connection.GormDB.Model(&user.Company).Debug().Find(&company)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", company)
	}
	company.CompanyAddress = "update address"
	db.Save(&company)
}

func BelongToUpdate3() {
	user := model.ReferenceForeignKeyBelongToUser{}
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Joins("Company").First(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	user.Company.CompanyAddress = "update address"
	//通过指定约束constraint:OnUpdate:CASCADE来实现修改主表的主键时候实现外键联动
	user.Company.CompanyName = "new company"
	//不要使用db.save(),此时db已经被赋值了
	db = connection.GormDB.Debug().Save(&user.Company)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
}

func BelongToUpdate4() {
	user := model.ReferenceForeignKeyBelongToUser{}
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToUser{}).Debug().Joins("Company").First(&user, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	} else {
		fmt.Printf("%+v\n", user)
	}
	newCompanyName := "company2"
	user.CompanyInfo = &newCompanyName
	db = db.Save(&user)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}
func BelongToDelete1() {
	company := model.ReferenceForeignKeyBelongToCompany{}
	db := connection.GormDB.Model(&model.ReferenceForeignKeyBelongToCompany{}).Debug().Delete(&company, 1)
	fmt.Printf("RowAffected %v\n", db.RowsAffected)
	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			log.Println(db.Error.Error())
		} else {
			log.Fatalln(db.Error.Error())
		}
	}
}

func BelongToReferenceForeignKeyDemo() {
	dropReferenceForeignKeyTable()
	autoMigrateReferenceForeignKeyTable()
	BelongToCreate2()
	BelongToUpdate3()
}
