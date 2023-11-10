package util

import "net/http"

func CheckOrigin(allowedOrigins []string) func(*http.Request) bool {
	return func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		for _, o := range allowedOrigins {
			if origin == o {
				return true
			}
		}
		return false
	}
}
