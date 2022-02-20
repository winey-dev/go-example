package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	//	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Result struct {
	ScrollID string `json:"_scroll_id,omitempty"`
	Hit      Hit    `json:"hits,omitempty"`
}

type Hit struct {
	Hits []Hits `json:"hits,omitempty"`
}

type Hits struct {
	Index  string                 `json:"_index"`
	Type   string                 `json:"_type"`
	Score  string                 `json:"_score"`
	Source map[string]interface{} `json:"_source"`
	Sort   []int                  `json:"sort"`
}

func (e *elastic) Insert(index string, v interface{}) (*esapi.Response, error) {
	body, err := json.Marshal(v)
	if err != nil {
		log.Printf("json marshal failed. err=%v\n", err)
		return nil, err
	}
	return e.cli.Index(index, bytes.NewReader(body))
}

func READMAP(r io.Reader) map[string]interface{} {
	str := READ(r)
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		log.Printf("json unmarshal failed. err=%v\n", err)
		return nil
	}
	return m
}

func READ(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()
}

func (e *elastic) Search(index string, page int) ([]Hits, error) {
	list := Hit{}

	res, err := e.cli.Search(
		e.cli.Search.WithIndex(index),
		e.cli.Search.WithScroll(time.Minute),
		e.cli.Search.WithSort("id"),
		e.cli.Search.WithSize(page),
	)

	if err != nil {
		log.Printf("search api failed. err=%v\n", err)
		return nil, err
	}

	str := READ(res.Body)
	res.Body.Close()
	if res.IsError() {
		log.Printf("search api response: %d\n", res.StatusCode)
		return nil, fmt.Errorf("response code %d", res.StatusCode)
	}

	result := Result{}
	err = json.Unmarshal([]byte(str), &result)
	if err != nil {
		log.Printf("json unmarshal failed. err=%v\n", err)
		return nil, err
	}
	log.Printf("Search Succ..\n")
	log.Printf("%v", result)

	defer e.cli.ClearScroll(e.cli.ClearScroll.WithScrollID(result.ScrollID))
	list.Hits = append(list.Hits, result.Hit.Hits...)

	for {
		res, err = e.cli.Scroll(e.cli.Scroll.WithScrollID(result.ScrollID), e.cli.Scroll.WithScroll(time.Minute))
		if err != nil {
			log.Printf("scroll api failed. err=%v\n", err)
			return nil, err
		}

		str = READ(res.Body)
		res.Body.Close()
		if res.IsError() {
			log.Printf("search api response: %d\n", res.StatusCode)
			return nil, fmt.Errorf("response code %d", res.StatusCode)
		}
		err = json.Unmarshal([]byte(str), &result)
		if err != nil {
			log.Printf("json unmarshal failed. err=%v\n", err)
			return nil, err
		}
		if len(result.Hit.Hits) == 0 {
			break
		}
		list.Hits = append(list.Hits, result.Hit.Hits...)
	}

	return list.Hits, err
}
