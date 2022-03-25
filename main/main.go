package main

import (
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
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "+Nk3RJlDNpE=tEM89iU5",
		CACert: cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatal("Error")
	}
	res, err := es.Info()
	if err != nil {
		log.Fatal("Error 1")
	}
	fmt.Println(res)
	if res.IsError() {
		fmt.Println("Error 2")
	}
}
