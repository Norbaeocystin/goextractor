package main

import (
	"goextractor/pkg/goextractor"
	"log"
)

func main() {
	b, _ := goextractor.GetWithHeaders("https://mustangtoken.net/", goextractor.HeadersDefault)
	doc, _ := goextractor.GetDoc(b)
	results := goextractor.Extract(doc, goextractor.Ids2name)
	for k, v := range results {
		log.Println(k, v)
	}
}
