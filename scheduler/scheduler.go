package scheduler

import (
	"taskscheduler/source"
	"time"
)

type Scheduler struct {
	repo            *source.Repository
	pollingInterval time.Duration
}

func NewScheduler(repo *source.Repository, pollingInterval time.Duration) *Scheduler {
	return &Scheduler{
		repo:            repo,
		pollingInterval: pollingInterval,
	}
}

func (s *Scheduler) Start() {

}
