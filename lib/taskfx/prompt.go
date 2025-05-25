package taskfx

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

var ErrRegisterNotAvailable = errors.New("register not available")
var ErrAgentUnHealthy = errors.New("agent unhealthy")

func (i *TaskSrv) StartPrompt() error {
	if i.current.Status != StatusRegistration {
		return ErrRegisterNotAvailable
	}
	i.current.Status = StatusPrompt

	go i.monitor()
	go i.prompt()

	return nil
}

func (i *TaskSrv) monitor() {
	for {
		time.Sleep(5 * time.Second)

		if time.Since(i.lastHealthy) > 5*time.Second {
			i.errch <- ErrAgentUnHealthy
			break
		}
	}
}

var ErrPromptExit = errors.New("exit")

func (i *TaskSrv) prompt() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")
		if !scanner.Scan() {
			break
		}
		i.current.Text = scanner.Text()

		if i.current.Text == "exit" {
			i.errch <- ErrPromptExit
			break
		}
		i.current.Status = StatusProceeding

		<-i.completedch
	}
}
