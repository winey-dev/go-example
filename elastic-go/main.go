package main

import (
	"bytes"
	"fmt"
	"io"
	"log"

	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/yiaw/go-example/elastic-go/app"
)

const (
	INDEX = "test_1"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func read(r io.Reader) string {
	var b bytes.Buffer
	b.ReadFrom(r)
	return b.String()

}
func main() {
	cfg := elastic.Config{
		Addresses: []string{"http://10.10.9.128:30613"},
	}

	es, err := app.EsNewClient(cfg)
	if err != nil {
		log.Printf("es new client failed. err=%v\n", err)
		return
	}

	/*
		maps := []app.MappingInfo{
			{Key: "id", Type: "integer"},
			{Key: "name", Type: "text"},
		}

		res, err := es.IndicesCreate(INDEX, maps)
		if err != nil {
			log.Printf("index create failed. err=%v", err)
			return
		}

		if res.IsError() {
			log.Printf("Index Create Status = %d\n", res.StatusCode)
			log.Println(READ(res.Body))
			return
		}

		for i := 0; i < 100; i++ {
			u := User{
				ID:   i,
				Name: fmt.Sprintf("smlee_%d", i),
			}
			res, err := es.Insert(INDEX, u)
			if err != nil {
				log.Printf("index insert failed. err=%v\n", err)
				return
			}
			if res.IsError() {
				log.Printf("Index insert Status = %d\n", res.StatusCode)
				log.Println(READ(res.Body))
				return
			}
		}
	*/

	hits, err := es.Search(INDEX, 5)
	if err != nil {
		log.Printf("Search Failed\n", err)
		return
	}
	fmt.Println(hits)
}
