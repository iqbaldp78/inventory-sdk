package object

import (
	"testing"
	"time"
)

func TestTimestampValue(t *testing.T) {
	input := []string{"2016-01-02 00:00:00", "02-01-2016"}
	output := []struct{
		val interface{}
		err interface{}
	}{
		{"2016-01-02 00:00:00", nil},
		{nil, nil},
	}

	for index := range input {
		obj := Timestamp{}
		obj.Time, _ = time.Parse(timestampLayout, input[index])
		val, err := obj.Value()
		if val != output[index].val && err != output[index].err {
			t.Errorf("Expected result to be `%v`. Got `%v %v`", output[index], val, err)
		}
	}
}

func TestTimestampScan(t *testing.T) {
	dt, _ := time.Parse(timestampLayout, "2016-01-02 00:00:00")
	input := []interface{}{
		nil,
		"2016-01-02 00:00:00",
		[]byte("2016-01-02 00:00:00"),
		dt,
	}
	output := []time.Time{
		time.Time{},
		dt,
		dt,
		dt,
	}

	for index := range input {
		obj := Timestamp{}
		obj.Scan(input[index])
		if !obj.Equal(output[index]) {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], obj.Time)
		}
	}
}

func TestTimestampMarshalJSON(t *testing.T) {
	input := "2016-01-02 00:00:00"
	output := []byte("\"2016-01-02 00:00:00\"")

	obj := Timestamp{}
	obj.Time, _ = time.Parse(timestampLayout, input)
	result, _ := obj.MarshalJSON()
	if string(result) != string(output) {
		t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
	}
}

func TestTimestampUnmarshalJSON(t *testing.T) {
	input := []byte("\"2016-01-02 00:00:00\"")
	output := "2016-01-02 00:00:00"

	obj := Timestamp{}
	obj.UnmarshalJSON(input)
	result, _ := time.Parse(timestampLayout, output)
	if !obj.Equal(result) {
		t.Errorf("Expected result to be `%v`. Got `%v`", obj.Time, result)
	}
}
