package taskfx

import (
	"strings"
	"testing"

	"github.com/enuesaa/taskhop/cli"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/mock/gomock"
)

func TestValidate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := repository.NewMockI(ctl)
	cmdsyml := strings.NewReader(`
title: Example

cmds:
  - echo hello!
  - pwd
`)
	mockRepo.EXPECT().Read("cmds.yml").Return(cmdsyml, nil)

	var i I
	fxtest.New(
		t,
		fx.Supply(
			fx.Annotate(&cli.Cli{}, fx.As(new(cli.ICli))),
			fx.Annotate(mockRepo, fx.As(new(repository.I))),
		),
		fx.Provide(New),
		fx.Populate(&i),
		fx.NopLogger,
	).RequireStart().RequireStop()

	task, err := i.Read()
	assert.NoError(t, err)
	assert.NoError(t, i.Validate(task))
}
