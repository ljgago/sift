package middleware

import (
	"crypto/rand"
	"net/http"
	"sync"

	"github.com/gorilla/csrf"

	"github.com/ljgago/sift/internal/log"
)

var (
	once sync.Once
	key  = make([]byte, 32)
)

func CSRF(next http.Handler) http.Handler {
	once.Do(func() {
		_, err := rand.Read(key)
		if err != nil {
			log.Errorf("Error generating random key: %s", err.Error())
		}
	})

	CSRF := csrf.Protect(key, csrf.CookieName("_csrf_token"))
	return CSRF(next)
}
