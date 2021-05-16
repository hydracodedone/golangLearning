package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Company struct {
	ID   int `gorm:"primary_key,auto_increment"`
	Name string
}

type User2 struct {
	gorm.Model
	Name         string
	CompanyRefer int
	Company2     Company2 `gorm:"foreignKey:CompanyRefer"`
}

type Company2 struct {
	ID   int `gorm:"primary_key,auto_increment"`
	Name string
}

type User3 struct {
	gorm.Model
	Name       string
	Company3ID string
	Company3   Company3 `gorm:"references:Code"` // 使用 Code 作为引用
}

type Company3 struct {
	ID   int `gorm:"primary_key,auto_increment"`
	Name string
	Code string `gorm:"unique"`
}

type User4 struct {
	gorm.Model
	Name       string
	Company4ID int
	Company4   Company4 `gorm:"association_foreignkey:Refer"` // use Refer 作为关联外键
}

type Company4 struct {
	ID    int `gorm:"primary_key,auto_increment"`
	Name  string
	Code  string `gorm:"unique"`
	Refer int
}
