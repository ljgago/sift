package github

import "time"

type RepoStats struct {
	FullName        string     `json:"full_name,omitempty"`
	StargazersCount int        `json:"stargazers_count,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
	URL             string     `json:"url,omitempty"`
}
