package goextractor

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// Returns *goquery Document from []byte
func GetDoc(b []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	return doc, err
}

// Extract external links from goquery document based on ids2name
func Extract(doc *goquery.Document, ids2name map[string]string) map[string][]string {
	results := make(map[string][]string)
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href := selection.AttrOr("href", "")
		// clean first href
		hrefClean := Clean(href)
		for id, name := range ids2name {
			if strings.Contains(hrefClean, id) {
				result, ok := results[name]
				chref := CleanLink(href)

				if id == "play.google.com" {
					chref = href
				}

				if ok {
					// check if the link is not already added
					alreadyAdded := false
					for _, value := range result {
						if value == chref {
							alreadyAdded = true
						}
					}
					if !alreadyAdded {
						result = append(result, chref)
						results[name] = result
					}
				} else {
					results[name] = []string{
						chref,
					}
				}
			}
		}
	})
	return results
}

// Clean href and returns just domain
func Clean(href string) string {
	href = strings.TrimPrefix(href, "https://")
	href = strings.TrimPrefix(href, "http://")
	href = strings.TrimPrefix(href, "www.")
	for _, char := range []string{"/", "?", "#"} {
		if strings.Contains(href, char) {
			href = strings.SplitN(href, char, 2)[0]
		}
	}
	return href
}

// Removes paramaters and hashtag stuff
func CleanLink(href string) string {
	for _, char := range []string{"?", "#"} {
		if strings.Contains(href, char) {
			href = strings.SplitN(href, char, 2)[0]
		}
	}
	return href
}
