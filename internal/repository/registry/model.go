package registry

import "time"

type SearchPackageResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Objects []struct {
		Downloads struct {
			Monthly int `json:"monthly"`
			Weekly  int `json:"weekly"`
		} `json:"downloads"`
		Dependents  any       `json:"dependents"`
		Updated     time.Time `json:"updated"`
		SearchScore float64   `json:"searchScore"`
		Package     struct {
			Name          string   `json:"name"`
			Keywords      []string `json:"keywords"`
			Version       string   `json:"version"`
			Description   string   `json:"description"`
			SanitizedName string   `json:"sanitized_name"`
			Publisher     struct {
				Email string `json:"email"`
				Actor struct {
					Name  string `json:"name"`
					Type  string `json:"type"`
					Email string `json:"email"`
				} `json:"actor"`
				TrustedPublisher struct {
					OidcConfigID string `json:"oidcConfigId"`
					ID           string `json:"id"`
				} `json:"trustedPublisher"`
				Username string `json:"username"`
			} `json:"publisher"`
			Maintainers []struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			} `json:"maintainers"`
			License string    `json:"license"`
			Date    time.Time `json:"date"`
			Links   struct {
				Homepage   string `json:"homepage"`
				Repository string `json:"repository"`
				Bugs       string `json:"bugs"`
				Npm        string `json:"npm"`
			} `json:"links"`
		} `json:"package"`
		Score struct {
			Final  float64 `json:"final"`
			Detail struct {
				Popularity  int `json:"popularity"`
				Quality     int `json:"quality"`
				Maintenance int `json:"maintenance"`
			} `json:"detail"`
		} `json:"score"`
		Flags struct {
			Insecure int `json:"insecure"`
		} `json:"flags"`
	} `json:"objects,omitempty"`
	Total int        `json:"total,omitempty"`
	Time  *time.Time `json:"time,omitempty"`
}

type PackageMetadataResponse struct {
	ReadmeFilename string `json:"readmeFilename"`
}

type PackageResponse struct {
	Dist struct {
		Tarball string `json:"tarball"`
	} `json:"dist"`
	Repository struct {
		URL string `json:"url"`
	} `json:"repository"`
}
