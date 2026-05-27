package source

import "time"

type Source struct {
	ID                     int
	CompanyName            string
	CareerURL              string
	SourceType             string
	ScrapeFrequencyMinutes int
	IsActive               bool
	NextScrapeAt           time.Time
	LastScrapedAt          *time.Time
	LastScheduledAt        *time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}
