package db

import (
	"database/sql"
	"log"
	"time"
)

// DBに接続する
func ConnectDB() *sql.DB{
	db, err := sql.Open("sqlite3", "/test.db")
	if err != nil {
        log.Fatal(err)
    }
	return db
}

// 誕生日をチェックして表示する
func CheckBirthdays(db *sql.DB) {
	today := time.Now().Format("01-02")
	rows, err := db.Query("SELECT * FROM birthdays WHERE birthday = ?", today)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 一致する誕生日がある場合は表示する
	for rows.Next() {
		var id int
		var name string
		var birthday string
		if err := rows.Scan(&id, &name, &birthday); err != nil {
			log.Fatal(err)
		}
		log.Printf("Happy Birthday! %s\n", name)
	}
}

// DBを切断する
func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalln(err)
	}
}