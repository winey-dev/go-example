package app

import (
	"bytes"
	"encoding/json"
	"log"

	//	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type MappingInfo struct {
	Key    string
	Type   string
	Format string // if type is date or time, set format
}

type mtype struct {
	mtype  string `json:"type,omitempty"`
	format string `json:"format,omitempty"`
}

type properties map[string]mtype

type mappings struct {
	properties properties `json:"properties"`
}

type indexMapping struct {
	mappings mappings `json:"mappings"`
}

func (e *elastic) IndicesCreate(index string, ms []MappingInfo) (*esapi.Response, error) {
	p := make(properties)
	for _, m := range ms {
		p[m.Key] = mtype{mtype: m.Type, format: m.Format}
	}

	query := indexMapping{
		mappings: mappings{
			properties: p,
		},
	}

	body, err := json.Marshal(query)
	if err != nil {
		log.Printf("json marshal failed. err=%v\n", err)
		return nil, err
	}
	return e.cli.Indices.Create(index, e.cli.Indices.Create.WithBody(bytes.NewReader(body)))
}
