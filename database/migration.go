package database

import "database/sql"

func RunMigrations(db *sql.DB) error {
	query := "CREATE TABLE IF NOT EXISTS sources(" +
		"id SERIAL PRIMARY KEY, " +
		"company_name TEXT NOT NULL, " +
		"career_url TEXT NOT NULL, " +
		"source_type TEXT NOT NULL, " +
		"scrape_frequency_minutes INT NOT NULL, " +
		"is_active BOOLEAN NOT NULL DEFAULT TRUE, " +
		"next_scrape_at TIMESTAMP NOT NULL, " +
		"last_scrape_at TIMESTAMP, " +
		"last_scheduled_at TIMESTAMP, " +
		"created_at TIMESTAMP NOT NULL DEFAULT NOW(), " +
		"updated_at TIMESTAMP NOT NULL DEFAULT NOW())"
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
