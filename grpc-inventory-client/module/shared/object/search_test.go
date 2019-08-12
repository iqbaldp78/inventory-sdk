package object

import (
	"testing"
)

func TestSearchResultInitial(t *testing.T) {
	input := []SearchResult{{}, {1, nil}}
	output := []bool{true, false}

	for index := range input {
		result := input[index].IsInitial()
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}
