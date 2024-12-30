package cmdfx

import (
	"errors"
	"io"
	"os/exec"
)

func (i *CmdSrv) Exec(writer io.Writer, command string, workdir string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = workdir

	cmd.Stdout = writer
	cmd.Stderr = writer

	if err := cmd.Start(); err != nil {
		return err
	}
	if err := cmd.Wait(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return nil
		}
		return err
	}
	return nil
}
