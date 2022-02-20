package app

import (
	"log"

	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type elastic struct {
	cli *es.Client
}

func EsNewClient(cfg es.Config) (*elastic, error) {
	cli, err := es.NewClient(cfg)
	if err != nil {
		log.Printf("es new client failed. err=%v\n", err)
		return nil, err
	}

	return &elastic{
		cli: cli,
	}, nil
}

func (e *elastic) Info() (*esapi.Response, error) {
	return e.cli.Info()
}
