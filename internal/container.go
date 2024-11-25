package internal

import (
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/cmdsfile"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
	"github.com/enuesaa/taskhop/internal/runnermg"
)

func NewContainer(
	cmdexeci cmdexec.I,
	cmdsfilei cmdsfile.I,
	fsi fs.I,
	loggingi logging.I,
	runnermgi runnermg.I,
) Container {
	return Container{
		cmdexec:  cmdexeci,
		cmdsfile: cmdsfilei,
		fs:       fsi,
		logging:  loggingi,
		runnermg: runnermgi,
	}
}

type Container struct {
	cmdexec  cmdexec.I
	cmdsfile cmdsfile.I
	fs       fs.I
	logging  logging.I
	runnermg runnermg.I
}
