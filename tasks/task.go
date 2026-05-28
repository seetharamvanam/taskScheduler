package tasks

import "time"

type Task struct {
	ID           int
	SourceID     int
	CareerURL    string
	SourceType   string
	Status       string
	AttemptCount int
	CreatedAt    time.Time
	StartedAt    *time.Time
	CompletedAt  *time.Time
	ErrorMessage string
}
