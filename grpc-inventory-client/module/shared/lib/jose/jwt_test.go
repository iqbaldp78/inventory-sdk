package jose

import (
	"testing"
	"time"
)

func TestGenerateJWT(t *testing.T) {
	input := struct {
		data        Data
		expireHours time.Duration
	}{Data{1, 1, map[string]int{}}, time.Minute}
	output := ""

	result, _ := GenerateJWT(input.data, input.expireHours)
	if result == output {
		t.Errorf("Expected result not to be `%v`. Got `%v`", output, result)
	}
}

func TestGetClaims(t *testing.T) {
	input := []string{
		`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjg2MzUxNjg3NzAsImlhdCI6MTU2MTUyNzgzNSwiaXNzIjoiYXBpLmJvaWxlcnBsYXRlLmNvbSIsInBheSI6IntcInBheWxvYWRcIjpcImV5SnBaQ0k2TVN3aVozSnZkWEJmYVdRaU9qRXNJbkp2YkdVaU9udDlmUVwiLFwicHJvdGVjdGVkXCI6XCJleUpoYkdjaU9pSklVekkxTmlKOVwiLFwic2lnbmF0dXJlXCI6XCJmVERKbWpCQmdGZU9JcXN3U3dfQ1h0S2NURnZ6UVVIQTRVZjl6M3FNdTVZXCJ9Iiwic3ViIjoiYWNjZXNzIGdyYW50ZWQgZm9yIDYwMDAwMDAwMDAwIGhvdXIocykifQ.t4iO45XwlWVvMNuuBHJM-0Mr_HEqrBcfemT8sSMLz1Q`,
		`eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjg2MzUxNjg3NzAsImlhdCI6MTU2MTUyNzgzNSwiaXNzIjoiYXBpLmJvaWxlcnBsYXRlLmNvbSIsInBheSI6IntcInBheWxvYWRcIjpcImV5SnBaQ0k2TVN3aVozSnZkWEJmYVdRaU9qRXNJbkp2YkdVaU9udDlmUVwiLFwicHJvdGVjdGVkXCI6XCJleUpoYkdjaU9pSklVekkxTmlKOVwiLFwic2lnbmF0dXJlXCI6XCJmVERKbWpCQmdGZU9JcXN3U3dfQ1h0S2NURnZ6UVVIQTRVZjl6M3FNdTVZXCJ9Iiwic3ViIjoiYWNjZXNzIGdyYW50ZWQgZm9yIDYwMDAwMDAwMDAwIGhvdXIocykifQ`,
	}
	output := []int{
		5,
		0,
	}

	for index := range input {
		result, _ := GetClaims(input[index])
		if len(result) != output[index] {
			t.Errorf("Expected result to be `%v`. Got `%v`", output[index], len(result))
		}
	}
}
