package runner

import "github.com/enuesaa/taskhop/runner/usecase"

func InvokeConnect(u *usecase.UseCase) error {
	return u.Connect()
}
