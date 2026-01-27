package packages

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/ljgago/sift/internal/markdown"
	"github.com/ljgago/sift/internal/tgz"
)

type Service struct {
	repository Repository
}

func (s *Service) Package(ctx context.Context, name string) (*[]byte, error) {
	pkg, err := s.repository.Package(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("Service GetPackage -> %w", err)
	}

	u, _ := url.Parse(pkg.Dist.Tarball)
	imgURL, _ := parseImgURL(pkg.Repository.URL)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Service GetPackage -> %w", err)
	}
	defer resp.Body.Close()

	markdownCondent, _ := tgz.ReadFile(resp.Body, "README.md")
	// if err != nil {
	// 	return nil, fmt.Errorf("Service GetPackage -> %w", err)
	// }

	htmlContent, err := markdown.ToHTML(markdownCondent, imgURL)
	if err != nil {
		return nil, fmt.Errorf("Service GetPackage -> %w", err)
	}

	content := markdown.ProcessFinalHTML(string(htmlContent), imgURL)
	byteSlice := []byte(content)

	return &byteSlice, nil
	// return &htmlContent, nil
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func parseImgURL(rawURL string) (string, error) {
	if rawURL == "" {
		return "", nil
	}

	cleanURL := strings.TrimPrefix(rawURL, "git+")
	cleanURL = strings.TrimSuffix(cleanURL, ".git")

	domains := map[string]string{
		"github.com": "https://raw.githubusercontent.com",
	}

	u, err := url.Parse(cleanURL)
	if err != nil {
		return "", err
	}

	path := strings.Split(strings.TrimPrefix(u.Path, "/"), "/")

	domain := fmt.Sprintf("%s://%s", u.Scheme, u.Host)
	user := path[0]
	repo := path[1]

	if value, ok := domains[u.Host]; ok {
		domain = value
	}

	finalURL := fmt.Sprintf("%s/%s/%s/HEAD/", domain, user, repo)

	return finalURL, nil
}
