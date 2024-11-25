package cmdexec

import (
	"io"
	"os/exec"
)

func (i *Impl) MultiExec(writer io.Writer, commands []string) error {
	for _, command := range commands {
		cmd := exec.Command("bash", "-c", command)

		cmd.Stdout = writer
		cmd.Stderr = writer

		if err := cmd.Start(); err != nil {
			return err
		}
		if err := cmd.Wait(); err != nil {
			return err
		}
	}
	return nil
}
