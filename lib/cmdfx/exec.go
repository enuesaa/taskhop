package cmdfx

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
)

func (i *CmdSrv) Exec(writer io.Writer, command string, workdir string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Dir = workdir

	var buf bytes.Buffer
	multiwriter := io.MultiWriter(writer, &buf)
	defer func() {
		if buf.Len() == 0 {
			writer.Write([]byte("")) //nolint:errcheck
		}
	}()
	cmd.Stdout = multiwriter
	cmd.Stderr = multiwriter

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
