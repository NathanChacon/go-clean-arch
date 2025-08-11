package jobEntity

import (
	"time"

	domainErrors "jobs.api.com/internal/domain/errors"
)

type Job struct {
	UUID        string    `db:"uuid"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Location    string    `db:"location"`
	CompanyID   string    `db:"company_id"`
	CreatedBy   string    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
}

type NewJobParams struct {
	UUID        string
	Title       string
	Description string
	Location    string
	CompanyID   string
	CreatedBy   string
}

func NewJobEntity(params NewJobParams) (*Job, error) {
	if params.UUID == "" {
		return nil, domainErrors.ErrJobUUIDRequired
	}
	if params.Title == "" {
		return nil, domainErrors.ErrJobTitleRequired
	}
	if params.Description == "" {
		return nil, domainErrors.ErrJobDescriptionRequired
	}
	if len(params.Description) < 20 {
		return nil, domainErrors.ErrJobDescriptionTooShort
	}
	if params.Location == "" {
		return nil, domainErrors.ErrJobLocationRequired
	}
	if params.CompanyID == "" {
		return nil, domainErrors.ErrJobCompanyIDRequired
	}
	if params.CreatedBy == "" {
		return nil, domainErrors.ErrJobCreatedByRequired
	}

	return &Job{
		UUID:        params.UUID,
		Title:       params.Title,
		Description: params.Description,
		Location:    params.Location,
		CompanyID:   params.CompanyID,
		CreatedBy:   params.CreatedBy,
	}, nil
}
