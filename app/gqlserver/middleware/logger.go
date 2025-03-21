package middleware

import (
	"net/http"
	"time"

	"github.com/enuesaa/taskhop/lib"
)

func Logger(li lib.Lib) Fn {
	handle := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			li.Log.Info(r.Context(), "%s %s %s", r.Method, r.URL.Path, time.Since(start))
		})
	}
	return handle
}
