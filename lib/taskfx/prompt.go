package taskfx

import (
	"bufio"
	"fmt"
	"os"
)

func (i *TaskSrv) Prompt() error {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("> ")
	if !scanner.Scan() {
		i.endCh <- true
		return fmt.Errorf("end")
	}

	text := scanner.Text()
	if text == "exit" {
		i.endCh <- true
		return fmt.Errorf("end")
	}
	i.current.Cmd = text
	i.current.Status = StatusProceeding

	return nil
}
