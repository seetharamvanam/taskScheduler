package source

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetActiveSources() ([]Source, error) {
	rows, err := r.db.Query("SELECT id, " +
		"company_name,career_url, source_type, scrape_frequency_minutes, " +
		"is_active, next_scrape_at  FROM sources WHERE is_active = TRUE;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sources []Source
	for rows.Next() {
		var source Source

		err := rows.Scan(
			&source.ID,
			&source.CompanyName,
			&source.CareerURL,
			&source.SourceType,
			&source.ScrapeFrequencyMinutes,
			&source.IsActive,
			&source.NextScrapeAt,
		)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, rows.Err()
}

func (r *Repository) GetDueSources(now time.Time) ([]Source, error) {
	rows, err := r.db.Query("SELECT id, company_name,career_url, source_type, scrape_frequency_minutes, is_active, next_scrape_at  FROM sources WHERE is_active = TRUE AND next_scrape_at <= $1;", now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sources []Source
	for rows.Next() {
		var source Source

		err := rows.Scan(
			&source.ID,
			&source.CompanyName,
			&source.CareerURL,
			&source.SourceType,
			&source.ScrapeFrequencyMinutes,
			&source.IsActive,
			&source.NextScrapeAt,
		)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, rows.Err()
}
