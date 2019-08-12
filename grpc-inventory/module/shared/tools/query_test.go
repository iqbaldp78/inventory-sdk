package tools


import (
	"testing"
)

func TestSearchFilter(t *testing.T) {
	input:=struct{
		searchKey string
		target foo
	}{"abc def", foo{"test1", "test1", false}}
	output:="abc ilike '%abc%' OR cde ilike '%abc%' OR abc ilike '%def%' OR cde ilike '%def%'"

	// for index := range input {
		result := SearchFilter(input.searchKey, input.target)
		if result != output {
			t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
		}
	// }
}