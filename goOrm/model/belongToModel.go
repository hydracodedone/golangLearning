package model

type NormalBelongToCompany struct {
	SID         int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);"`
}

type NormalBelongToUser struct {
	ID                       int    `gorm:"primaryKey"`
	UserName                 string `gorm:"type:varchar(30)"`
	NormalBelongToCompanySID string `gorm:"type:varchar(30)"` //foreignKey
	NormalBelongToCompany    NormalBelongToCompany
}

func (n *NormalBelongToCompany) TableName() string {
	return "normal_belong_to_company"
}
func (n *NormalBelongToUser) TableName() string {
	return "normal_belong_to_user"
}

type ForeignKeyBelongToCompany struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);"`
}

type ForeignKeyBelongToUser struct {
	ID        int                       `gorm:"primaryKey"`
	UserName  string                    `gorm:"type:varchar(30)"`
	CompanyId string                    `gorm:"type:varchar(30)"` //foreignKey
	Company   ForeignKeyBelongToCompany `gorm:"foreignKey:CompanyId"`
}

func (n *ForeignKeyBelongToCompany) TableName() string {
	return "foreign_key_belong_to_company"
}
func (n *ForeignKeyBelongToUser) TableName() string {
	return "foreign_key_belong_to_user"
}

type ReferenceBelongToCompany struct {
	ID          int    `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);unique"`
}

type ReferenceBelongToUser struct {
	ID                         int                      `gorm:"primaryKey"`
	UserName                   string                   `gorm:"type:varchar(30)"`
	ReferenceBelongToCompanyID string                   `gorm:"type:varchar(30)"` //foreignKey
	ReferenceBelongToCompany   ReferenceBelongToCompany `gorm:"references:CompanyName"`
}

func (n *ReferenceBelongToCompany) TableName() string {
	return "reference_belong_to_company"
}
func (n *ReferenceBelongToUser) TableName() string {
	return "reference_belong_to_user"
}

type ReferenceForeignKeyBelongToCompany struct {
	ID             int    `gorm:"primaryKey"`
	CompanyName    string `gorm:"type:varchar(30);unique"`
	CompanyAddress string
}

type ReferenceForeignKeyBelongToUser struct {
	ID          int                                `gorm:"primaryKey"`
	UserName    string                             `gorm:"type:varchar(30)"`
	CompanyInfo *string                            `gorm:"type:varchar(30);"` //foreignKey
	Company     ReferenceForeignKeyBelongToCompany `gorm:"references:CompanyName;foreignKey:CompanyInfo;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (n *ReferenceForeignKeyBelongToCompany) TableName() string {
	return "reference_foreign_key_belong_to_company"
}
func (n *ReferenceForeignKeyBelongToUser) TableName() string {
	return "reference_foreign_key_belong_to_user"
}

type ErrorReferenceForeignKeyBelongToCompany struct {
	ID          string `gorm:"primaryKey"`
	CompanyName string `gorm:"type:varchar(30);"`
}

type ErrorReferenceForeignKeyBelongToUser struct {
	ID          int                                     `gorm:"primaryKey"`
	UserName    string                                  `gorm:"type:varchar(30)"`
	CompanyName string                                  `gorm:"type:varchar(30);unique"` //foreignKey
	Company     ErrorReferenceForeignKeyBelongToCompany `gorm:"foreignKey:CompanyName;references:CompanyName;"`
}

func (n *ErrorReferenceForeignKeyBelongToCompany) TableName() string {
	return "error_reference_foreign_key_belong_to_company"
}
func (n *ErrorReferenceForeignKeyBelongToUser) TableName() string {
	return "error_reference_foreign_key_belong_to_user"
}
