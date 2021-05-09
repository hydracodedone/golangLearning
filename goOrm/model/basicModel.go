package model

import (
	"database/sql"
	"time"

	"gorm.io/plugin/soft_delete"
)

type Update struct {
	UpdatedAt       time.Time
	UpdatedNano     int64 `gorm:"autoUpdateTime:nano"`  // 使用时间戳填纳秒数充更新时间
	UpdatedMilli    int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
	UpdateTimeStamp int64 `gorm:"autoUpdateTime"`       // 使用时间戳秒数填充创建时间
}
type Delete struct {
	//二者只能支持一种
	// DeletedAt   gorm.DeletedAt        `gorm:"index"`
	DeletedFlag soft_delete.DeletedAt `gorm:"softDelete:flag"`
}
type Create struct {
	CreatedAt       int
	CreateNano      int64 `gorm:"autoCreateTime:nano"`  // 使用时间戳填纳秒数充创建时间
	CreateMilli     int64 `gorm:"autoCreateTime:milli"` // 使用时间戳毫秒数填充创建时间
	CreateTimeStamp int64 `gorm:"autoCreateTime"`       // 使用时间戳秒数填充创建时间
}

type Embed struct {
	Info string
}

type Authority struct {
	OnlyCreate string `gorm:"->:false;<-:create"` //只能创建,不能读取
	OnlyUpdate string `gorm:"->:false；<-:update"` //只能更新，不能读取
	OnlyRead   string `gorm:"<-:false"`           //不能写入，只能读取
}
type NullAndNotNUllData struct {
	NullBool            sql.NullBool
	NormalBool          bool
	NullString          sql.NullString `gorm:"default:'this is not null'"`
	DefaultNormalString string         `gorm:"default:'this is default'"`
	NormalString        string
}
type Basic struct {
	// gorm.Model
	ID      uint  `gorm:"primarykey"`
	Embeded Embed `gorm:"embedded;embeddedPrefix:embed_"`
	Update
	Create
	Delete
	Authority
	NullAndNotNUllData
}

func (b *Basic) TableName() string {
	return "base_model"
}
