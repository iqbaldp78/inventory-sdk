package tools

import (
	"testing"
	"time"
)

func TestSetUTCPlus7(t *testing.T) {
	str1 := "01 Jan 16 07:00 GMT"
	str2 := "01 Jan 16 00:00 GMT"
	input, _ := time.Parse(time.RFC822, str1)
	output, _ := time.Parse(time.RFC822, str2)

	result := SetUTCPlus7(input)
	if !result.Equal(output) {
		t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
	}
}

func TestLocalTime(t *testing.T) {
	str1 := "01 Jan 16 00:00 GMT"
	str2 := "01 Jan 16 07:00 GMT"
	input, _ := time.Parse(time.RFC822, str1)
	output, _ := time.Parse(time.RFC822, str2)

	result := LocalTime(input, "Asia/Jakarta")
	if !result.Equal(output) {
		t.Errorf("Expected result to be `%v`. Got `%v`", output, result)
	}
}

func TestTimeToString(t *testing.T) {
	input := []struct {
		input string
		to    string
	}{
		{"01 Jan 16 07:00 GMT", "date"},
		{"01 Jan 16 07:00 GMT", "timestamp"},
		{"2016-01-01 07:00:00", "date"},
		{"2016-01-01", "timestamp"},
	}
	output := []string{"2016-01-01", "2016-01-01 07:00:00", "", ""}
	for index := range input {
		temp, _ := time.Parse(time.RFC822, input[index].input)
		result := TimeToString(temp, input[index].to)
		if result != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], result)
		}
	}
}
