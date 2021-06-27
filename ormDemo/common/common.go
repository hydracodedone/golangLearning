package common

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB(path string) *gorm.DB {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("Init DB Fail :<%s>\n", err.Error())
	}
	db.Debug()
	return db
}

func SetDB(db *gorm.DB) {
	dbConfig := db.DB()

	err := dbConfig.Ping()
	if err != nil {
		log.Fatalf("Ping DB Fail :<%s>\n", err.Error())

	}
	dbConfig.SetMaxIdleConns(2)
	dbConfig.SetMaxOpenConns(1)
	dbConfig.SetConnMaxIdleTime(time.Minute)
	dbConfig.SetConnMaxLifetime(time.Minute)
}

func MigrateDB(db *gorm.DB, models ...interface{}) {
	db.AutoMigrate(models...)
}

func CleanDB(db *gorm.DB) {

}
func CloseDB(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Close DB Fail:%s\n", err.Error())
	}
}
