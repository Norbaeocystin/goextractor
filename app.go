package main

import (
	"goextractor/pkg/goextractor"
	"log"
)

func main() {
	b, _ := goextractor.GetWithHeaders("https://www.kohls.com/", goextractor.HeadersDefault)
	doc, _ := goextractor.GetDoc(b)
	results := goextractor.Extract(doc, goextractor.Ids2name)
	for k, v := range results {
		log.Println(k, v)
	}
}
