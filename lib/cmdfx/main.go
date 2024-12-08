package cmdfx

import (
	"io"
	"os/exec"

	"github.com/enuesaa/taskhop/lib/cmdfx/repository"
)

func New(repo repository.I) I {
	return &Impl{
		repo: repo,
	}
}

type I interface {
	Exec(writer io.Writer, command string) (*exec.Cmd, error)
	MultiExec(writer io.Writer, commands []string, workdir string) error
	Kill(cmd *exec.Cmd) error
}
type Impl struct {
	repo repository.I
}

// TODO:
// おそらく cmd store みたいなのが必要
