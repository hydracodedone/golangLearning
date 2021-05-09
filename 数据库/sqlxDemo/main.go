package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func initDB() (db *sqlx.DB) {
	dbUser := "root"
	dbPwd := "123456"
	dbHost := "localhost"
	dbName := "user"
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbName)
	db, err := sqlx.Connect("mysql", dns)
	if err != nil {
		fmt.Printf("ERROR:OPEN DATABASE FAIL->%s", err)
		return nil
	}
	return db
}
func closeDB(db *sqlx.DB) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		fmt.Printf("ERROR:CLOSE DATABASES FAIL->%s", err)
	}
}
func checkConnectionDb(db *sqlx.DB) bool {
	if db == nil {
		return false
	}
	err := db.Ping()
	if err != nil {
		fmt.Printf("ERROR:CLOSE DATABASES FAIL->%s", err)
		return false
	}
	return true
}

func main() {

}
