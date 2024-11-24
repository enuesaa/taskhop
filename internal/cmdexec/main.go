package cmdexec

import (
	"io"
	"os/exec"
)

func New() I {
	return &Impl{}
}

type I interface {
	Exec(writer io.Writer, workdir string, command string, args []string) (*exec.Cmd, error)
	Kill(cmd *exec.Cmd) error
}
type Impl struct{}
