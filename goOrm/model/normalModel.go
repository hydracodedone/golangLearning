package model

import "gorm.io/gorm"

type NormalModel struct {
	gorm.Model
	Info string `gorm:"unique"`
}

func (n *NormalModel) TableName() string {
	return "normal_model"
}
