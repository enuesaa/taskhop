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
		return fmt.Errorf("end")
	}

	text := scanner.Text()
	if text == "exit" {
		return fmt.Errorf("end")
	}
	i.current.Cmds = []string{text}

	return nil
}
