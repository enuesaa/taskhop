package connector

import "github.com/enuesaa/taskhop/internal"

func New(container internal.Container) *Connector {
	return &Connector{container}
}

type Connector struct {
	internal.Container
}
