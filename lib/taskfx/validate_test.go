package taskfx

import (
	"strings"
	"testing"

	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib/taskfx/repository"
	"github.com/enuesaa/taskhop/lib/testsuite"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestValidate(t *testing.T) {
	suite := testsuite.New(t)
	defer suite.End()

	repo := repository.NewMockI(suite.Ctl())
	cmdsyml := strings.NewReader(`
title: Example

cmds:
  - echo hello!
  - pwd
`)
	repo.EXPECT().Read("cmds.yml").Return(cmdsyml, nil)

	var i ITaskSrv
	suite.UseOptions(
		fx.Supply(
			fx.Annotate(&conf.Config{}, fx.As(new(conf.Config))),
			fx.Annotate(repo, fx.As(new(repository.IRepository))),
		),
		fx.Provide(New),
		fx.Populate(&i),
	)

	task, err := i.Read()
	assert.NoError(t, err)
	assert.NoError(t, i.Validate(task))
}
