package search

type Result struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
	Items      any `json:"items"`
}

type Query struct {
	Text        string
	Size        string
	From        string
	Quality     string
	Popularity  string
	Maintenance string
}
