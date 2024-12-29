package testsuite

import (
	"testing"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"go.uber.org/mock/gomock"
)

func New(t *testing.T) TestSuite {
	return TestSuite{
		t: t,
	}
}

type TestSuite struct {
	t   *testing.T
	ctl *gomock.Controller
}

func (ts *TestSuite) Ctl() *gomock.Controller {
	ts.ctl = gomock.NewController(ts.t)
	return ts.ctl
}

func (ts *TestSuite) UseOptions(options ...fx.Option) {
	fxtest.New(
		ts.t,
		fx.Options(options...),
		fx.NopLogger,
	).RequireStart().RequireStop()
}

func (ts *TestSuite) End() {
	if ts.ctl != nil {
		ts.ctl.Finish()
	}
}
