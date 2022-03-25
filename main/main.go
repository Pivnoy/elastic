package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/elastic/go-elasticsearch/v8"
	"io/ioutil"
	"log"
)

func main() {
	cert, err := ioutil.ReadFile("/home/evg/http_ca.crt")
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200/",
		},
		Username: "elastic",
		Password: "+Nk3RJlDNpE=tEM89iU5",
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("Error")
	}
	res, err := es.Info()
	if err != nil {
		log.Fatal("Error 1")
	}
	if res.IsError() {
		fmt.Println("Error 2")
	}
	var buffer bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"status": "CI PASS",
			},
		},
	}
	if err := json.NewEncoder(&buffer).Encode(query); err != nil {
		log.Fatal("Error in buffer request")
	}

	res, err = es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("ci_runs"),
		es.Search.WithBody(&buffer),
		es.Search.WithTrackTotalHits(true),
	)
	if res.IsError() {
		log.Fatal("Error in searching")
	}
	var e map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		log.Fatal("Error in parse")
	}
	hitsMap := e["hits"].(map[string]interface{})
	tests := hitsMap["hits"].([]interface{})
	var results = make([]map[string]interface{}, 10)
	for _, k := range tests {
		results = append(results, k.(map[string]interface{}))
	}
	var source = make([]map[string]interface{}, 10)
	for _, k := range results {
		if k["_source"] != nil {
			source = append(source, k["_source"].(map[string]interface{}))
		}
		//fmt.Println(k["_source"])
	}
	for _, l := range source {
		if l["testlist"] != nil {
			fmt.Println(l["testlist"])
		}
	}
}
