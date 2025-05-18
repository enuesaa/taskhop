package gqlserver

import (
	"io"
	"net/http"
)

func (h *Handler) Assets(rw http.ResponseWriter, r *http.Request) {
	archive, err := h.li.Arv.Archive(h.config.Workdir)
	if err != nil {
		h.li.Log.Info(r.Context(), "failed to create archive file: %s", err.Error())
		rw.WriteHeader(500)
		return
	}
	rw.Header().Set("Content-Type", "application/zip")
	rw.Header().Set("Content-Disposition", "attachment; filename=\"archive.zip\"")

	if _, err := io.Copy(rw, archive); err != nil {
		h.li.Log.Info(r.Context(), "failed to write archive to response: %s", err.Error())
		rw.WriteHeader(500)
	}
}

func (h *Handler) Upload(rw http.ResponseWriter, r *http.Request) {
	if err := h.li.Arv.UnArchive(r.Body, h.config.Workdir); err != nil {
		h.li.Log.Info(r.Context(), "failed to download archive file: %s", err.Error())
		rw.WriteHeader(500)
		return
	}
}
