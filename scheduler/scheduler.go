package scheduler

import (
	"fmt"
	"taskscheduler/source"
	"taskscheduler/tasks"
	"time"
)

type Scheduler struct {
	sourceRepo      *source.Repository
	taskRepo        *tasks.Repository
	pollingInterval time.Duration
}

func NewScheduler(sourceRepo *source.Repository,
	pollingInterval time.Duration, taskRepo *tasks.Repository) *Scheduler {
	return &Scheduler{
		sourceRepo:      sourceRepo,
		taskRepo:        taskRepo,
		pollingInterval: pollingInterval,
	}
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(s.pollingInterval)
	defer ticker.Stop()
	for tick := range ticker.C {
		sources, err := s.sourceRepo.GetDueSources(tick)
		if err != nil {
			fmt.Println("error getting due sources:", err)
			continue
		}
		if len(sources) == 0 {
			fmt.Println("no sources found")
		}
		for _, src := range sources {
			err := s.taskRepo.CreateTask(src.ID, src.CareerURL, src.SourceType)
			if err != nil {
				fmt.Println("error creating task:", err)
				continue
			}
			nextScrapeTime := tick.Add(time.Duration(src.ScrapeFrequencyMinutes) * time.Minute)
			err = s.sourceRepo.MarkScheduled(src.ID, tick, nextScrapeTime)
			if err != nil {
				fmt.Println("error marking source as scheduled:", err)
				continue
			}
		}
	}
}
