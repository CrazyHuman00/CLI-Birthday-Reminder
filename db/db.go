package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DBに接続する
func ConnectDB() *gorm.DB{
	url := fmt.Sprintf("sqlite3://%s", "test.db")
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
        log.Fatalln(err)
    }
	return db
}

// DBを切断する
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}