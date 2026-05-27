package main

import (
	"taskscheduler/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	err = database.RunMigrations(db)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
