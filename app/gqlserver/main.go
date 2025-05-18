package gqlserver

import (
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
)

func NewHandler(li *lib.Lib, config *conf.Config) *Handler {
	return &Handler{
		li: li,
		config: config,
	}
}

type Handler struct {
	li *lib.Lib
	config *conf.Config
}
