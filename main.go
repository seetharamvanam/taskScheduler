package main

import (
	"taskscheduler/database"
	"taskscheduler/scheduler"
	"taskscheduler/source"
	"taskscheduler/tasks"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	err = database.RunMigrationsForCreateSources(db)
	if err != nil {
		panic(err)
	}
	err = database.RunMigrationsForScrapeTasks(db)
	if err != nil {
		panic(err)
	}

	// Creating Repositories
	sourceRepo := source.NewRepository(db)
	taskRepo := tasks.NewRepository(db)

	//Creating a scheduler
	newScheduler := scheduler.NewScheduler(sourceRepo, 5, taskRepo)
	newScheduler.Start()
	defer db.Close()
}
