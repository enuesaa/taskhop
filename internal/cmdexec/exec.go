package cmdexec

import (
	"io"
	"os/exec"
)

func (i *Impl) Exec(writer io.Writer, workdir string, command string, args []string) (*exec.Cmd, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = workdir

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return cmd, nil
}
