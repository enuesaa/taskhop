package usecase

import (
	"github.com/enuesaa/taskhop/internal/cmdexec"
	"github.com/enuesaa/taskhop/internal/fs"
	"github.com/enuesaa/taskhop/internal/logging"
)

func New(fsi fs.I, logi logging.I, cmdexeci cmdexec.I) Usecase {
	return Usecase{
		fs: fsi,
		log: logi,
		cmdexec: cmdexeci,
	}
}

type Usecase struct {
	fs fs.I
	log logging.I
	cmdexec cmdexec.I
}
