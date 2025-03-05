package elastic

import (
	"github.com/elastic/go-elasticsearch/v8"
)

const ModuleName = "elasticsearch"

type Elastic struct {
	*elasticsearch.Client
	addr string
}

func New(addr string) *Elastic {
	return &Elastic{addr: addr}
}

func (e *Elastic) Name() string {
	return ModuleName
}

func (e *Elastic) Start() error {
	if e == nil {
		return nil
	}

	client, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{e.addr}})
	if err != nil {
		return err
	}

	e.Client = client
	return nil
}

func (e *Elastic) Stop() error {
	return nil
}
