package tools

import "reflect"

//GetTag used for get tag from struct
func GetTag(tagName string, obj interface{}) (result []string) {
	ref := reflect.TypeOf(obj)
	if ref.Kind().String() != "struct" {
		return
	}

	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue
		}
		result = append(result, tag)
	}

	return
}
