package goextractor

import (
	"fmt"
	"strings"
)

func GetString(records []map[string]interface{}) string {
	data := make(map[string]map[int]string)
	for index, record := range records {
		for key,value := range record {
			var s string
			s = fmt.Sprintf("%v", value)
			array, ok := value.([]string)
			if ok {
				s = strings.Join(array, ",")
			}
			s = strings.Replace(s, "\n"," ",-1)
			s = strings.Replace(s, "\r"," ",-1)
			s = strings.Replace(s, "\t"," ",-1)
			doc, ok := data[key]
			if ok == false {
				doc := make(map[int]string)
				doc[index] = s
				data[key] = doc
			} else {
				if doc == nil {
					doc := make(map[int]string)
					doc[index] = s
					data[key] = doc
				} else {
					doc[index] = s
					data[key] = doc
				}
			}
		}
		index += 1
	}
	lines := MakeLines(data, len(records))
	text := strings.Join(lines, "\n")
	return text
}


func MakeLines(data map[string]map[int]string, lenght int)[]string{
	lines := make([]string, lenght + 1)
	cyclus := make([]int, lenght)
	header := make( []string, 0)
	for k, _ := range data {
		header = append(header, k)
	}
	firstline := strings.Join(header, "\t")
	lines[0] = firstline
	for idx, _ := range cyclus {
		line := make([]string,0)
		for _, column := range header {
			value, bool := data[column][idx]
			if bool == false {
				value = ""
			}
			line = append(line,value)

		}
		lines[idx+1] = strings.Join(line, "\t")
	}
	return lines
}

