package tools

import (
	"testing"
)

func TestSliceStringToString(t *testing.T) {
	input := []struct{
		list []string
		separator string
		wrapper []string
	}{
		{[]string{"a", "b"}, ",", []string{"'"}},
		{[]string{"a", "b"}, ",", []string{}},
	}
	output := []string{
		"'a' , 'b'",
		"a , b",
	}

	for index := range input {
		result := SliceStringToString(input[index].list, input[index].separator, input[index].wrapper...)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestSliceIntToString(t *testing.T) {
	input := struct{
		slice []int
		delim string
	}{[]int{1, 2}, ","}
	output := "1,2"
	result := SliceIntToString(input.slice, input.delim)
	if result != output {
		t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
	}
}

func TestStringToSliceInt(t *testing.T) {
	input := []string{"1, 2, 3", "1, a, 3"}
	output := [][]int{{1, 2, 3}, {1, 0, 3}}
	for index := range input {
		result := StringToSliceInt(input[index])
		for i := range result {
			if result[i] != output[index][i] {
				t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
			}
		}
	}
}
