package src

import "net/http"

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		w.Header().Set("x-content-type-options", "nosniff")
		next.ServeHTTP(w, r)
	})
}
