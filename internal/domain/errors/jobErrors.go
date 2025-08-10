package domainErrors

import "errors"

var (
	ErrJobNotFound             = errors.New("job not found")
	ErrJobUUIDRequired         = errors.New("job uuid is required")
	ErrJobTitleRequired        = errors.New("job title is required")
	ErrJobDescriptionRequired  = errors.New("job description is required")
	ErrJobDescriptionTooShort  = errors.New("job description must be at least 20 characters long")
	ErrJobLocationRequired     = errors.New("job location is required")
	ErrJobCompanyIDRequired    = errors.New("job companyID is required")
	ErrJobCreatedByRequired    = errors.New("job createdBy is required")
	ErrJobCreatedAtRequired    = errors.New("job createdAt is required")
	ErrJobMissingRequiredField = errors.New("missing required field")
)
