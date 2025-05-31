package app

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/enuesaa/taskhop/app/gqlserver"
	"github.com/enuesaa/taskhop/conf"
	"github.com/enuesaa/taskhop/lib"
	"github.com/enuesaa/taskhop/lib/taskfx"
	"go.uber.org/fx"
)

func NewCommander(config *conf.Config, li *lib.Lib, lc fx.Lifecycle, shutdowner fx.Shutdowner) *Commander {
	commander := &Commander{
		config:     config,
		lib:        li,
		lc:         lc,
		shutdowner: shutdowner,
	}
	return commander
}

type Commander struct {
	config     *conf.Config
	lib        *lib.Lib
	lc         fx.Lifecycle
	shutdowner fx.Shutdowner
	server     *http.Server
}

func (a *Commander) Run(ctx context.Context) error {
	router := gqlserver.New(a.config, a.lib)
	a.server = &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	go a.listen(ctx)
	go a.monitor(ctx)
	a.printBanner(ctx)

	a.lc.Append(fx.Hook{
		OnStop: a.close,
	})
	return nil
}

func (a *Commander) listen(ctx context.Context) {
	if err := a.server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return
		}
		a.lib.Log.Info(ctx, "Error: %s", err.Error())
	}
}

func (a *Commander) monitor(ctx context.Context) {
	for err := range a.lib.Task.Subscribe() {
		if !errors.Is(err, taskfx.ErrPromptExit) {
			a.lib.Log.Info(ctx, "Error: %s", err.Error())
		}
		a.shutdowner.Shutdown() //nolint:errcheck
	}
}

func (a *Commander) close(ctx context.Context) error {
	if a.server != nil {
		return a.server.Shutdown(ctx)
	}
	return nil
}

func (a *Commander) printBanner(ctx context.Context) {
	addr := a.calcCommanderAddr()
	a.lib.Log.Info(ctx, "running!")
	a.lib.Log.Info(ctx, "┌───────────────────────────────────────────────────────")
	a.lib.Log.Info(ctx, "│ To launch the agent:")
	a.lib.Log.Info(ctx, "│   taskhop-agent %s:3000", addr)
	a.lib.Log.Info(ctx, "└───────────────────────────────────────────────────────")
}

func (a *Commander) calcCommanderAddr() string {
	// see https://stackoverflow.com/questions/23558425
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "localhost"
	}
	defer conn.Close() //nolint:errcheck

	addr, ok := conn.LocalAddr().(*net.UDPAddr)
	if !ok {
		return "localhost"
	}
	return addr.IP.To4().String()
}
