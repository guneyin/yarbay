package yarbay

import (
	"context"
	"errors"
	"fmt"
	"github.com/guneyin/yarbay/modules"
	"os/signal"
	"syscall"
)

type App struct {
	config       *Config
	mc           *modules.Controller
	bootstrapped bool
}

func NewApp(config *Config) *App {
	return &App{
		config: config,
		mc:     modules.NewController(),
	}
}

func (a *App) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	okCh := make(chan bool)
	errCh := make(chan error)

	go func() {
		ok, err := a.mc.Boostrap()
		if err != nil {
			errCh <- err
		}
		okCh <- ok
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("error on start: %v", err)
	case <-okCh:
		a.bootstrapped = true
		return nil
	case <-ctx.Done():
		a.Stop()
		return errors.New("interrupted")
	}
}

func (a *App) Stop() {
	a.mc.Shutdown()
}

func (a *App) Bootstrapped() bool {
	return a.bootstrapped
}

func (a *App) Market() modules.Market {
	return a.mc
}
