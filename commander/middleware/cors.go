package middleware

import "github.com/go-chi/cors"

func Cors() Fn {
	handle := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	return handle
}
