package stars

import (
	"context"
	"fmt"
)

type Service struct {
	repository Repository
}

func (s *Service) Stars(ctx context.Context, owner string, repo string) (*Stars, error) {
	stats, err := s.repository.Stats(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("Service Stars: %w", err)
	}

	stars := Stars{
		RepositoryName: stats.FullName,
		StarsCount:     stats.StargazersCount,
		UpdatedAt:      stats.UpdatedAt,
	}

	return &stars, nil
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
