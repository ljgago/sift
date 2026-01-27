package stars

import (
	"context"

	"github.com/ljgago/sift/internal/repository/github"
)

type Repository interface {
	Stats(ctx context.Context, owner string, repo string) (*github.RepoStats, error)
}
