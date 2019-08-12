package tools

import (
	"testing"
)

func TestSliceInSliceInt(t *testing.T) {
	input := []struct{
		slice1 []int
		slice2 []int
	}{
		{[]int{1, 2}, []int{3, 2}},
		{[]int{1, 2}, []int{3, 4}},
	}
	output := []bool{true, false}
	for index := range input {
		result := sliceInSliceInt(input[index].slice1, input[index].slice2)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestSliceInSliceString(t *testing.T) {
	input := []struct{
		slice1 []string
		slice2 []string
	}{
		{[]string{"abc", "cde"}, []string{"efg", "cde"}},
		{[]string{"abc", "cde"}, []string{"efg", "hij"}},
	}
	output := []bool{true, false}
	for index := range input {
		result := sliceInSliceString(input[index].slice1, input[index].slice2)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestValueInSliceInt(t *testing.T) {
	input := []struct{
		value int
		values []int
	}{
		{2, []int{3, 2}},
		{2, []int{3, 4}},
	}
	output := []bool{true, false}
	for index := range input {
		result := valueInSliceInt(input[index].value, input[index].values)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestValueInSliceString(t *testing.T) {
	input :=[]struct{
		value string
		values []string
	}{
		{"cde", []string{"efg", "cde"}},
		{"cde", []string{"efg", "hij"}},
	}
	output := []bool{true, false}
	for index := range input {
		result := valueInSliceString(input[index].value, input[index].values)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestUniqueInt(t *testing.T) {
	input := []int{1, 2, 2, 3}
	output := []int{1, 2, 3}
	UniqueInt(&input)
	for i := range output {
		if output[i] != input[i] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output, input)
		}
	}
}

func TestUniqueString(t *testing.T) {
	input := []string{"a", "b", "b", "c"}
	output := []string{"a", "b", "c"}
	UniqueString(&input)
	for i := range output {
		if output[i] != input[i] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output, input)
		}
	}
}

func TestValueInSlice(t *testing.T) {
	input := []struct{
		value interface{}
		values interface{}
	}{
		{2, []bool{true, false}},
		{2, []string{"3", "2"}},
		{2, []int{3, 2}},
		{"2", []string{"3", "2"}},
	}
	output := []bool{false, false, true, true}

	for index := range input {
		result := ValueInSlice(input[index].value, input[index].values)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}

func TestSliceInSlice(t *testing.T) {
	input := []struct{
		slice1 interface{}
		slice2 interface{}
	}{
		{[]bool{false, true}, []bool{true, false}},
		{[]int{3, 2}, []string{"3", "2"}},
		{[]int{1, 2}, []int{3, 2}},
		{[]string{"abc", "cde"}, []string{"efg", "cde"}},
	}
	output := []bool{false, false, true, true}

	for index := range input {
		result := SliceInSlice(input[index].slice1, input[index].slice2)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}
