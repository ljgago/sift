package search

import (
	"context"

	"github.com/ljgago/sift/internal/repository/registry"
	"github.com/ljgago/sift/internal/search"
)

type Repository interface {
	Packages(ctx context.Context, args *search.QueryArgs) (*registry.SearchPackageResponse, error)
}
