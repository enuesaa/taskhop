package middleware

import "net/http"

type Fn func(http.Handler) http.Handler
