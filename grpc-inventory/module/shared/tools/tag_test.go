package tools

import (
	"testing"
)

type foo struct{
	f1 string `filter:"abc"`
	f2 string `filter:"cde"`
	f3 bool 
}

func TestGetTag(t *testing.T) {
	input := []struct{
		tagName string
		obj interface{}
	}{
		{"filter", "test1"},
		{"filter", foo{"test1", "test1", false}},
	}
	output := [][]string{
		{},
		{"abc", "cde"},
	}

	for index := range input {
		result := GetTag(input[index].tagName, input[index].obj)
		for i := range result {
			if result[i] != output[index][i] {
				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
			}
		}
	}
}
