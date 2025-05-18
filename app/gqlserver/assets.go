package gqlserver

import (
	"io"
	"net/http"

	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
)

func HandleAssets(li lib.Lib, config *conf.Config) http.HandlerFunc {
	handle := func(rw http.ResponseWriter, r *http.Request) {
		archive, err := li.Arv.Archive(config.Workdir)
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

func HandleUpload(li lib.Lib, config *conf.Config) http.HandlerFunc {
	handle := func(rw http.ResponseWriter, r *http.Request) {
		if err := li.Arv.UnArchive(r.Body, config.Workdir); err != nil {
			li.Log.Info(r.Context(), "failed to download archive file: %s", err.Error())
			http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
	return handle
}
