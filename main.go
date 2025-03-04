package main

import (
	"report-birthday/cmd"
	"report-birthday/db"
)

func main() {
	database := db.ConnectDB()
	defer database.Close()
	db.CheckBirthdays(database)
	cmd.Execute()
}