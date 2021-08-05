package middleware

import (
	"net/http"
)

func Default(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")

		if (*r).Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			r.Context().Done()
			return
		}

		ctx := r.Context()
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
