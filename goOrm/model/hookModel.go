package model

import (
	"fmt"

	"gorm.io/gorm"
)

/*
hook流程

开始事务
BeforeSave
BeforeCreate BefroreUpdate BeforeDelete
关联前的
db
关联后的
AfterCreate AfterUpdate AfterDelete
AfterSave
提交或回滚事务
*/
type HookModel struct {
	gorm.Model
	Info *string
}

func (b *HookModel) TableName() string {
	return "hook_model"
}

func (b *HookModel) BeforeSave(gormDB *gorm.DB) (err error) {
	info := "before_save"
	b.Info = &info
	fmt.Printf("before save [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) BeforeCreate(gormDB *gorm.DB) (err error) {
	info := "before_create"
	b.Info = &info
	fmt.Printf("before create [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) AfterCreate(gormDB *gorm.DB) (err error) {
	info := "after_create"
	b.Info = &info
	fmt.Printf("after create [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) BeforeUpdate(gormDB *gorm.DB) (err error) {
	if gormDB.Statement.Changed() {
		fmt.Println("info column changed")
	}
	info := "before_update"
	b.Info = &info
	fmt.Printf("before update [%+v]\n", *(b.Info))
	return nil

}
func (b *HookModel) AfterUpdate(gormDB *gorm.DB) (err error) {
	info := "after_update"
	b.Info = &info
	fmt.Printf("after update [%+v]\n", *(b.Info))
	return nil
}

func (b *HookModel) BeforeDelete(gormDB *gorm.DB) (err error) {
	info := "before_delete"
	b.Info = &info
	fmt.Printf("before delete [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) AfterDelete(gormDB *gorm.DB) (err error) {
	info := "after_delete"
	b.Info = &info
	fmt.Printf("after delete [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) AfterFind(gormDB *gorm.DB) (err error) {
	info := "after_find"
	b.Info = &info
	fmt.Printf("after find [%+v]\n", *(b.Info))
	return nil
}
func (b *HookModel) AfterSave(gormDB *gorm.DB) (err error) {
	info := "after_save"
	b.Info = &info
	fmt.Printf("after save [%+v]\n", *(b.Info))
	return nil
}
