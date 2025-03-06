package yarbay

import (
	"github.com/guneyin/yarbay/modules/db"
	"github.com/guneyin/yarbay/modules/elastic"
	"github.com/guneyin/yarbay/modules/fiber"
	"github.com/guneyin/yarbay/modules/grpc"
	"github.com/guneyin/yarbay/modules/nats"
	"github.com/guneyin/yarbay/modules/otel"
	"github.com/guneyin/yarbay/modules/store"
	"testing"
	"time"
)

type TestApp struct {
	*App
}

func NewTestApp() *TestApp {
	app := NewApp(&Config{
		Name:  "test-app",
		Title: "Test App",
	})
	return &TestApp{app}
}

func (ta *TestApp) testRun() error {
	errCh := make(chan error)
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		errCh <- ta.Start()
	}()

	select {
	case err := <-errCh:
		return err
	case <-ticker.C:
		if ta.bootstrapped {
			break
		}
	}

	return nil
}

func (ta *TestApp) RunTest(t *testing.T, tests ...func(t *testing.T, a *TestApp)) error {
	if err := ta.testRun(); err != nil {
		return err
	}

	for _, test := range tests {
		test(t, ta)
	}

	ta.Stop()
	return nil
}

func (ta *TestApp) WithFiber(fiber *fiber.Fiber) *TestApp {
	ta.mc.RegisterModule(fiber)
	return ta
}

func (ta *TestApp) WithDB(db *db.DB) *TestApp {
	ta.mc.RegisterModule(db)
	return ta
}

func (ta *TestApp) WithStore(store *store.Store) *TestApp {
	ta.mc.RegisterModule(store)
	return ta
}

func (ta *TestApp) WithNATS(nats *nats.NATS) *TestApp {
	ta.mc.RegisterModule(nats)
	return ta
}

func (ta *TestApp) WithGRPC(grpc *grpc.GRPC) *TestApp {
	ta.mc.RegisterModule(grpc)
	return ta
}

func (ta *TestApp) WithOtel(otel *otel.Otel) *TestApp {
	ta.mc.RegisterModule(otel)
	return ta
}

func (ta *TestApp) WithElastic(elastic *elastic.Elastic) *TestApp {
	ta.mc.RegisterModule(elastic)
	return ta
}
