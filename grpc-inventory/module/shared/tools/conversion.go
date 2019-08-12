package tools

import (
	"fmt"
	"strconv"
	"strings"
)

//SliceStringToString used for convert slice string to string
func SliceStringToString(list []string, separator string, wraper ...string) string {
	output := ""
	length := len(list)
	for index, value := range list {
		if output == "" {
			if len(wraper) > 0 {
				output = fmt.Sprintf("%v%v%v", wraper[0], value, wraper[0])
			} else {
				output = fmt.Sprintf("%v", value)
			}
		} else {
			if len(wraper) > 0 {
				output = fmt.Sprintf("%v %v%v%v", output, wraper[0], value, wraper[0])
			} else {
				output = fmt.Sprintf("%v %v", output, value)
			}
		}

		if index < (length - 1) {
			output = fmt.Sprintf("%v %v", output, separator)
		}
	}
	return output
}

//SliceIntToString used for convert slice integer to string
func SliceIntToString(slice []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(slice), " ", delim, -1), "[]")
}

//StringToSliceInt used for convert string to slice integer
func StringToSliceInt(source string) []int {
	arrayString := strings.Split(source, ", ")
	result := make([]int, len(arrayString))
	for i := range result {
		result[i], _ = strconv.Atoi(arrayString[i])
	}
	return result
}
