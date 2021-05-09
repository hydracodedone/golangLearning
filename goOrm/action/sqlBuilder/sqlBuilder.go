package sqlbuilder

import (
	"database_demo/connection"
	"fmt"

	"gorm.io/gorm"
)

func _() {
	dryRunSession := connection.GormDB.Session(&gorm.Session{DryRun: true})
	statement := dryRunSession.Exec("select * from demo.basics where id = ? limit ?", 1, 2).Statement
	sqlString := statement.SQL.String()
	sqlVars := statement.Vars
	fmt.Printf("the sql is [%s],the param is [%v]\n", sqlString, sqlVars)
	fmt.Printf("the actual execurte sql is [%s]\n", connection.GormDB.Dialector.Explain(sqlString, sqlVars...))
}
