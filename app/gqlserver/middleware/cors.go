package middleware

import "github.com/go-chi/cors"

func Cors() Middleware {
	handle := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	return handle
}
