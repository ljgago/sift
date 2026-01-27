package stars

import "time"

type Stars struct {
	RepositoryName string     `json:"repository_name"`
	StarsCount     int        `json:"stars_count"`
	UpdatedAt      *time.Time `json:"updated_at"`
}
