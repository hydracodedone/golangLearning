package session

import (
	"context"

	"gorm.io/gorm"
)

func NewSession(ctx context.Context,db *gorm.DB) *gorm.DB {
	sessionConfig := &gorm.Session{
		DryRun:                   false,
		PrepareStmt:              false,
		NewDB:                    false,
		Initialized:              false,
		SkipHooks:                false,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: false,
		AllowGlobalUpdate:        true,
		FullSaveAssociations:     true,
		//不需一定声明查询字段,可以使用 *
		QueryFields:     false,
		CreateBatchSize: 50,
		Logger:         nil,
		Context: ctx,
	}
	if db == nil {
		return db.Session(sessionConfig)
	}
	return nil
}
