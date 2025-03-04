package yarbay

import (
	"context"
	"errors"
	"fmt"
	"github.com/guneyin/yarbay/modules"
	"os/signal"
	"syscall"
)

type Config struct {
	Name  string
	Title string
}

type App struct {
	config *Config
	mc     *modules.Controller
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

	errCh := make(chan error)

	go func() {
		errCh <- a.mc.Boostrap(ctx)
	}()

	select {
	case err := <-errCh:
		return fmt.Errorf("error on start: %v", err)
	case <-ctx.Done():
		a.mc.Shutdown()
		return errors.New("interrupted")
	}
}

func (a *App) Market() modules.Market {
	return a.mc
}
