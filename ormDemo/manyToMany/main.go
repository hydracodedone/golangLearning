package main

import (
	"hydracode.com/ormDemo/common"
)

type Teacher struct {
	ID      int
	TName   string
	Student []Student `gorm:"many2many:teacher_to_student"`
}
type Student struct {
	ID      int
	SName   string
	Teacher []Teacher `gorm:"many2many:teacher_to_student"`
}

func main() {
	db := common.InitDB("./manyToMany/gorm.db")
	common.MigrateDB(db, &Teacher{}, &Student{})
	common.SetDB(db)
	common.CloseDB(db)
}
