package gqlclient

import (
	"io"

	"github.com/enuesaa/taskhop/app/gqlclient/adapter"
	"github.com/enuesaa/taskhop/app/gqlclient/logwriter"
	"github.com/enuesaa/taskhop/conf"
)

func New(config *conf.Config) *Connector {
	adap := adapter.New(config.Address)

	connector := &Connector{
		config: config,
		adap:   adap,
		LogWriter: logwriter.New(adap),
	}
	return connector
}

type Connector struct {
	config *conf.Config
	adap   *adapter.Adapter
	LogWriter io.Writer
}
