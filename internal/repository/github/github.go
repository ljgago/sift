package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Github struct {
	URL string
}

func (g *Github) Stats(ctx context.Context, owner string, repo string) (*RepoStats, error) {
	u, _ := url.Parse(g.URL)
	u = u.JoinPath(owner).JoinPath(repo)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Github GetStars: %w", err)
	}
	defer resp.Body.Close()

	repoStats := RepoStats{}

	err = json.NewDecoder(resp.Body).Decode(&repoStats)
	if err != nil {
		return nil, fmt.Errorf("repository Github GetRepoStats: %w", err)
	}

	return &repoStats, nil
}

func New(url string) *Github {
	return &Github{
		URL: url,
	}
}
