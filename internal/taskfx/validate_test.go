package taskfx

import (
	"testing"

	"github.com/enuesaa/taskhop/internal/cli"
	"github.com/enuesaa/taskhop/internal/taskfx/repository"
	"github.com/stretchr/testify/assert"
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

	taski := New(cli.Config{}, &mock)
	task, err := taski.Read()
	assert.NoError(t, err)
	assert.NoError(t, taski.Validate(task))
}
