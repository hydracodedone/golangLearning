package model

type FirstTable struct {
	TID      uint   `gorm:"primary_key,auto_increment"`
	TestName string `gorm:"column:'test_info_name'"`
}

func (firstDemo *FirstTable) TableName() string {
	return "first_table_info"
}
