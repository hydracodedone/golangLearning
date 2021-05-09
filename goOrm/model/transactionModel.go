package model

import "gorm.io/gorm"

type TransactionModel struct {
	gorm.Model
	Info string `gorm:"unique"`
}

func (t *TransactionModel) TableName() string {
	return "transaction_model"
}
