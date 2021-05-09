package model

type NormalHasManyUser struct {
	ID       int    `gorm:"primaryKey"`
	UserName string `gorm:"type:varchar(30);"`
	Cards    []NormalHasManyCard
}

type NormalHasManyCard struct {
	ID                  int    `gorm:"primaryKey"`
	CardName            string `gorm:"type:varchar(30)"`
	NormalHasManyUserID string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *NormalHasManyUser) TableName() string {
	return "normal_has_many_user"
}
func (n *NormalHasManyCard) TableName() string {
	return "normal_has_many_card"
}

type ForeignKeyHasManyUser struct {
	ID       int                     `gorm:"primaryKey"`
	UserName string                  `gorm:"type:varchar(30);"`
	Cards    []ForeignKeyHasManyCard `gorm:"foreignKey:UserID"`
}

type ForeignKeyHasManyCard struct {
	ID       int    `gorm:"primaryKey"`
	CardName string `gorm:"type:varchar(30)"`
	UserID   string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *ForeignKeyHasManyUser) TableName() string {
	return "foreign_key_has_many_user"
}
func (n *ForeignKeyHasManyCard) TableName() string {
	return "foreign_key_has_many_card"
}

type ReferenceHasManyUser struct {
	ID int `gorm:"primaryKey"`
	//外键关联主表的非主键字段必须存在唯一索引
	UserName string                 `gorm:"type:varchar(30);unique"`
	Cards    []ReferenceHasManyCard `gorm:"references:UserName"`
}

type ReferenceHasManyCard struct {
	ID                     int    `gorm:"primaryKey"`
	CardName               string `gorm:"type:varchar(30)"`
	ReferenceHasManyUserID string `gorm:"type:varchar(30)"` //foreignKey
}

func (n *ReferenceHasManyUser) TableName() string {
	return "reference_has_many_user"
}
func (n *ReferenceHasManyCard) TableName() string {
	return "reference_has_many_card"
}

type ReferenceForeignKeyHasManyUser struct {
	ID       int                              `gorm:"primaryKey"`
	UserName string                           `gorm:"type:varchar(30);unique"`
	Cards    []ReferenceForeignKeyHasManyCard `gorm:"references:UserName;foreignKey:UserNames;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ReferenceForeignKeyHasManyCard struct {
	ID       int     `gorm:"primaryKey"`
	CardName string  `gorm:"type:varchar(30)"`
	UserNames *string `gorm:"type:varchar(30);"` //foreignKey
	CardInfo string
	User     ReferenceForeignKeyHasManyUser `gorm:"references:UserName;foreignKey:UserNames;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (n *ReferenceForeignKeyHasManyUser) TableName() string {
	return "reference_foreign_key_has_many_user"
}
func (n *ReferenceForeignKeyHasManyCard) TableName() string {
	return "reference_foreign_key_has_many_card"
}
