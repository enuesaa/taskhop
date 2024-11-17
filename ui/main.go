package ui

import (
	"embed"
	"fmt"
	"mime"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:generate pnpm install
//go:generate pnpm gen:gql
//go:generate pnpm build

//go:embed all:dist/*
var dist embed.FS

func Handle() echo.HandlerFunc {
	handler := func(c echo.Context) error {
		path := c.Request().URL.Path // like `/`

		// index page
		ext := filepath.Ext(path)
		if ext == "" || strings.HasSuffix(path, "/") {
			path = "/index.html"
		}

		// read dist
		path = fmt.Sprintf("dist%s", path)
		f, err := dist.ReadFile(path)
		if err != nil {
			return err
		}

		// response
		mimeType := mime.TypeByExtension(ext)

		return c.Blob(200, mimeType, f)
	}

	return handler
}
