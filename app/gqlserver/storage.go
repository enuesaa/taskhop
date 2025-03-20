package gqlserver

import (
	"io"
	"net/http"

	"github.com/enuesaa/taskhop/lib"
)

func HandleStorage(li lib.Lib) http.HandlerFunc {
	handle := func(rw http.ResponseWriter, r *http.Request) {
		archive, err := li.Arv.Archive()
		if err != nil {
			li.Log.Info(r.Context(), "failed to create archive file: %s", err.Error())
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/zip")
		rw.Header().Set("Content-Disposition", "attachment; filename=\"archive.zip\"")

		if _, err := io.Copy(rw, archive); err != nil {
			li.Log.Info(r.Context(), "failed to write archive to response: %s", err.Error())
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		}
	}
	return handle
}
