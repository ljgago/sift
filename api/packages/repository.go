package packages

import (
	"context"

	"github.com/ljgago/sift/internal/repository/registry"
)

type Repository interface {
	PackageMetadata(ctx context.Context, name string) (*registry.PackageMetadataResponse, error)
	Package(ctx context.Context, name string) (*registry.PackageResponse, error)
}
