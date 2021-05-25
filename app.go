package main

import (
	"bufio"
	"encoding/json"
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
	tsv := flag.Bool("tsv", false, "output as tsv - delimiter tab, default false - output will be as new line delmited json")
	retries := flag.Int("retries", 3, "number of retries if nothing returns")
	flag.Parse()
	if *url != "" && *file == ""{
		// b, _ := goextractor.GetWithHeaders("https://mdex.com/#/", goextractor.HeadersDefault)
		var data map[string]interface{}

		for i := 0; i <= *retries; i++ {
			if *chrome {
				data = goextractor.GetDataViaChrome(*url)
			} else {
				data = goextractor.GetData(*url)
			}

			if len(data) == 1 && i != 3 {
				continue
			}
			break
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
				// retries logic added
				for i := 0; i <= *retries; i++ {
					if *chrome {
						data = goextractor.GetDataViaChrome(line)
					} else {
						data = goextractor.GetData(line)
					}

					if len(data) == 1 && i != 3 {
						continue
					}

					if len(data) > 0 && *tsv {
						results = append(results, data)
					} else {
						b, _ := json.Marshal(data)
						fmt.Println(string(b))
					}

					break
				}
			}
		}
		if *tsv {
			txt := goextractor.GetString(results)
			fmt.Println(txt)
		}
	}
}
