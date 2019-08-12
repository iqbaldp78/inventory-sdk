package object

import (
	"testing"
	"time"
)

func TestDateValue(t *testing.T) {
	input := []string{"2016-01-02", "02-01-2016"}
	output := []struct{
		val interface{}
		err interface{}
	}{
		{"2016-01-02", nil},
		{nil, nil},
	}

	for index := range input {
		obj := Date{}
		obj.Time, _ = time.Parse(dateLayout, input[index])
		val, err := obj.Value()
		if val != output[index].val && err != output[index].err {
			t.Errorf("Expected result to be `%v`. Got `%v %v`", output[index], val, err)
		}
	}
}

func TestDateScan(t *testing.T) {
	dt, _ := time.Parse(dateLayout, "2016-01-02")
	input := []interface{}{
		nil,
		"2016-01-02",
		[]byte("2016-01-02"),
		dt,
	}
	output := []time.Time{
		time.Time{},
		dt,
		dt,
		dt,
	}

	for index := range input {
		obj := Date{}
		obj.Scan(input[index])
		if !obj.Equal(output[index]) {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], obj.Time)
		}
	}
}

func TestDateMarshalJSON(t *testing.T) {
	input := "2016-01-02"
	output := []byte("\"2016-01-02\"")

	obj := Date{}
	obj.Time, _ = time.Parse(dateLayout, input)
	result, _ := obj.MarshalJSON()
	if string(result) != string(output) {
		t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
	}
}

func TestDateUnmarshalJSON(t *testing.T) {
	input := []byte("\"2016-01-02\"")
	output := "2016-01-02"

	obj := Date{}
	obj.UnmarshalJSON(input)
	result, _ := time.Parse(dateLayout, output)
	if !obj.Equal(result) {
		t.Errorf("Expected result to be `%v`. Got `%v`", obj.Time, result)
	}
}
