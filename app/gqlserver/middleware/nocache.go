package middleware

import "github.com/go-chi/chi/v5/middleware"

func NoCache() Middleware {
	return middleware.NoCache
}
