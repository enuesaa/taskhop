package runner

import (
	"fmt"

	"github.com/enuesaa/taskhop/runner/usecase"
)

func InvokeMain(u *usecase.UseCase) error {
	task, err := u.Register()
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", task)

	if err := u.UnArchive(); err != nil {
		return err
	}

	if err := u.Run(task); err != nil {
		return err
	}

	return nil
}
