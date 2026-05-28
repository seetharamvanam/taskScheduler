package tasks

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) CreateTask(sourceId int, careerUrl string, sourceType string) error {
	_, err := repository.db.Exec("INSERT INTO scrape_tasks("+
		"source_id, "+
		"career_url, "+
		"source_type, "+
		"status, "+
		"created_at) "+
		"VALUES ($1, $2, $3, 'PENDING', NOW());", sourceId, careerUrl, sourceType)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) GetPendingTasks() ([]Task, error) {
	rows, err := repository.db.Query("SELECT id,source_id, career_url, source_type, status, attempt_count, created_at  FROM scrape_tasks WHERE status = 'PENDING';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.SourceID,
			&task.CareerURL,
			&task.SourceType,
			&task.Status,
			&task.AttemptCount,
			&task.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, rows.Err()
}

func (repository *Repository) MarkRunning(taskId int, startedAt time.Time) error {
	_, err := repository.db.Exec("UPDATE scrape_tasks SET status = 'RUNNING', started_at = $1 WHERE id = $2;", startedAt, taskId)
	if err != nil {
		return err
	}
	return nil
}

func (repository *Repository) MarkCompleted(taskId int, completedAt time.Time) error {
	_, err := repository.db.Exec("UPDATE scrape_tasks SET status = 'COMPLETED', completed_at = $1 WHERE id = $2;", completedAt, taskId)
	if err != nil {
		return err
	}
	return nil
}
