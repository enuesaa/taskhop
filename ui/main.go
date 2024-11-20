package ui

import (
	"embed"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

//go:generate pnpm install
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

func Handle() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path // like `/`

		// index page
		if strings.HasSuffix(path, "/") {
			path = "/index.html"
		}

		// read dist
		path = fmt.Sprintf("dist%s", path)
		f, err := dist.ReadFile(path)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		// response
		ext := filepath.Ext(path)
		mimeType := mime.TypeByExtension(ext)
		if mimeType == "" {
			mimeType = "application/octet-stream"
		}
		w.Header().Set("Content-Type", mimeType)

		w.WriteHeader(http.StatusOK)
		w.Write(f)
	}

	return handler
}
