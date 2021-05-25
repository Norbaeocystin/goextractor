package main

import (
	"bufio"
	"flag"
	"fmt"
	"goextractor/pkg/goextractor"
	"log"
	"os"
	"strings"
)

func main() {
	// https://binance.org
	url := flag.String("url", "", "url to scrap")
	file := flag.String("file", "", "file to use to fetch urls")
	chrome := flag.Bool("chrome",false, "use chrome ")
	flag.Parse()
	if *url != "" && *file == ""{
		// b, _ := goextractor.GetWithHeaders("https://mdex.com/#/", goextractor.HeadersDefault)
		var data map[string]interface{}
		if *chrome {
			data = goextractor.GetDataViaChrome(*url)
		}else {
			data = goextractor.GetData(*url)
		}
		fmt.Println(data)
		txt := goextractor.GetString([]map[string]interface{}{data})
		fmt.Println(txt)
	}
	if *file != ""{
		f, err := os.Open(*file)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		results := make([]map[string]interface{},0)

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "https://") || strings.Contains(line, "http://"){
				var data map[string]interface{}
				if *chrome {
					data = goextractor.GetDataViaChrome(line)
				}else {
					data = goextractor.GetData(*url)
				}
				if len(data) > 0{
					results = append(results, data)
				}
			}
		}
		txt := goextractor.GetString(results)
		fmt.Println(txt)
	}
}
