package jobEntity

import "time"

type Job struct {
	UUID        string
	Title       string
	Description string
	Location    string
	CreatedBy   string
	CreatedAt   time.Time
}
