package worker

import (
	"fmt"
	"taskscheduler/tasks"
	"time"
)

type Worker struct {
	taskRepo        *tasks.Repository
	pollingInterval time.Duration
}

func NewWorker(taskRepo *tasks.Repository, pollingInterval time.Duration) *Worker {
	return &Worker{taskRepo: taskRepo, pollingInterval: pollingInterval}
}

func (w *Worker) Start() {
	ticker := time.NewTicker(w.pollingInterval)
	defer ticker.Stop()
	for tick := range ticker.C {
		tasksList, err := w.taskRepo.GetPendingTasks()
		if err != nil {
			fmt.Println("Error getting pending tasks:", err)
			continue
		}
		for _, task := range tasksList {
			err = w.taskRepo.MarkRunning(task.ID, tick)
			if err != nil {
				fmt.Println("Error marking task as running:", err)
				continue
			}
			fmt.Println(task)
			time.Sleep(2 * time.Second)
			err = w.taskRepo.MarkCompleted(task.ID, time.Now())
			if err != nil {
				fmt.Println("Error marking task as completed:", err)
				continue
			}

		}
	}
}
