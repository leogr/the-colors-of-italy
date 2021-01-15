package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/leogr/the-colors-of-italy/pkg/crawler"
)

func main() {

	data, err := crawler.Governo()
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile("data.json", b, 0644); err != nil {
		log.Fatal(err)
	}
}
