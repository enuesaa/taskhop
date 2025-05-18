package gqlclient

import (
	"io"

	"github.com/enuesaa/taskhop/app/gqlclient/adapter"
	"github.com/enuesaa/taskhop/app/gqlclient/logwriter"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
)

func New(config *conf.Config, li *lib.Lib) *UseCase {
	adap := adapter.New(config.Address)

	u := &UseCase{
		li:        li,
		config:    config,
		adap:      adap,
		LogWriter: logwriter.New(adap),
	}
	return u
}

type UseCase struct {
	li        *lib.Lib
	config    *conf.Config
	adap      *adapter.Adapter
	LogWriter io.Writer
}
