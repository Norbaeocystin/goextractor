package main

import (
	"goextractor/pkg/goextractor"
	"log"
)

func main() {
	// b, _ := goextractor.GetWithHeaders("https://mdex.com/#/", goextractor.HeadersDefault)
	s, _ := goextractor.GetByChrome("https://evodefi.com/",)
	doc, _ := goextractor.GetDoc([]byte(s))
	results := goextractor.Extract(doc, goextractor.Ids2name)
	for k, v := range results {
		log.Println(k, v)
	}
}
