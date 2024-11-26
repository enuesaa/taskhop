package cmdexec

import (
	"io"
	"os/exec"

	"github.com/enuesaa/taskhop/internal/cmdexec/repository"
)

func New(repo repository.I) I {
	return &Impl{
		repo: repo,
	}
}

type I interface {
	Exec(writer io.Writer, command string) (*exec.Cmd, error)
	MultiExec(writer io.Writer, commands []string) error
	Kill(cmd *exec.Cmd) error
}
type Impl struct {
	repo repository.I
}

// TODO:
// おそらく cmd store みたいなのが必要
