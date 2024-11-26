package internal

import (
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/runnermg"
)

func NewContainer(
	cmdexeci cmdexec.I,
	cmdsfilei cmdsfile.I,
	loggingi logging.I,
	runnermgi runnermg.I,
) Container {
	return Container{
		Cmdexec:  cmdexeci,
		Cmdsfile: cmdsfilei,
		Logging:  loggingi,
		Runnermg: runnermgi,
	}
}

type Container struct {
	Cmdexec  cmdexec.I
	Cmdsfile cmdsfile.I
	Logging  logging.I
	Runnermg runnermg.I
}
