package modules

import (
	"github.com/guneyin/yarbay/logger"
)

type Module interface {
	Name() string
	Start() error
	Stop() error
}

type Controller struct {
	modules map[string]Module
}

func NewController() *Controller {
	return &Controller{
		modules: make(map[string]Module),
	}
}

func (mc *Controller) Boostrap() (bool, error) {
	errCh := make(chan error)
	for _, mod := range mc.modules {
		go func() {
			logger.Info("[INFO] starting %s module\n", mod.Name())
			if err := mod.Start(); err != nil {
				errCh <- err
			}
		}()
	}

	for err := range errCh {
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (mc *Controller) Shutdown() {
	for _, mod := range mc.modules {
		logger.Warn("[INFO] shutting down %s module\n", mod.Name())
		if err := mod.Stop(); err != nil {
			logger.Error(err)
		}
	}
}

func (mc *Controller) RegisterModule(m Module) {
	mc.modules[m.Name()] = m
}

func (mc *Controller) GetModule(name string) Module {
	return mc.modules[name]
}
