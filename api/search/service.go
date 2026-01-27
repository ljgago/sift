package search

import (
	"context"
	"fmt"
	"math"

	"github.com/ljgago/sift/internal/search"
)

type Service struct {
	repository Repository
}

func (s *Service) Packages(ctx context.Context, query string) (*Result, error) {
	args, err := search.ParseQuery(query)
	if err != nil {
		return nil, fmt.Errorf("Service ListPackages: %w", err)
	}

	response, err := s.repository.Packages(ctx, args)
	if err != nil {
		return nil, fmt.Errorf("Service ListPackages: %w", err)
	}

	totalItems := response.Total
	totalPages := int(math.Ceil(float64(totalItems) / float64(args.Size)))
	items := response.Objects

	result := Result{
		Page:       args.Page,
		Size:       args.Size,
		TotalPages: totalPages,
		TotalItems: totalItems,
		Items:      items,
	}

	return &result, nil
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}
