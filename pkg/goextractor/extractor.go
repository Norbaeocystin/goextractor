package goextractor

// Get Data Via Using chrome
func GetDataViaChrome(urlstring string) map[string]interface{} {
	s, _ := GetByChrome(urlstring)
	doc, _ := GetDoc([]byte(s))
	results := Extract(doc, Ids2name)
	data := make(map[string]interface{})
	data["source"] = urlstring
	for k, v := range results {
		data[k] = v
	}
	return data
}

// Get Data
func GetData(urlstring string) map[string]interface{} {
	s, _ := GetWithHeaders(urlstring, HeadersDefault)
	doc, _ := GetDoc([]byte(s))
	results := Extract(doc, Ids2name)
	data := make(map[string]interface{})
	data["source"] = urlstring
	for k, v := range results {
		data[k] = v
	}
	return data
}