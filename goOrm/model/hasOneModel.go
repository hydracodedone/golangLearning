package model

type NormalHasOneUser struct {
	ID       int    `gorm:"primaryKey"`
	UserName string `gorm:"type:varchar(30);"`
	Card     NormalHasOneCard
}

type NormalHasOneCard struct {
	ID                 int    `gorm:"primaryKey"`
	CardName           string `gorm:"type:varchar(30)"`
	NormalHasOneUserID string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *NormalHasOneUser) TableName() string {
	return "normal_has_one_user"
}
func (n *NormalHasOneCard) TableName() string {
	return "normal_has_one_card"
}

type ForeignKeyHasOneUser struct {
	ID       int                  `gorm:"primaryKey"`
	UserName string               `gorm:"type:varchar(30);"`
	Card     ForeignKeyHasOneCard `gorm:"foreignKey:UserID"`
}

type ForeignKeyHasOneCard struct {
	ID       int    `gorm:"primaryKey"`
	CardName string `gorm:"type:varchar(30)"`
	UserID   string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *ForeignKeyHasOneUser) TableName() string {
	return "foreign_key_has_one_user"
}
func (n *ForeignKeyHasOneCard) TableName() string {
	return "foreign_key_has_one_card"
}

type ReferenceHasOneUser struct {
	ID int `gorm:"primaryKey"`
	//外键关联主表的非主键字段必须存在唯一索引
	UserName string              `gorm:"type:varchar(30);unique"`
	Card     ReferenceHasOneCard `gorm:"references:UserName"`
}

type ReferenceHasOneCard struct {
	ID                    int    `gorm:"primaryKey"`
	CardName              string `gorm:"type:varchar(30)"`
	ReferenceHasOneUserID string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *ReferenceHasOneUser) TableName() string {
	return "reference_has_one_user"
}
func (n *ReferenceHasOneCard) TableName() string {
	return "reference_has_one_card"
}

type ReferenceForeignKeyHasOneUser struct {
	ID       int                           `gorm:"primaryKey"`
	UserName string                        `gorm:"type:varchar(30);unique"`
	Card     ReferenceForeignKeyHasOneCard `gorm:"references:UserName;foreignKey:UserName;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ReferenceForeignKeyHasOneCard struct {
	ID       int     `gorm:"primaryKey"`
	CardName string  `gorm:"type:varchar(30)"`
	UserName *string `gorm:"type:varchar(30);"` //foreignKey
	CardInfo string
}

func (n *ReferenceForeignKeyHasOneUser) TableName() string {
	return "reference_foreign_key_has_one_user"
}
func (n *ReferenceForeignKeyHasOneCard) TableName() string {
	return "reference_foreign_key_has_one_card"
}
