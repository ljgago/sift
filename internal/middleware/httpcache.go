package middleware

import (
	"net/http"
)

// HTTPCache is a simple middleware that sets HTTP cache headers.
func HTTPCache(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Cache valid for 5 min (max-age=300), but after this time for 24 hs (stale-while-revalidate=86400),
		// the cache can be revalidated in the background
		w.Header().Set("Cache-Control", "private, max-age=300, stale-while-revalidate=86400")

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
