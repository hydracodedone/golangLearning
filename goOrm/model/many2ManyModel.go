package model

import (
	"gorm.io/plugin/soft_delete"
)

type NormalMany2ManyUser struct {
	ID        int `gorm:"primaryKey"`
	UserName  string
	UserInfo  string
	Languages []NormalMany2ManyLanguage `gorm:"many2many:nomal_many_2_many_mid_table;"`
}

func (n *NormalMany2ManyUser) TableName() string {
	return "normal_many_to_many_user"
}

type NormalMany2ManyLanguage struct {
	ID           int `gorm:"primaryKey"`
	LanguageName string
	LanguageInfo string
}

func (n *NormalMany2ManyLanguage) TableName() string {
	return "normal_many_to_many_language"
}

type ReverseReferMany2ManyUser struct {
	ID        int `gorm:"primaryKey"`
	UserName  string
	UserInfo  string
	Languages []ReverseReferMany2ManyLanguage `gorm:"many2many:reverse_refer_many_2_many_mid_table;"`
}

func (n *ReverseReferMany2ManyUser) TableName() string {
	return "reverse_refer_many_to_many_user"
}

type ReverseReferMany2ManyLanguage struct {
	ID           int `gorm:"primaryKey"`
	LanguageName string
	LanguageInfo string
	Users        []ReverseReferMany2ManyUser `gorm:"many2many:reverse_refer_many_2_many_mid_table;"`
}

func (n *ReverseReferMany2ManyLanguage) TableName() string {
	return "reverse_refer_many_to_many_language"
}

type ForeignKeyMany2ManyUser struct {
	ID          int    `gorm:"primaryKey"`
	UserName    string `gorm:"unique"`
	UserInfo    string
	Languages   []ForeignKeyMany2ManyLanguage `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:UserName;joinForeignKey:UserName;references:LanguageName;joinReferences:LanguageName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DeletedFlag soft_delete.DeletedAt         `gorm:"softDelete:flag"`
}

func (n *ForeignKeyMany2ManyUser) TableName() string {
	return "foreign_key_many_to_many_user"
}

type ForeignKeyMany2ManyLanguage struct {
	ID           int    `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []ForeignKeyMany2ManyUser `gorm:"many2many:foreign_key_many_2_many_mid_table;foreignKey:LanguageName;joinForeignKey:LanguageName;references:UserName;joinReferences:UserName;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedFlag  soft_delete.DeletedAt     `gorm:"softDelete:flag"`
}

func (n *ForeignKeyMany2ManyLanguage) TableName() string {
	return "foreign_key_many_to_many_language"
}

type CustomMany2ManyUser struct {
	ID        int    `gorm:"primaryKey"`
	UserName  string `gorm:"unique"`
	UserInfo  string
	Languages []CustomMany2ManyLanguage `gorm:"many2many:custom_many_to_many_user_language_mid_table;foreignKey:UserName;joinForeignKey:UserName;references:LanguageName;joinReferences:LanguageName"`
}

func (n *CustomMany2ManyUser) TableName() string {
	return "custom_many_to_many_user"
}

type CustomMany2ManyLanguage struct {
	ID           int    `gorm:"primaryKey"`
	LanguageName string `gorm:"unique"`
	LanguageInfo string
	Users        []CustomMany2ManyUser `gorm:"many2many:custom_many_to_many_user_language_mid_table;foreignKey:LanguageName;joinForeignKey:LanguageName;references:UserName;joinReferences:UserName"`
}

func (n *CustomMany2ManyLanguage) TableName() string {
	return "custom_many_to_many_language"
}

type CustomMany2ManyUserLanguageMidTable struct {
	UserName     string `gorm:"primaryKey"`
	LanguageName string `gorm:"primaryKey"`
	OtherInfo    string
}

func (n *CustomMany2ManyUserLanguageMidTable) TableName() string {
	return "custom_many_to_many_user_language_mid_table"
}
