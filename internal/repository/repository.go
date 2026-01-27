package repository

import (
	"github.com/ljgago/sift/internal/config"
	"github.com/ljgago/sift/internal/repository/github"
	"github.com/ljgago/sift/internal/repository/registry"
)

const (
	GithubBaseURL = "https://api.github.com/repos"
)

type Repository struct {
	Registry *registry.Registry
	Github   *github.Github
}

func New(cfg config.Config) *Repository {
	return &Repository{
		Registry: registry.New(cfg.Server.URL),
		Github:   github.New(GithubBaseURL),
	}
}
