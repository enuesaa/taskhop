package storage

import (
	"io"
	"net/http"

	"github.com/enuesaa/taskhop/internal"
)

func Handle(c internal.Container) http.HandlerFunc {
	handle := func(rw http.ResponseWriter, r *http.Request) {
		archive, err := c.Arvi.Archive()
		if err != nil {
			c.Logi.Info(r.Context(), "failed to create archive file: %s", err.Error())
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/zip")
		rw.Header().Set("Content-Disposition", "attachment; filename=\"archive.zip\"")
	
		if _, err := io.Copy(rw, archive); err != nil {
			c.Logi.Info(r.Context(), "failed to write archive to response: %s", err.Error())
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		}
	}
	return handle
}
