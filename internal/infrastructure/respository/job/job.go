package jobRepository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	jobEntity "jobs.api.com/internal/domain/entities/job"
)

type JobMySQLRepository struct {
	db *sqlx.DB
}

func NewJobMySQLRepository(db *sqlx.DB) *JobMySQLRepository {
	return &JobMySQLRepository{db: db}
}

func (repository *JobMySQLRepository) Create(jobPayload jobEntity.Job) error {
	query := `
        INSERT INTO jobs (uuid, title, description, location, created_by)
        VALUES (?, ?, ?, ?, ?)
    `
	_, err := repository.db.Exec(query,
		jobPayload.UUID,
		jobPayload.Title,
		jobPayload.Description,
		jobPayload.Location,
		jobPayload.CreatedBy,
	)
	return err
}

func (repository *JobMySQLRepository) List() ([]jobEntity.Job, error) {
	jobs := []jobEntity.Job{}
	err := repository.db.Select(&jobs, `SELECT * FROM jobs ORDER BY created_at DESC`)
	return jobs, err
}

func (repository *JobMySQLRepository) GetByID(id string) (*jobEntity.Job, error) {
	var j jobEntity.Job
	err := repository.db.Get(&j, `SELECT * FROM jobs WHERE id = ?`, id)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &j, err
}
