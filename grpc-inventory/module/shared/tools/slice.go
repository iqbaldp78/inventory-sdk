package tools

import (
	"reflect"
	"strings"
)

//TypeSliceInt used as constant to represent type slice of integer
const TypeSliceInt = "[]int"

//TypeSliceString used as constant to represent type slice of string
const TypeSliceString = "[]string"

//TypeInt used as constant to represent type of integer
const TypeInt = "int"

//TypeString used as constant to represent type of string
const TypeString = "string"

//SliceInSlice used for check value exist on two slice
func SliceInSlice(slice1, slice2 interface{}) bool {
	type1 := reflect.TypeOf(slice1).String()
	type2 := reflect.TypeOf(slice2).String()
	if type1 != type2 {
		return false
	}

	switch type1 {
	case TypeSliceInt:
		return sliceInSliceInt(slice1.([]int), slice2.([]int))
	case TypeSliceString:
		return sliceInSliceString(slice1.([]string), slice2.([]string))
	}

	return false
}

func sliceInSliceInt(slice1, slice2 []int) bool {
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

func sliceInSliceString(slice1, slice2 []string) bool {
	for _, v1 := range slice1 {
		for _, v2 := range slice2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

//ValueInSlice used for checking value is exist on slice
func ValueInSlice(value interface{}, values interface{}) bool {
	type1 := reflect.TypeOf(value).String()
	type2 := reflect.TypeOf(values).String()
	if type2 == TypeSliceInt || type2 == TypeSliceString {
		if !strings.HasSuffix(type2, type1) {
			return false
		}

		switch reflect.TypeOf(value).String() {
		case TypeInt:
			return valueInSliceInt(value.(int), values.([]int))
		case TypeString:
			return valueInSliceString(value.(string), values.([]string))
		}
	}

	return false
}

func valueInSliceInt(value int, values []int) bool {
	for _, row := range values {
		if value == row {
			return true
		}
	}
	return false
}

func valueInSliceString(value string, values []string) bool {
	for _, row := range values {
		if value == row {
			return true
		}
	}
	return false
}

//UniqueInt used for remove duplicate from slice of integer
func UniqueInt(list *[]int) {
	found := make(map[int]bool)
	j := 0
	for i, x := range *list {
		if !found[x] {
			found[x] = true
			(*list)[j] = (*list)[i]
			j++
		}
	}
	*list = (*list)[:j]
}

//UniqueString used for remove duplicate from slice of string
func UniqueString(list *[]string) {
	found := make(map[string]bool)
	j := 0
	for i, x := range *list {
		if !found[x] {
			found[x] = true
			(*list)[j] = (*list)[i]
			j++
		}
	}
	*list = (*list)[:j]
}
