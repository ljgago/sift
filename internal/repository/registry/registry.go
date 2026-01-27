package registry

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ljgago/sift/internal/search"
)

const SearchAPI = "/-/v1/search"

type Registry struct {
	URL string
}

func (r *Registry) Packages(ctx context.Context, args *search.QueryArgs) (*SearchPackageResponse, error) {
	u, _ := url.Parse(r.URL)
	u = u.JoinPath(SearchAPI)

	q := u.Query()
	q.Add("text", args.Text)
	q.Add("from", strconv.Itoa(args.From))
	q.Add("size", strconv.Itoa(args.Size))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Registry Packages: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Registry Packages: %s", resp.Status)
	}
	defer resp.Body.Close()

	searchPackageResponse := SearchPackageResponse{}

	err = json.NewDecoder(resp.Body).Decode(&searchPackageResponse)
	if err != nil {
		return nil, fmt.Errorf("Registry Packages: %w", err)
	}

	return &searchPackageResponse, nil
}

func (r *Registry) PackageMetadata(ctx context.Context, name string) (*PackageMetadataResponse, error) {
	u, _ := url.Parse(r.URL)
	u = u.JoinPath(name)

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Registry PackageMetadata: %w", err)
	}
	defer resp.Body.Close()

	packageMetadataResponse := PackageMetadataResponse{}

	err = json.NewDecoder(resp.Body).Decode(&packageMetadataResponse)
	if err != nil {
		return nil, fmt.Errorf("Registry PackageMetadata: %w", err)
	}

	return &packageMetadataResponse, nil
}

func (r *Registry) Package(ctx context.Context, name string) (*PackageResponse, error) {
	u, _ := url.Parse(r.URL)
	u = u.JoinPath(name).JoinPath("latest")

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("Registry Package: %w", err)
	}
	defer resp.Body.Close()

	packageResponse := PackageResponse{}

	err = json.NewDecoder(resp.Body).Decode(&packageResponse)
	if err != nil {
		return nil, fmt.Errorf("Registry Package: %w", err)
	}

	return &packageResponse, nil
}

func New(url string) *Registry {
	return &Registry{
		URL: url,
	}
}
