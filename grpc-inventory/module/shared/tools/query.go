package tools

import (
	"fmt"
	"strings"
)

//SearchFilter used for generate search condition query
func SearchFilter(searchKey string, target interface{}) (output string) {
	var conditionList []string
	keys := strings.Split(searchKey, " ")
	for _, key := range keys {
		for _, field := range GetTag("filter", target) {
			conditionList = append(conditionList, fmt.Sprintf("%v ilike '%%%v%%'", field, key))
		}
	}
	return SliceStringToString(conditionList, "OR")
}
