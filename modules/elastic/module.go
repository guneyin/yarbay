package elastic

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	elasticContainer "github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

const ModuleName = "elasticsearch"

type Elastic struct {
	*elasticsearch.Client
	container *elasticContainer.ElasticsearchContainer
	addr      string
	test      bool
}

func New(addr string) *Elastic {
	return &Elastic{addr: addr}
}

func NewTest() *Elastic {
	return &Elastic{test: true}
}

func (e *Elastic) Name() string {
	return ModuleName
}

func (e *Elastic) Start() error {
	if e == nil {
		return nil
	}

	if e.test {
		container, err := elasticContainer.Run(context.Background(), "docker.elastic.co/elasticsearch/elasticsearch:8.9.0")
		if err != nil {
			return err
		}
		e.container = container

		host, err := container.Host(context.Background())
		if err != nil {
			return err
		}
		e.addr = host
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
