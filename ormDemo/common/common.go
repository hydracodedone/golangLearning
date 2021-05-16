package common

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./ormDemo/gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Init DB Fail :<%s>\n", err.Error())
	}

	return db
}

func SetDB(db *gorm.DB) {
	dbConfig, err := db.DB()
	if err != nil {
		log.Fatalf("Config DB Fail :<%s>\n", err.Error())
	}
	dbConfig.SetMaxIdleConns(2)
	dbConfig.SetMaxOpenConns(1)
	dbConfig.SetConnMaxIdleTime(time.Minute)
	dbConfig.SetConnMaxLifetime(time.Minute)
}

func MigrateDB(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("AutoMigrate DB Fail :<%s>\n", err.Error())
	}
}

func CleanDB(db *gorm.DB) {
}
func CloseDB(db *gorm.DB) {
	dbConfig, err := db.DB()
	if err != nil {
		log.Fatalf("Config DB Fail :<%s>\n", err.Error())
	}
	err = dbConfig.Close()
	if err != nil {
		log.Fatalf("Close DB Fail :<%s>\n", err.Error())
	}
}
