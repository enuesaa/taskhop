package taskfx

import (
	"testing"

	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestValidate(t *testing.T) {
	mock := repository.Mock{
		Cmds: `
title: Example

cmds:
  - echo hello!
  - pwd
`,
	}

	var i I
	fxtest.New(
		t,
		fx.Supply(
			cli.Cli{}, // TODO
			fx.Annotate(&mock, fx.As(new(repository.I))),
		),
		fx.Provide(New),
		fx.Populate(&i),
		fx.NopLogger,
	).RequireStart().RequireStop()

	task, err := i.Read()
	assert.NoError(t, err)
	assert.NoError(t, i.Validate(task))
}
