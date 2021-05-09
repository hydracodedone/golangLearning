package session

import (
	"database_demo/connection"
	"database_demo/model"
	"fmt"

	"gorm.io/gorm"
)

/*
	type Session struct {
		DryRun                 bool
		PrepareStmt            bool
		NewDB                  bool
		SkipHooks              bool
		SkipDefaultTransaction bool
		AllowGlobalUpdate      bool
		FullSaveAssociations   bool
		Context                context.Context
		Logger                 logger.Interface
		NowFunc                func() time.Time
	}
*/
func _() {
	sessionDB := connection.GormDB.Session(&gorm.Session{
		DryRun: true,
	})
	statement := sessionDB.Model(&model.Basic{}).First("id=?", 1).Statement
	fmt.Println(statement.SQL.String())
	fmt.Println(statement.Vars)
	fmt.Println(sessionDB.Dialector.Explain(statement.SQL.String(), statement.Vars...))
}
func sessionPrepareStatement() {
	sessionDB := connection.GormDB.Session(&gorm.Session{
		PrepareStmt: true,
	})
	sessionDB.Model(&model.Basic{}).First("id=?", 1)
	stmtManger, ok := sessionDB.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		return
	}
	preparedSQL := stmtManger.PreparedSQL
	fmt.Println(preparedSQL)
	sessionDB.Model(&model.Basic{}).First("id=?", 2)
}
func BasicHandleSession() {
	sessionPrepareStatement()
}
