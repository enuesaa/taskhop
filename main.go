package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	app := chi.NewRouter()
	app.Use(middleware.Logger)
	app.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	if err := http.ListenAndServe(":3000", app); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}
